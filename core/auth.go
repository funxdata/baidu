package core

import (
	"encoding/json"
	"net/http"
	"net/url"

	"github.com/sirupsen/logrus"
)

const (
	urlAuth = "https://aip.baidubce.com/oauth/2.0/token"
)

type AccessToken struct {
	AccessToken   string `json:"access_token"`
	SessionKey    string `json:"session_key"`
	Scope         string `json:"scope"`
	RefreshToken  string `json:"refresh_token"`
	SessionSecret string `json:"session_secret"`
	ExpiresIn     int64  `json:"expires_in"`
}

// AccessToken
func (b *Core) GetAccessToken() (string, error) {
	if b.tkn != nil {
		return b.tkn.AccessToken, nil
	}
	u, _ := url.Parse(urlAuth)
	q := u.Query()
	q.Set("grant_type", "client_credentials")
	q.Set("client_id", b.ApiKey)
	q.Set("client_secret", b.ApiSecret)
	u.RawQuery = q.Encode()
	resp, err := http.Get(u.String())
	if err != nil {
		logrus.Errorf("got access token failed, %s", err)
		return "", err
	}

	tkn := &AccessToken{}
	err = json.NewDecoder(resp.Body).Decode(tkn)
	if err != nil {
		logrus.Errorf("decode access token json failed, %s", err)
		return "", err
	}

	b.tkn = tkn
	return tkn.AccessToken, nil
}
