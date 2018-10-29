package body

import (
	"flag"
	"os"
	"testing"

	"github.com/sirupsen/logrus"

	"github.com/funxdata/baidu/core"
)

func Test_body(t *testing.T) {
	flag.Set("alsologtostderr", "true")
	tb := &BaiduBody{Core: core.NewCore(os.Getenv("BAIDU_KEY"), os.Getenv("BAIDU_SECRET"))}
	resp, err := tb.Track("~/Pictures/test.jpg")
	if err != nil {
		logrus.Info(err)
		return
	}
	logrus.Info(resp.PersonNum)
}
