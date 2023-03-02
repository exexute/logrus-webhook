package logrusWebhook

import "github.com/sirupsen/logrus"

const EnableSlsLog = "slslog"

type SlsHook struct {
	Writer   *SlsWriter
	LogLevel logrus.Level
}

func NewSlsHook(config *SlsConfig, logLevel logrus.Level) (*SlsHook, error) {
	w, err := RegisterSlsWriter(config)
	return &SlsHook{w, logLevel}, err
}

func (hook *SlsHook) Levels() (levels []logrus.Level) {
	return getLevels(hook.LogLevel)
}

func (hook *SlsHook) Fire(e *logrus.Entry) (err error) {
	if enable, hasKey := e.Data[EnableSlsLog]; !(hasKey && enable) {
		return
	}
	err = hook.Writer.Write(e.Message)
	return
}
