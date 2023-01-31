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
	"strconv"
	"sync"
	"time"
)

type FeiShuWriter struct {
	Url  string
	Sign string

	mu sync.Mutex
}

func RegisterFeiShuWriter(url, sign string) (*FeiShuWriter, error) {
	w := &FeiShuWriter{Url: url, Sign: sign}

	w.mu.Lock()
	defer w.mu.Unlock()

	err := w.WriteTextMsg("[logrus.webhook.feishu] register logrus hook")
	if err != nil {
		return nil, err
	}
	return w, nil
}

func (w *FeiShuWriter) Write(msg []byte) error {
	postData := msg
	if w.supportSign() {
		timestamp := time.Now().UnixNano() / 1e6
		sign, err := GenSign(w.Sign, timestamp)
		//sign := calcSign(timestamp, w.Sign)

		signedMsg, err := convertMsg(msg)
		if err != nil {
			return err
		}
		switch signedMsg.MsgType {
		case TextMsg:
			textMsg, err := convertTextMsg(msg)
			if err != nil {
				return err
			}
			textMsg.Sign = sign
			textMsg.Timestamp = strconv.FormatInt(timestamp, 10)
			postData, _ = json.Marshal(textMsg)
		}
	}

	req, err := http.NewRequest("POST", w.Url, bytes.NewBuffer(msg))
	if err != nil {
		logrus.Errorf("[logrus.webhook.feishu] create requrest failure, msg: %s, error msg: %s", string(postData), err.Error())
		return err
	}
	req.Header.Set("Content-Type", "application/json")

	hc := http.Client{}
	resp, err := hc.Do(req)
	if err != nil {
		logrus.Errorf("[logrus.webhook.feishu] call api failure, msg: %s, error msg: %s", string(postData), err.Error())
		return err
	}

	defer resp.Body.Close()
	data, _ := ioutil.ReadAll(resp.Body)
	apiResp := &FeiShuResponse{}
	err = json.Unmarshal(data, apiResp)
	if err != nil {
		logrus.Errorf("[logrus.webhook.feishu] unmarshal api resp failure, msg: %s, api resp msg: %s, error msg: %s", string(postData), string(data), apiResp.Msg)
		return err
	}
	if apiResp.Code != 0 {
		logrus.Errorf("[logrus.webhook.feishu] send notify msg failure, msg: %s, error code: %d, error msg: %s", string(postData), apiResp.Code, apiResp.Msg)
	}
	return nil
}

func (w *FeiShuWriter) WriteTextMsg(content string) error {
	msg := &FeiShuTextMsg{}
	msg.MsgType = TextMsg
	msg.Content.Text = content

	postData, _ := json.Marshal(msg)
	return w.Write(postData)
}

func (w *FeiShuWriter) supportSign() bool {
	return w.Sign != ""
}

func convertMsg(msg []byte) (*FeiShuMsg, error) {
	signedMsg := &FeiShuMsg{}
	err := json.Unmarshal(msg, signedMsg)
	if err != nil {
		return nil, err
	}
	return signedMsg, nil
}

func convertTextMsg(msg []byte) (*FeiShuTextMsg, error) {
	signedMsg := &FeiShuTextMsg{}
	err := json.Unmarshal(msg, signedMsg)
	if err != nil {
		return nil, err
	}
	return signedMsg, nil
}

func GenSign(secret string, timestamp int64) (string, error) {
	//timestamp + key 做sha256, 再进行base64 encode
	stringToSign := fmt.Sprintf("%v", timestamp) + "\n" + secret

	var data []byte
	h := hmac.New(sha256.New, []byte(stringToSign))
	_, err := h.Write(data)
	if err != nil {
		return "", err
	}

	signature := base64.StdEncoding.EncodeToString(h.Sum(nil))
	return signature, nil
}
