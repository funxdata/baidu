package impl

import (
	"context"
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/funxdata/baidu/core"
	"github.com/funxdata/baidu/internal/biz"
	pb "github.com/funxdata/baidu/service"
	"github.com/sirupsen/logrus"
)

type FaceServer struct {
	*core.Core
}

func NewFaceServer(key, secret string) *FaceServer {
	return &FaceServer{
		Core: core.NewCore(key, secret),
	}
}
func (s *FaceServer) Search(ctx context.Context, in *pb.SearchOption) (*pb.SearchResult, error) {
	req := map[string]interface{}{
		"image":         in.Image,
		"image_type":    in.ImageType,
		"group_id_list": in.GroupIdList,
	}
	var ret struct {
		core.BaiduResponse `json",inline"`
		pb.SearchResult    `json:"result"`
	}
	if err := s.Post(biz.URLFaceSearch, req, &ret); err != nil {
		logrus.Errorf("searched failed, %s", err)
		return nil, err
	}

	if err := ret.Err("search"); err != nil {
		return nil, err
	}

	return &ret.SearchResult, nil
}
func (s *FaceServer) Detect(ctx context.Context, in *pb.DetectOption) (*pb.DetectResult, error) {
	img, err := s.encodeImg(in.Image)
	if err != nil {
		return nil, err
	}
	req := map[string]interface{}{
		"image":        img,
		"image_type":   in.ImageType,
		"face_field":   "age,beauty,expression,face_shape,gender,glasses,landmark,race,quality,face_type",
		"max_face_num": 10,
	}
	var ret struct {
		core.BaiduResponse
		pb.DetectResult
	}
	err = s.Post(biz.URLFaceDetect, req, &ret)
	if err != nil {
		logrus.Errorf("detect face failed, %s", err)
		return nil, err
	}

	if ret.ErrorCode > 0 {
		return nil, fmt.Errorf("detect failed, (%v) %s", ret.ErrorCode, ret.ErrorMsg)
	}

	return &ret.DetectResult, nil
}

func (s *FaceServer) encodeImg(imgURL string) (encodeRet string, err error) {
	file, err := os.Open(imgURL)
	if err != nil {
		return "", err
	}
	defer file.Close()
	buf, err := ioutil.ReadAll(file)
	if err != nil {
		return "", err
	}
	encodeRet = base64.StdEncoding.EncodeToString(buf)
	return
}
