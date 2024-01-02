package logger

import "go.uber.org/zap"

type internalLogger struct {
	logger *zap.Logger
}
