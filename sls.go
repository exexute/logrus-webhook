package logrusWebhook

import "github.com/sirupsen/logrus"

type SlsHook struct {
	Writer    *SlsWriter
	LogLevels []logrus.Level
}

func NewSlsHook(config *SlsConfig, logLevels ...logrus.Level) (*SlsHook, error) {
	w, err := RegisterSlsWriter(config)
	return &SlsHook{w, logLevels}, err
}

func (hook *SlsHook) Levels() []logrus.Level {
	return hook.LogLevels
}

func (hook *SlsHook) Fire(e *logrus.Entry) (err error) {
	if _, isOk := e.Data["slslog"]; !isOk {
		return
	}
	err = hook.Writer.Write(e.Message)
	return
}
