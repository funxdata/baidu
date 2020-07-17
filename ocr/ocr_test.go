package ocr

import (
	"os"
	"testing"

	"github.com/funxdata/baidu/core"
	"github.com/stretchr/testify/assert"
)

var (
	skip    = true
	testNLP = &BaiduOCR{
		Core: core.NewCore("xxxxx", "yyyyy"),
	}
)

func TestAccurate(t *testing.T) {
	if skip {
		return
	}
	fpath := "/home/ckeyer/Pictures/bz2_20200713182710.jpg"
	f, err := os.Open(fpath)
	assert.Nil(t, err)

	ret, err := testNLP.Accurate(f)
	assert.Nil(t, err)

	t.Errorf("%+v", ret)
}
