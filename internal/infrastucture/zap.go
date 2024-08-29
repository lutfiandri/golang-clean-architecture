package infrastructure

import (
	"go.uber.org/zap"
)

func NewLogger() *zap.Logger {
	logger := zap.NewExample()
	return logger
}
