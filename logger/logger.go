package logger

import "go.uber.org/zap"

type internalLogger struct {
	logger *zap.Logger
}

var logger internalLogger

func init() {
	zapLogger, err := zap.NewDevelopment()
	if err != nil {
		panic(err)
	}
	logger.logger = zapLogger
}

func GetLogger() *internalLogger {
	return &logger
}
