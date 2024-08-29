package main

import (
	"github.com/lutfiandri/golang-clean-architecture/internal/bootstrap"
	"github.com/lutfiandri/golang-clean-architecture/internal/config"
	infrastructure "github.com/lutfiandri/golang-clean-architecture/internal/infrastucture"
)

func main() {
	viperConfig := infrastructure.NewViper()
	config.LoadEnv(viperConfig)

	log := infrastructure.NewLogger()
	db := infrastructure.NewDatabase(log)
	validate := infrastructure.NewValidator()
	app := infrastructure.NewFiber()

	bootstrap.BootstrapApp(bootstrap.BootstrapAppConfig{
		DB:       db,
		App:      app,
		Log:      log,
		Validate: validate,
	})
}
