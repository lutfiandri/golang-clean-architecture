package config

import (
	"github.com/spf13/viper"
	"go.uber.org/zap"
)

func NewLogger(viper *viper.Viper) *zap.Logger {
	logger := zap.NewExample()
	return logger
}
