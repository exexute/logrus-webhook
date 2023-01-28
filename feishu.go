package logrusWebhook

import (
	"github.com/sirupsen/logrus"
)

const (
	TextMsg = "text"
)

type FeiShuHook struct {
	Writer    *FeiShuWriter
	LogLevels []logrus.Level
}

func NewFeiShuHook(url, sign string, logLevels ...logrus.Level) (*FeiShuHook, error) {
	w, err := RegisterFeiShuWriter(url, sign)
	return &FeiShuHook{w, logLevels}, err
}

func (hook *FeiShuHook) Levels() []logrus.Level {
	return hook.LogLevels
}

func (hook *FeiShuHook) Fire(e *logrus.Entry) (err error) {
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
