package logger

import "github.com/sirupsen/logrus"

var (
	log *logrus.Logger
)

func IntitLogger() {
	log = logrus.StandardLogger()
}

func GetLogger() *logrus.Logger {
	if log == nil {
		IntitLogger()
	}

	return log
}
