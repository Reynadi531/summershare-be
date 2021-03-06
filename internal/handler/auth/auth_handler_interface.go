package auth

import (
	"github.com/gofiber/fiber/v2"
)

type AuthHandler interface {
	Login(c *fiber.Ctx) error
	Register(c *fiber.Ctx) error
	RefreshToken(c *fiber.Ctx) error
}
