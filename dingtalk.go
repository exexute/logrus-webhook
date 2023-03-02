package logrusWebhook

import "github.com/sirupsen/logrus"

const EnableDingTalk = "dingtalk"

type DingTalkHook struct {
	Writer   *DingTalkWriter
	LogLevel logrus.Level
}

func NewDingTalkHook(openApi, token, secret string, logLevel logrus.Level) (*DingTalkHook, error) {
	w, err := RegisterDingTalkWriter(openApi, token, secret)
	return &DingTalkHook{w, logLevel}, err
}

func (hook *DingTalkHook) Levels() []logrus.Level {
	return getLevels(hook.LogLevel)
}

func (hook *DingTalkHook) Fire(e *logrus.Entry) (err error) {
	if enable, hasKey := e.Data[EnableDingTalk]; !(enable && hasKey) {
		return
	}
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
