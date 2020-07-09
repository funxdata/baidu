package speech

import (
	"testing"

	"github.com/funxdata/baidu/core"
	"github.com/stretchr/testify/assert"
)

var (
	skip   = true
	testSP = &BaiduSpeech{
		Core: core.NewCore("xxxxx", "yyyyyy"),
	}
)

func TestSpeechRecognitionFile(t *testing.T) {
	if skip {
		return
	}
	fp := "./testdata/8k.wav"
	ret, err := testSP.SpeechRecognitionFile("wav", 8000, fp)
	assert.Nil(t, err)

	t.Logf("%+v", ret)
}
