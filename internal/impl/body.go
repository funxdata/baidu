package impl

import (
	"bytes"
	"context"
	"encoding/base64"
	"fmt"
	"image"
	"io/ioutil"
	"net/url"
	"os"
	"strconv"

	"github.com/funxdata/baidu/core"
	"github.com/funxdata/baidu/internal/biz"
	pb "github.com/funxdata/baidu/service"
	"github.com/sirupsen/logrus"
)

type BodyTrackServer struct {
	*core.Core
}

func NewBodyTrackServer(key, secret string) *BodyTrackServer {
	return &BodyTrackServer{
		Core: core.NewCore(key, secret),
	}
}

func (s *BodyTrackServer) Track(ctx context.Context, in *pb.TrackOption) (*pb.TrackResult, error) {
	img, w, h, err := s.encodeImg(in.Image)
	if err != nil {
		return nil, err
	}
	req := url.Values{
		"case_id":   []string{strconv.FormatInt(in.CaseId, 10)},
		"case_init": []string{strconv.FormatBool(in.CaseInit)},
		"image":     []string{img},
		"dynamic":   []string{strconv.FormatBool(in.Dynamic)},
	}
	if in.Dynamic {
		req.Set("area", fmt.Sprintf("%v,%v,%v,%v,%v,%v,%v,%v", "0", "0", w, "0", w, -h, "0", -h))
	}
	var ret struct {
		core.BaiduResponse
		pb.TrackResult
	}
	if err := s.PostForm(biz.URLBodyTracking, req, &ret); err != nil {
		logrus.Errorf("track failed, %s", err)
		return nil, err
	}
	if ret.ErrorCode > 0 {
		return nil, fmt.Errorf("track failed, (%v) %s", ret.ErrorCode, ret.ErrorMsg)
	}

	return &ret.TrackResult, nil
}

func (s *BodyTrackServer) encodeImg(imgURL string) (encodeRet string, width, height int, err error) {
	file, err := os.Open(imgURL)
	if err != nil {
		return "", 0, 0, err
	}
	defer file.Close()
	buf, err := ioutil.ReadAll(file)
	if err != nil {
		return "", 0, 0, err
	}
	temp := ioutil.NopCloser(bytes.NewBuffer(buf))

	img, _, err := image.DecodeConfig(temp)
	if err != nil {
		return "", 0, 0, err
	}
	width = img.Width
	height = img.Height

	encodeRet = base64.StdEncoding.EncodeToString(buf)
	return
}
