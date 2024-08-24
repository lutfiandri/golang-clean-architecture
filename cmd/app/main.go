package main

import (
	"fmt"

	"github.com/lutfiandri/golang-clean-architecture/internal/bootstrap"
	"github.com/lutfiandri/golang-clean-architecture/internal/config"
	"go.uber.org/zap"
)

func main() {
	viperConfig := config.NewViper()
	log := config.NewLogger(viperConfig)
	db := config.NewDatabase(viperConfig, log)
	validate := config.NewValidator(viperConfig)
	app := config.NewFiber(viperConfig)

	bootstrap.BootstrapApp(&bootstrap.BootstrapAppConfig{
		DB:       db,
		App:      app,
		Log:      log,
		Validate: validate,
		Config:   viperConfig,
	})

	webPort := viperConfig.GetInt("web.port")
	err := app.Listen(fmt.Sprintf(":%d", webPort))
	if err != nil {
		log.Fatal("Failed to start server: %v", zap.Error(err))
	}
}
