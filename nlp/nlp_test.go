package nlp

import (
	"testing"

	"github.com/funxdata/baidu/core"
	"github.com/stretchr/testify/assert"
)

var (
	skip    = true
	testNLP = &BaiduNLP{
		Core: core.NewCore("xxxxx", "yyyyy"),
	}
)

func TestLexer(t *testing.T) {
	if skip {
		return
	}
	req := &LexerRequest{
		Text: "周恩来是一位伟大的中国共产党员",
	}
	ret, err := testNLP.Lexer(req)
	assert.Nil(t, err)
	for _, v := range ret.Items {
		t.Logf("%+v", v)
	}

}

func TestCompareSimilarity(t *testing.T) {
	if skip {
		return
	}
	req := &CompareWords{
		Word1: "北京",
		Word2: "上海",
	}
	ret, err := testNLP.CompareSimilarity(req)
	assert.Nil(t, err)

	t.Logf("%+v", ret)
	t.Error("..")
}
