package middleware

import (
	"slices"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/lutfiandri/golang-clean-architecture/internal/config"
	"github.com/lutfiandri/golang-clean-architecture/internal/helper"
	"github.com/lutfiandri/golang-clean-architecture/internal/model"
)

const (
	CtxClaims = "claims"
)

func NewAuthenticator() fiber.Handler {
	return func(c *fiber.Ctx) error {
		authHeader := c.Get("Authorization")
		authHeaderList := strings.Split(authHeader, " ")

		if len(authHeaderList) < 2 || authHeaderList[0] != "Bearer" || authHeaderList[1] == "" {
			return fiber.NewError(fiber.StatusUnauthorized, "Invalid or missing Bearer token")
		}

		tokenString := authHeaderList[1]

		claims, err := helper.ParseJwt(tokenString, config.JWT_SECRET_KEY)
		if err != nil {
			return fiber.NewError(fiber.StatusUnauthorized, err.Error())
		}

		c.Locals("user", claims.User)
		return c.Next()
	}
}

func GetUser(c *fiber.Ctx) model.JwtUser {
	return c.Locals("user").(model.JwtUser)
}

func NewRoleAuthorizer(roles ...uint) fiber.Handler {
	return func(c *fiber.Ctx) error {
		user := GetUser(c)

		if !slices.Contains(roles, user.Role) {
			return fiber.NewError(fiber.StatusForbidden, "Doesn't have permission to access this resource")
		}

		return c.Next()
	}
}
