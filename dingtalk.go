package logrusWebhook

import "github.com/sirupsen/logrus"

type DingTalkHook struct {
	Writer    *DingTalkWriter
	LogLevels []logrus.Level
}

func NewDingTalkHook(openApi, token, secret string, logLevels ...logrus.Level) (*DingTalkHook, error) {
	w, err := RegisterDingTalkWriter(openApi, token, secret)
	return &DingTalkHook{w, logLevels}, err
}

func (hook *DingTalkHook) Levels() []logrus.Level {
	return hook.LogLevels
}

func (hook *DingTalkHook) Fire(e *logrus.Entry) (err error) {
	var at *DingTalkAt
	atValue, _ := e.Data["at"]
	if atValue != nil {
		at = atValue.(*DingTalkAt)
	}

	msgType, _ := e.Data["msgType"]
	switch msgType {
	case "markdown":
		err = hook.Writer.WriteMarkdownMsg(e.Message, at)
		break
	case "link":
		err = hook.Writer.WriteLinkMsg(e.Message, at)
		break
	default:
		err = hook.Writer.WriteTextMsg(e.Message, at)
		break
	}

	return
}
