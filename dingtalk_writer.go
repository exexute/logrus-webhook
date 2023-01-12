package logrusWebhook

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"github.com/sirupsen/logrus"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
	"sync"
	"time"
)

const (
	DefaultDingTalkApi = "https://oapi.dingtalk.com"
	UrlOfRobotSend     = "robot/send"
)

type DingTalkWriter struct {
	OpenApi   string
	Token     string
	Secret    string
	NotifyUrl string

	mu sync.Mutex
}

func RegisterDingTalkWriter(openApi, token, secret string) (*DingTalkWriter, error) {
	w := &DingTalkWriter{
		OpenApi:   openApi,
		Token:     token,
		Secret:    secret,
		NotifyUrl: urlJoin(openApi, UrlOfRobotSend),
	}

	w.mu.Lock()
	defer w.mu.Unlock()

	err := w.WriteTextMsg(fmt.Sprintf("[logrus.webhook.dingtalk] register logrus hook with %s", w.OpenApi), nil)
	if err != nil {
		return nil, err
	}
	return w, nil
}

func (w *DingTalkWriter) Write(msg interface{}) error {
	form := url.Values{
		"access_token": []string{w.Token},
	}
	if w.supportSign() {
		timestamp := time.Now().UnixNano()
		sign := calcSign(timestamp, w.Secret)

		form["timestamp"] = []string{strconv.FormatInt(timestamp, 10)}
		form["sign"] = []string{sign}
	}

	postData, _ := json.Marshal(msg)
	req, err := http.NewRequest("POST", w.NotifyUrl, bytes.NewBuffer(postData))
	if err != nil {
		logrus.Errorf("[logrus.webhook.dingtalk] create requrest failure, msg: %s, error msg: %s", string(postData), err.Error())
		return err
	}
	req.URL.RawQuery = form.Encode()
	req.Header.Set("Content-Type", "application/json")

	hc := http.Client{}
	resp, err := hc.Do(req)
	if err != nil {
		logrus.Errorf("[logrus.webhook.dingtalk] call api failure, msg: %s, error msg: %s", string(postData), err.Error())
		return err
	}

	defer resp.Body.Close()
	data, _ := ioutil.ReadAll(resp.Body)
	apiResp := &DingTalkResponse{}
	err = json.Unmarshal(data, resp)
	if err != nil {
		logrus.Errorf("[logrus.webhook.dingtalk] unmarshal api resp failure, msg: %s, api resp msg: %s, error msg: %s", string(postData), string(data), apiResp.ErrMsg)
		return err
	}
	if apiResp.ErrCode != 0 {
		logrus.Errorf("[logrus.webhook.dingtalk] send notify msg failure, msg: %s, error msg: %s", string(postData), apiResp.ErrMsg)
	}
	return nil
}

func (w *DingTalkWriter) supportSign() bool {
	return w.Secret != ""
}

func (w *DingTalkWriter) WriteTextMsg(content string, at *DingTalkAt) error {
	msg := DingTalkTextMsg{}
	msg.MsgType = "text"
	msg.Text.Content = content
	msg.At = at

	return w.Write(msg)
}

func (w *DingTalkWriter) WriteLinkMsg(content string, at *DingTalkAt) error {
	link, err := UnmarshalDingTalkLink(content)
	if err != nil {
		return err
	}
	msg := DingTalkLinkMsg{}
	msg.MsgType = "link"
	msg.At = at
	msg.Link = link

	return w.Write(msg)
}

func (w *DingTalkWriter) WriteMarkdownMsg(content string, at *DingTalkAt) error {
	markdown, err := UnmarshalDingTalkMarkdown(content)
	if err != nil {
		return err
	}
	msg := DingTalkMarkdownMsg{}
	msg.MsgType = "markdown"
	msg.At = at
	msg.Markdown = markdown

	return w.Write(msg)
}

func calcSign(timestamp int64, secret string) (sign string) {
	stringToSign := fmt.Sprintf("%d\n%s", timestamp, secret)
	mac := hmac.New(sha256.New, []byte(secret))
	mac.Write([]byte(stringToSign))
	expectedMac := mac.Sum(nil)
	return base64.StdEncoding.EncodeToString(expectedMac)
}

func urlJoin(openApi, path string) (fullUrl string) {
	u, err := url.Parse(openApi)
	if err != nil {
		logrus.Error("[logrus.webhook.dingtalk] urljoin failure, openApi: %s, path: %s", openApi, path)
		return fmt.Sprintf("%s%s", openApi, path)
	}
	u.Path = path
	fullUrl, err = url.PathUnescape(u.String())
	if err != nil {
		logrus.Error("[logrus.webhook.dingtalk] urljoin failure when un-escape fullUrl, openApi: %s, path: %s", openApi, path)
		return fmt.Sprintf("%s%s", openApi, path)
	}
	return
}
