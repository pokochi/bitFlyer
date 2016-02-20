package bitFlyer

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"
)

const (
	GET_METHOD  = "GET"
	POST_METHOD = "POST"
	END_POINT   = "https://api.bitflyer.jp"
)

type Lightning struct {
	AccessKey    string
	AccessSecret string
}

func (l *Lightning) Request(method string, path string, params map[string]string) (err error, s string) {
	if method != GET_METHOD {
		return errors.New("Request Error"), ""
	}

	req, _ := http.NewRequest(method, END_POINT+path, nil)
	if params != nil {
		values := url.Values{}
		for k, v := range params {
			values.Add(k, v)
		}
		req.URL.RawQuery = values.Encode()
	}

	client := new(http.Client)
	resp, _ := client.Do(req)

	defer resp.Body.Close()
	byteArray, _ := ioutil.ReadAll(resp.Body)

	return nil, string(byteArray)
}

func (l *Lightning) RequestPrivate(method string, path string, params map[string]string) (err error, s string) {
	if method != GET_METHOD && method != POST_METHOD {
		return errors.New("RequestPrivate Error"), ""
	}

	values := url.Values{}
	if params != nil {
		for k, v := range params {
			values.Add(k, v)
		}
	}

	req, err := http.NewRequest(method, END_POINT+path, strings.NewReader(values.Encode()))
	if err != nil {
		return errors.New("RequestPrivate NewRequest Error"), ""
	}

	t := strconv.FormatInt(time.Now().Unix(), 10)
	sign := l.CreateSign(method, path, t)

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("ACCESS-KEY", l.AccessKey)
	req.Header.Set("ACCESS-TIMESTAMP", t)
	req.Header.Set("ACCESS-SIGN", sign)

	client := new(http.Client)
	resp, _ := client.Do(req)
	defer resp.Body.Close()

	byteArray, _ := ioutil.ReadAll(resp.Body)
	return nil, string(byteArray)
}

func (l *Lightning) CreateSign(method string, path string, t string) (s string) {
	text := t + method + path

	// API secretでSHA256 署名を生成する
	hash := hmac.New(sha256.New, []byte(l.AccessSecret))
	hash.Write([]byte(text))
	s = hex.EncodeToString(hash.Sum(nil))

	return s
}
