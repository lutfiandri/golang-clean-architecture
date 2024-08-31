package infrastructure

import (
	"fmt"

	"github.com/spf13/viper"
)

func NewViper(configFile string) *viper.Viper {
	config := viper.New()

	config.SetConfigFile(configFile)
	config.SetConfigType("env")

	err := config.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("fatal error config file: %w", err))
	}

	return config
}
