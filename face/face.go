package face

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"image"
	"image/jpeg"
	"strings"

	"github.com/funxdata/baidu/core"
	"github.com/sirupsen/logrus"
)

const (
	// h5语音验证码
	urlFacesessioncode = "/rest/2.0/face/v1/faceliveness/sessioncode"
	// h5活体视频分析
	urlFaceVerify = "/rest/2.0/face/v1/faceliveness/verify"
	// 人脸检测
	urlFaceDetect = "/rest/2.0/face/v3/detect"
	// 人脸对比
	urlFaceMatch = "/rest/2.0/face/v3/match"
	// 人脸搜索
	urlFaceSearch = "/rest/2.0/face/v3/search"
	// 人脸库管理-人脸注册
	urlFaceUserAdd = "/rest/2.0/face/v3/faceset/user/add"
	// 人脸库管理-人脸更新
	urlFaceUserUpdate = "/rest/2.0/face/v3/faceset/user/update"
	// 人脸库管理-删除用户
	urlFaceUserDelete = "/rest/2.0/face/v3/faceset/user/delete"
	// 人脸库管理-用户信息查询
	urlFaceUserGet = "/rest/2.0/face/v3/faceset/user/get"
	// 人脸库管理-获取组列表
	urlFaceGroups = "/rest/2.0/face/v3/faceset/group/getlist"
	// 人脸库管理-获取用户列表
	urlFaceGroupUsers = "/rest/2.0/face/v3/faceset/group/getusers"
	// 人脸库管理-复制用户
	urlFaceCopy = "/rest/2.0/face/v3/faceset/user/copy"
	// 人脸库管理-获取用户人脸列表
	urlFaces = "/rest/2.0/face/v3/faceset/face/getlist"
	// 人脸库管理-创建用户组
	urlFaceGroupAdd = "/rest/2.0/face/v3/faceset/group/add"
	// 人脸库管理-删除用户组
	urlFaceGroupDelete = "/rest/2.0/face/v3/faceset/group/delete"
	// 人脸库管理-删除人脸
	urlFaceDelete = "/rest/2.0/face/v3/faceset/face/delete"
	// 在线活体检测
	urlFaceFaceverify = "/rest/2.0/face/v3/faceverify"
)

type BaiduFace struct {
	*core.Core
}

type SearchResult struct {
	FaceToken string          `json:"face_token"`
	UserList  []BaiduFaceUser `json:"user_list"`
}

type BaiduFaceUser struct {
	GroupId  string  `json:"group_id"`
	UserId   string  `json:"user_id"`
	UserInfo string  `json:"user_info"`
	Score    float64 `json:"score"`
}

// SearchFace
func (b *BaiduFace) Search(imgurl string, groups ...string) (*SearchResult, error) {
	req := map[string]interface{}{
		"image":      imgurl,
		"image_type": "FACE_TOKEN",
		// "image_type":    "URL",
		"group_id_list": strings.Join(groups, ","),
	}
	var ret struct {
		core.BaiduResponse `json",inline"`
		SearchResult       `json:"result"`
	}
	if err := b.Post(urlFaceSearch, req, &ret); err != nil {
		logrus.Errorf("searched failed, %s", err)
		return nil, err
	}

	if err := ret.Err("search"); err != nil {
		return nil, err
	}

	return &ret.SearchResult, nil
}

type DetectResult struct {
	FaceList []FaceInfo `json:"face_list"`
}
type InfoType struct {
	Probability float64 `json:"probability"`
	Type        string  `json:"type"`
}
type Point struct {
	X float64 `json:"x"`
	Y float64 `json:"y"`
}
type FaceInfo struct {
	FaceToken string `json:"face_token"`
	Age       int    `json:"age"`
	Angle     struct {
		Pitch float64 `json:"pitch"`
		Roll  float64 `json:"roll"`
		Yaw   float64 `json:"yaw"`
	} `json:"angle"`
	Beauty          float64
	Expression      InfoType `json:"expression"`
	FaceProbability float64  `json:"face_probability"`
	FaceShape       InfoType `json:"face_shape"`
	FaceType        InfoType `json:"face_type"`
	Gender          InfoType `json:"gender"`
	Glasses         InfoType `json:"glasses"`
	Race            InfoType `json:"race"`
	Landmark        []Point  `json:"landmark"`
	Landmark72      []Point  `json:"landmark72"`
	Location        struct {
		Height   int     `json:"height"`
		Width    int     `json:"width"`
		Rotation float64 `json:"rotation"`
		Left     float64 `json:"left"`
		Top      float64 `json:"top"`
	} `json:"location"`
}

func (b *BaiduFace) Detect(imgurl string) (*DetectResult, error) {
	req := map[string]interface{}{
		"image":        imgurl,
		"image_type":   "URL",
		"face_field":   "age,beauty,expression,face_shape,gender,glasses,landmark,race,quality,face_type",
		"max_face_num": 10,
		// "face_type":    "",
	}
	var ret struct {
		core.BaiduResponse `json",inline"`
		DetectResult       `json:"result"`
	}
	err := b.Post(urlFaceDetect, req, &ret)
	if err != nil {
		logrus.Errorf("detect face %s failed, %s", imgurl, err)
		return nil, err
	}

	if ret.ErrorCode > 0 {
		return nil, fmt.Errorf("detect failed, (%v) %s", ret.ErrorCode, ret.ErrorMsg)
	}

	return &ret.DetectResult, nil
}

// img
func (b *BaiduFace) img(obj interface{}) map[string]interface{} {
	if imgurl, ok := obj.(string); ok {
		return map[string]interface{}{
			"image":      imgurl,
			"image_type": "URL",
		}
	}

	if img, ok := obj.(image.Image); ok {
		buf, out := new(bytes.Buffer), new(bytes.Buffer)
		jpeg.Encode(buf, img, nil)
		base64.NewEncoder(base64.StdEncoding, out)
		return map[string]interface{}{
			"image":      out.String(),
			"image_type": "FACE_TOKEN",
		}
	}
	return nil
}
