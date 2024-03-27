package logger

import (
	"time"

	"github.com/sirupsen/logrus"
)

var logger = logrus.New()

func setupJsonLogging() {
	logger.SetFormatter(&logrus.JSONFormatter{
		TimestampFormat: time.RFC3339Nano,
	})

	logger.SetReportCaller(true)
}

func setupTextLogging() {
	logger.SetFormatter(&logrus.TextFormatter{
		FullTimestamp: true,
	})
	logger.SetReportCaller(true)

}
func setFormat(format string) {
	switch format {
	case "json":
		setupJsonLogging()
	default:
		setupTextLogging()
	}
}

func setLevel(level string) {
	lvl, err := logrus.ParseLevel(level)
	if err != nil {
		logger.Fatal(err)
	}
	logger.SetLevel(lvl)
}

func InitLogger(loggingFormat, loggingLevel string) *logrus.Entry {
	setFormat(loggingFormat)
	setLevel(loggingLevel)
	return GetLogger()
}

func GetLogger() *logrus.Entry {
	return logrus.NewEntry(logger)
}
