package baidu

import (
	"github.com/funxdata/baidu/core"
	"github.com/funxdata/baidu/face"
	"github.com/funxdata/baidu/nlp"
	"github.com/funxdata/baidu/ocr"
	"github.com/funxdata/baidu/speech"
)

type Baidu struct {
	*core.Core
}

func New(apiKey, apiSecret string) *Baidu {
	return &Baidu{core.NewCore(apiKey, apiSecret)}
}

// Face
func (b *Baidu) Face() *face.BaiduFace {
	return &face.BaiduFace{b.Core}
}

// Face
func (b *Baidu) Speech() *speech.BaiduSpeech {
	return &speech.BaiduSpeech{b.Core}
}

// NLP
func (b *Baidu) NLP() *nlp.BaiduNLP {
	return &nlp.BaiduNLP{b.Core}
}

// OCR .
func (b *Baidu) OCR() *ocr.BaiduOCR {
	return &ocr.BaiduOCR{b.Core}
}
