package config

import (
	"os"

	log "github.com/sirupsen/logrus"
)

func LogInit(app AppConfig, logrus *log.Logger) {
	logLevel, _ := log.ParseLevel(app.LogLevel)
	logrus.SetLevel(logLevel)
	logrus.SetFormatter(&log.JSONFormatter{})
	logrus.SetReportCaller(true)
	logrus.SetOutput(os.Stdout)
}
