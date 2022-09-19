package log

import (
	"os"

	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"gopkg.in/natefinch/lumberjack.v2"
)

var Logger *logrus.Logger

func SetLogger() {
	lum := &lumberjack.Logger{
		Filename:   "internal/log/log.log",
		MaxSize:    500,
		MaxBackups: 3,
		MaxAge:     28,
	}

	if viper.Get("production") == true {
		logrus.SetFormatter(&logrus.JSONFormatter{})
		logrus.SetLevel(logrus.WarnLevel)
		logrus.SetOutput(lum)
	} else {
		logrus.SetFormatter(&logrus.TextFormatter{})
		logrus.SetLevel(logrus.InfoLevel)
		logrus.SetOutput(os.Stdout)
	}
	logrus.SetReportCaller(true)
	Logger = logrus.New()
}
