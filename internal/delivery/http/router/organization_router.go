package router

import (
	"github.com/gofiber/fiber/v2"
	"github.com/lutfiandri/golang-clean-architecture/internal/delivery/http/controller.go"
)

func SetupOrganizationRouter(app *fiber.App, controller controller.OrganizationController) {
	api := app.Group("/organizations")

	api.Post("/", controller.Create)
	api.Get("/", controller.GetMany)
	api.Get("/:id", controller.Get)
	api.Put("/:id", controller.Update)
	api.Delete("/:id", controller.Delete)
}
