package logrusWebhook

import (
	"bytes"
	"encoding/json"
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

func (w *FeiShuWriter) Write(msg interface{}) error {
	if w.supportSign() {
		timestamp := time.Now().UnixNano() / 1e6
		sign := calcSign(timestamp, w.Sign)
		msg.(*FeiShuMsg).Timestamp = strconv.FormatInt(timestamp, 10)
		msg.(*FeiShuMsg).Sign = sign
	}

	postData, _ := json.Marshal(msg)
	req, err := http.NewRequest("POST", w.Url, bytes.NewBuffer(postData))
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
	err = json.Unmarshal(data, resp)
	if err != nil {
		logrus.Errorf("[logrus.webhook.feishu] unmarshal api resp failure, msg: %s, api resp msg: %s, error msg: %s", string(postData), string(data), apiResp.Msg)
		return err
	}
	if apiResp.Code != 0 {
		logrus.Errorf("[logrus.webhook.feishu] send notify msg failure, msg: %s, error msg: %s", string(postData), apiResp.Msg)
	}
	return nil
}

func (w *FeiShuWriter) WriteTextMsg(content string) error {
	msg := &FeiShuTextMsg{}
	msg.MsgType = "text"
	msg.Content.Text = content

	return w.Write(msg)
}

func (w *FeiShuWriter) supportSign() bool {
	return w.Sign != ""
}
