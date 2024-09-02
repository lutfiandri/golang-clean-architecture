package organization_test

import (
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/lutfiandri/golang-clean-architecture/internal/bootstrap"
	"github.com/lutfiandri/golang-clean-architecture/internal/config"
	"github.com/lutfiandri/golang-clean-architecture/internal/entity"
	"github.com/lutfiandri/golang-clean-architecture/internal/infrastructure"
	"gorm.io/gorm"
)

var (
	app *fiber.App
	db  *gorm.DB
)

func bootstrapApp() {
	viperConfig := infrastructure.NewViper("../../../.env.testing")
	config.LoadEnv(viperConfig)

	log := infrastructure.NewLogger()

	validate := infrastructure.NewValidator()
	db = infrastructure.NewDatabase(log)
	app = infrastructure.NewFiber(&infrastructure.FiberConfig{
		HealthCheck: false,
		Logger:      false,
	})

	bootstrap.BootstrapApp(bootstrap.BootstrapAppConfig{
		DB:       db,
		App:      app,
		Log:      log,
		Validate: validate,
	})
}

func migrateUp() {
	db.Migrator().DropTable(&entity.Organization{})
	db.AutoMigrate(&entity.Organization{})
}

func migrateDown() {
	db.Migrator().DropTable(&entity.Organization{})
}

func seed() {
	tx := db.Begin()
	seedOrganization(tx)
	tx.Commit()
}

func TestMain(m *testing.M) {
	// before test
	bootstrapApp()
	migrateUp()
	seed()

	// test
	m.Run()

	// after test
	migrateDown()
}
