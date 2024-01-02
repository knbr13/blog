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

func (l *internalLogger) Error(msg string, fields ...zapcore.Field) {
	l.logger.Error(msg, fields...)
}

func (l *internalLogger) Debug(msg string, fields ...zapcore.Field) {
	l.logger.Debug(msg, fields...)
}

func (l *internalLogger) Warn(msg string, fields ...zapcore.Field) {
	l.logger.Warn(msg, fields...)
}

func (l *internalLogger) Fatal(msg string, fields ...zapcore.Field) {
	l.logger.Fatal(msg, fields...)
}

func (l *internalLogger) Panic(msg string, fields ...zapcore.Field) {
	l.logger.Panic(msg, fields...)
}
