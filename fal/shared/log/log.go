package log

import (
	"github.com/sirupsen/logrus"
	"os"
)

type Fields = logrus.Fields

func Init(showDebug bool) {
	logrus.SetOutput(os.Stdout)

	logrus.SetFormatter(&logrus.TextFormatter{
		PadLevelText:    true,
		FullTimestamp:   true,
		TimestampFormat: "2006-01-02 15:04:05", // YYYY-MM-DD HH:mm:ss
	})

	logLevel := logrus.InfoLevel
	if showDebug {
		logLevel = logrus.DebugLevel
	}

	logrus.SetLevel(logLevel)
}

func Debug(msg string) {
	logrus.Debug(msg)
}

func DebugF(msg string, fields Fields) {
	logrus.WithFields(fields).Debug(msg)
}

func Info(msg string) {
	logrus.Info(msg)
}

func InfoF(msg string, fields Fields) {
	logrus.WithFields(fields).Info(msg)
}

func Warn(msg string) {
	logrus.Warn(msg)
}

func WarnF(msg string, fields Fields) {
	logrus.WithFields(fields).Warn(msg)
}

func Error(msg string) {
	logrus.Error(msg)
}

func ErrorF(msg string, fields Fields) {
	logrus.WithFields(fields).Error(msg)
}
