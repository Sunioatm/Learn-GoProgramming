package handlers

import (
	"goredis/services"

	"github.com/gofiber/fiber/v2"
)

type productHandler struct {
	productService services.ProductService
}

func NewProductHandler(productService services.ProductService) ProductHandler {
	return productHandler{productService: productService}
}

func (h productHandler) GetProducts(c *fiber.Ctx) error {

	products, err := h.productService.GetProducts()
	if err != nil {
		return err
	}

	response := fiber.Map{
		"status":  "OK",
		"product": products,
	}

	return c.JSON(response)
}
