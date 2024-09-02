package exception

import "github.com/gofiber/fiber/v2"

var (
	ErrOrganizationAlreadyExists = fiber.NewError(fiber.StatusConflict, "organization already exists")
	ErrOrganizationNotFound      = fiber.NewError(fiber.StatusNotFound, "organization not found")
)
