package bootstrap

import (
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v3"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type BootstrapAppConfig struct {
	App      *fiber.App
	DB       *gorm.DB
	Log      *zap.Logger
	Validate *validator.Validate
	Config   *viper.Viper
}

func BootstrapApp(config *BootstrapAppConfig) {
	config.App.Listen(":"+config.Config.GetString("APP_PORT"), fiber.ListenConfig{
		EnablePrefork: config.Config.GetBool("APP_PREFORK"),
	})
}
