package handlers

import "github.com/gofiber/fiber/v2"

type ProductHandler interface {
	GetProducts(c *fiber.Ctx) error
}
