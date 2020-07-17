package ocr

import (
	"encoding/base64"
	"io"
	"io/ioutil"
	"net/url"

	"github.com/funxdata/baidu/core"
)

const (
	urlOCRAccurate = "/rest/2.0/ocr/v1/accurate"
)

// BaiduOCR 百度自然语言处理
type BaiduOCR struct {
	*core.Core
}

type AccurateRequest struct {
	Image string `json:"image"`
}

// AccurateResult .
type AccurateResult struct {
	core.BaiduResponse

	WordsResultNum int `json:"words_result_num"`
	WordsResult    []struct {
		Location struct {
			Width  int `json:"width"`
			Top    int `json:"top"`
			Height int `json:"height"`
			Left   int `json:"left"`
		} `json:"location"`

		Words string `json:"words"`
	} `json:"words_result"`
}

// Accurate 通用文字识别（高精度含位置版）
// https://cloud.baidu.com/doc/OCR/s/tk3h7y2aq
func (o *BaiduOCR) Accurate(r io.Reader) (*AccurateResult, error) {
	data, err := ioutil.ReadAll(r)
	if err != nil {
		return nil, err
	}
	body := make(url.Values)
	body.Set("image", base64.StdEncoding.EncodeToString(data))
	ret := &AccurateResult{}

	err = o.PostForm(urlOCRAccurate, body, &ret)
	if err != nil {
		return nil, err
	}

	if err := ret.Err(); err != nil {
		return nil, err
	}

	return ret, nil
}
