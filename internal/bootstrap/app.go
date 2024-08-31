package bootstrap

import (
	"github.com/lutfiandri/golang-clean-architecture/internal/delivery/http/controller.go"
	"github.com/lutfiandri/golang-clean-architecture/internal/delivery/http/router"
	"github.com/lutfiandri/golang-clean-architecture/internal/entity"
	"github.com/lutfiandri/golang-clean-architecture/internal/repository"
	"github.com/lutfiandri/golang-clean-architecture/internal/usecase"

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

func BootstrapApp(cfg BootstrapAppConfig) {
	// setup automigrate
	cfg.DB.AutoMigrate(&entity.Organization{}, &entity.Role{}, &entity.User{})

	// repository
	organizationRepository := repository.NewOrganizationRepository(cfg.Log)

	// usecase
	organizationUseCase := usecase.NewOrganizationUseCase(cfg.DB, organizationRepository)

	// controller
	organizationController := controller.NewOrganizationController(cfg.App, cfg.Validate, organizationUseCase)

	// router
	router.SetupOrganizationRouter(cfg.App, organizationController)
}
