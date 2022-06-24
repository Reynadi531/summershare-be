package post

import "github.com/gofiber/fiber/v2"

type PostHandler interface {
	Create(c *fiber.Ctx) error
	Update(c *fiber.Ctx) error
	ViewByID(c *fiber.Ctx) error
	ViewAll(c *fiber.Ctx) error
	ViewSelf(c *fiber.Ctx) error
}
