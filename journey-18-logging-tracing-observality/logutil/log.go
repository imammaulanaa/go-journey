package logutil

import (
	"github.com/sirupsen/logrus"
)

var log = logrus.New()

func LogInfo(message string) {
	log.WithField("level", "info").Info(message)
}

func LogError(message string) {
	log.WithField("level", "error").Error(message)
}
