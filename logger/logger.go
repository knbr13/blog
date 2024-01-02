package logger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

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

func (l *internalLogger) Info(msg string, fields ...zapcore.Field) {
	l.logger.Info(msg, fields...)
}
