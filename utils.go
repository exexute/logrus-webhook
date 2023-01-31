package logrusWebhook

import "github.com/sirupsen/logrus"

func getLevels(minLevel logrus.Level) (levels []logrus.Level) {
	for _, level := range logrus.AllLevels {
		if level <= minLevel {
			levels = append(levels, level)
		}
	}
	return
}
