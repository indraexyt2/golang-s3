package helpers

import "github.com/sirupsen/logrus"

type ILogger interface {
	Info(msg string, args ...interface{})
	Error(msg string, args ...interface{})
}

type LogrusWrapper struct {
	logger *logrus.Logger
}

func (l *LogrusWrapper) Info(msg string, args ...interface{}) {
	l.logger.Infof(msg, args...)
}

func (l *LogrusWrapper) Error(msg string, args ...interface{}) {
	l.logger.Errorf(msg, args...)
}

var Logger ILogger

func SetupLogger() {
	logger := logrus.New()
	logger.SetFormatter(&logrus.JSONFormatter{
		PrettyPrint: true,
	})

	Logger = &LogrusWrapper{logger: logger}
	logger.Info("Setup logger")
}
