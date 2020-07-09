package nlp

import (
	"github.com/funxdata/baidu/core"
)

// 自然语言处理

const (
	urlNLPLexer  = "/rpc/2.0/nlp/v1/lexer"
	urlNLPEmpSim = "/rpc/2.0/nlp/v2/word_emb_sim"
)

// BaiduNLP 百度自然语言处理
type BaiduNLP struct {
	*core.Core
}

// LexerRequest 词法分析请求
type LexerRequest struct {
	Text string `json:"text"`
}

// LexerResult 词法分析结果
type LexerResult struct {
	core.BaiduResponse
	Text  string      `json:"text"`
	Items []LexerItem `json:"items"`
}

// LexerItem 每个元素对应结果中的一个词
type LexerItem struct {
	// Item 词汇的字符串
	Item string `json:"item"`
	// Length 字节级length
	Length int `json:"byte_length"`
	// Offset 在text中的字节级offset
	Offset int `json:"byte_offset"`
	// Formal 词汇的标准化表达，主要针对时间、数字单位，没有归一化表达的，此项为空串
	Formal string `json:"formal"`
	// NE 命名实体类型，命名实体识别算法使用。词性标注算法中，此项为空串
	// PER 人名
	// LOC 地名
	// ORG 机构名
	// TIME 时间
	Ne string `json:"ne"`
	// Pos 词性，词性标注算法使用。命名实体识别算法中，此项为空串
	// n	普通名词	f	方位名词	s	处所名词	t	时间名词
	// nr	人名	ns	地名	nt	机构团体名	nw	作品名
	// nz	其他专名	v	普通动词	vd	动副词	vn	名动词
	// a	形容词	ad	副形词	an	名形词	d	副词
	// m	数量词	q	量词	r	代词	p	介词
	// c	连词	u	助词	xc	其他虚词	w	标点符号
	Pos string `json:"pos"`
	// URI 链指到知识库的URI，只对命名实体有效。对于非命名实体和链接不到知识库的命名实体，此项为空串
	URI string `json:"uri"`
	// 地址成分，非必需，仅对地址型命名实体有效，没有地址成分的，此项为空数组。
	LocDetails []string `json:"loc_details"`
	// BasicWords 基本词成分
	BasicWords []string `json:"basic_words"`
}

// Lexer 词法分析
// 官方文档 https://cloud.baidu.com/doc/NLP/s/fk6z52f2u
func (b *BaiduNLP) Lexer(req *LexerRequest) (*LexerResult, error) {
	ret := &LexerResult{}

	err := b.Post(urlNLPLexer, req, ret)
	if err != nil {
		return nil, err
	}

	return ret, nil
}

// CompareWords 词语比较参数
type CompareWords struct {
	Word1 string `json:"word_1"`
	Word2 string `json:"word_2"`
}

// CompareSimilarityResult 词义相似度比较结果
type CompareSimilarityResult struct {
	Words CompareWords `json:"words"`
	// Score 相似度结果，(0,1]，分数越高说明相似度越高
	Score float64 `json:"score"`
}

// CompareSimilarity 词义相似度
// 官方文档 https://cloud.baidu.com/doc/NLP/s/Fk6z52fjc
func (b *BaiduNLP) CompareSimilarity(req *CompareWords) (*CompareSimilarityResult, error) {
	ret := &CompareSimilarityResult{}

	err := b.Post(urlNLPEmpSim, req, ret)
	if err != nil {
		return nil, err
	}

	return ret, nil
}
