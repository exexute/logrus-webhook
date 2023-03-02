package logrusWebhook

import (
	"github.com/sirupsen/logrus"
)

const (
	TextMsg      = "text"
	EnableFeiShu = "feishu"
)

type FeiShuHook struct {
	Writer   *FeiShuWriter
	LogLevel logrus.Level
}

func NewFeiShuHook(url, sign string, logLevel logrus.Level) (*FeiShuHook, error) {
	w, err := RegisterFeiShuWriter(url, sign)
	return &FeiShuHook{w, logLevel}, err
}

func (hook *FeiShuHook) Levels() []logrus.Level {
	return getLevels(hook.LogLevel)
}

func (hook *FeiShuHook) Fire(e *logrus.Entry) (err error) {
	if enable, hasKey := e.Data[EnableFeiShu]; !(hasKey && enable.(bool)) {
		return
	}
	msgType, _ := e.Data["msgType"]
	switch msgType {
	//case "markdown":
	//	err = hook.Writer.WriteMarkdownMsg(e.Message, at)
	//	break
	//case "link":
	//	err = hook.Writer.WriteLinkMsg(e.Message, at)
	//	break
	default:
		err = hook.Writer.WriteTextMsg(e.Message)
		break
	}

	return
}
