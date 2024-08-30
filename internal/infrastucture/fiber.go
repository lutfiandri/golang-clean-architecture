package infrastructure

import (
	"errors"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/healthcheck"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/lutfiandri/golang-clean-architecture/internal/config"
	"github.com/lutfiandri/golang-clean-architecture/internal/helper"
	"github.com/lutfiandri/golang-clean-architecture/internal/model"
	"gorm.io/gorm"
)

func NewFiber() *fiber.App {
	app := fiber.New(fiber.Config{
		AppName:      config.APP_NAME,
		Prefork:      config.APP_PREFORK,
		ErrorHandler: NewErrorHandler(),
	})

	app.Use(healthcheck.New())
	app.Use(logger.New())
	app.Use(recover.New(recover.Config{
		EnableStackTrace: true,
	}))
	app.Use(cors.New(cors.Config{
		AllowOrigins: config.CORS_ALLOW_ORIGIN,
		AllowMethods: "GET, POST, PUT, PATCH, DELETE",
	}))

	return app
}

func NewErrorHandler() fiber.ErrorHandler {
	return func(c *fiber.Ctx, err error) error {
		code := fiber.StatusInternalServerError

		var e *fiber.Error
		if errors.As(err, &e) {
			code = e.Code
		}

		c.Set(fiber.HeaderContentType, fiber.MIMEApplicationJSONCharsetUTF8)

		validationErrorsResponse := helper.GetValidationErrorsResponse(c)

		switch err {
		case gorm.ErrRecordNotFound:
			code = fiber.StatusNotFound
		}

		response := model.NewErrorResponse(err, validationErrorsResponse)
		return c.Status(code).JSON(response)
	}
}
