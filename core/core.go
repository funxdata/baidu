package core

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strings"

	"github.com/sirupsen/logrus"
)

type Core struct {
	AppId     string
	ApiKey    string
	ApiSecret string

	tkn     *AccessToken
	baseURL *url.URL
}

// NewBaidu
func NewCore(apiKey, apiSecret string, baseURL ...string) *Core {
	b := &Core{
		ApiKey:    apiKey,
		ApiSecret: apiSecret,
	}
	if len(baseURL) == 1 && baseURL[0] != "" {
		b.baseURL, _ = url.Parse(baseURL[0])
	} else {
		b.baseURL, _ = url.Parse("https://aip.baidubce.com/")
	}
	return b
}

// URL
func (b *Core) URL(path string) *url.URL {
	u, _ := url.Parse(b.baseURL.String())
	if path != "" {
		u.Path = path
	}
	return u
}

// Post
func (b *Core) Post(path string, object interface{}, ret interface{}) error {
	tkn, err := b.GetAccessToken()
	if err != nil {
		return err
	}

	buf := new(bytes.Buffer)
	json.NewEncoder(buf).Encode(object)
	u := b.URL(path)
	q := u.Query()
	q.Set("access_token", tkn)
	q.Set("charset", "UTF-8")
	u.RawQuery = q.Encode()

	req, err := http.NewRequest("POST", u.String(), buf)
	if err != nil {
		return err
	}

	req.Header.Set("Content-Type", "application/json")

	return b.Do(req, ret)
	// resp, err := http.Post(u.String(), "application/json", buf)
	// if err != nil {
	// 	return err
	// }

	// // bs, _ := ioutil.ReadAll(resp.Body)
	// // logrus.Infof("body: %s", bs)

	// err = json.NewDecoder(resp.Body).Decode(ret)
	// if err != nil {
	// 	logrus.Errorf("decode body to %T failed, %s", ret, err)
	// 	return err
	// }
	// return nil
}

// Do .
func (b *Core) Do(req *http.Request, ret interface{}) error {
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}

	// bs, _ := ioutil.ReadAll(resp.Body)
	// logrus.Infof("body: %s", bs)

	err = json.NewDecoder(resp.Body).Decode(ret)
	if err != nil {
		logrus.Errorf("decode body to %T failed, %s", ret, err)
		return err
	}
	return nil
}

//post form
func (b *Core) PostForm(path string, form url.Values, ret interface{}) error {
	tkn, err := b.GetAccessToken()
	if err != nil {
		return err
	}

	u := b.URL(path)
	q := u.Query()
	q.Set("access_token", tkn)
	q.Set("charset", "UTF-8")
	u.RawQuery = q.Encode()
	resp, err := http.Post(u.String(),
		"application/x-www-form-urlencoded", strings.NewReader(form.Encode()))

	if err != nil {
		return err
	}
	err = json.NewDecoder(resp.Body).Decode(ret)
	if err != nil {
		logrus.Errorf("decode body to %T failed, %s", ret, err)
		return err
	}
	return nil
}

type BaiduResponse struct {
	ErrorCode int    `json:"error_code"`
	ErrorMsg  string `json:"error_msg"`
	LogId     int64  `json:"log_id"`
	Timestamp int64  `json:"timestamp"`
}

// Err
func (b BaiduResponse) Err(action ...string) error {
	if b.ErrorCode == 0 {
		return nil
	}
	if len(action) > 0 {
		return fmt.Errorf("%s failed, (%v) %s", strings.Join(action, " "), b.ErrorCode, b.ErrorMsg)
	}
	return fmt.Errorf("(%v) %s", b.ErrorCode, b.ErrorMsg)
}
