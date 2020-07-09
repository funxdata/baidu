package speech

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"io"
	"net/http"
	"os"

	"github.com/funxdata/baidu/core"
)

const (
	urlServerAPI = "/server_api"
)

// BaiduSpeech 百度语言
type BaiduSpeech struct {
	*core.Core
}

type RecognitionRequest struct {
	// Format 语音文件的格式，pcm/wav/amr/m4a。不区分大小写。推荐pcm文件
	Format string `json:"format"`
	// Rate 采样率，16000、8000，固定值
	Rate int `json:"rate"`
	// Channel 声道数，仅支持单声道，请填写固定值 1
	Channel int `json:"channel"`
	// CUID 用户唯一标识，用来区分用户，计算UV值。建议填写能区分用户的机器 MAC 地址或 IMEI 码，长度为60字符以内。
	CUID string `json:"cuid"`
	// Token 开放平台获取到的开发者[access_token]获取 Access Token "access_token")
	Token string `json:"token"`
	// Speech 本地语音文件的的二进制语音数据 ，需要进行base64 编码。与len参数连一起使用。
	Speech string `json:"speech"`
	// Len 本地语音文件的的字节数，单位字节
	Len int64 `json:"len"`

	// DevPID 不填写lan参数生效，都不填写，默认1537（普通话 输入法模型），dev_pid参数见本节开头的表格
	DevPID int `json:"dev_pid"`
	// LmID 自训练平台模型id，填dev_pid = 8001 或 8002生效
	LmID int `json:"lm_id"`
}

type RecognitionResult struct {
	ErrNo    int    `json:"err_no"`
	ErrMsg   string `json:"err_msg"`
	CorpusNo string `json:"corpus_no"`
	SN       string `json:"sn"`

	Result []string `json:"result"`
}

// SpeechRecognition 语音识别
func (v *BaiduSpeech) SpeechRecognition(in *RecognitionRequest) (*RecognitionResult, error) {
	var (
		err error
		buf = new(bytes.Buffer)
	)
	in.Channel = 1
	in.CUID = "baidu_workshop"
	in.DevPID = 1537
	in.Token, err = v.GetAccessToken()
	if err != nil {
		return nil, err
	}

	json.NewEncoder(buf).Encode(in)

	req, err := http.NewRequest("POST", "http://vop.baidu.com/server_api", buf)
	if err != nil {
		return nil, err
	}

	ret := &RecognitionResult{}
	if err := v.Do(req, ret); err != nil {
		return nil, err
	}
	return ret, nil
}

func (v *BaiduSpeech) SpeechRecognitionReador(format string, rate int, r io.Reader) (*RecognitionResult, error) {
	var (
		req = &RecognitionRequest{
			Format: format,
			Rate:   rate,
		}
		speechBuf = new(bytes.Buffer)
		err       error
	)

	req.Len, err = io.Copy(speechBuf, r)
	if err != nil {
		return nil, err
	}

	req.Speech = base64.StdEncoding.EncodeToString(speechBuf.Bytes())
	return v.SpeechRecognition(req)
}

func (v *BaiduSpeech) SpeechRecognitionFile(format string, rate int, filepath string) (*RecognitionResult, error) {
	f, err := os.Open(filepath)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	return v.SpeechRecognitionReador(format, rate, f)
}
