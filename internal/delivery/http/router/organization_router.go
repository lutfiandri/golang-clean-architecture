package router

import (
	"github.com/gofiber/fiber/v2"
	"github.com/lutfiandri/golang-clean-architecture/internal/delivery/http/controller.go"
)

func StartOrganizationRouter(app *fiber.App, controller controller.OrganizationController) {
	api := app.Group("/organizations")

	api.Post("/", controller.Create)
}
