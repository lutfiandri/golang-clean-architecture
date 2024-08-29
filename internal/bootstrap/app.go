package bootstrap

import (
	"strconv"

	"github.com/lutfiandri/golang-clean-architecture/internal/config"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type BootstrapAppConfig struct {
	App      *fiber.App
	DB       *gorm.DB
	Log      *zap.Logger
	Validate *validator.Validate
}

func BootstrapApp(bootstrapAppConfig BootstrapAppConfig) {
	bootstrapAppConfig.App.Listen(":" + strconv.Itoa(config.APP_PORT))
}
