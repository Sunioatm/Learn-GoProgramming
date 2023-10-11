package handlers

import (
	"context"
	"encoding/json"
	"fmt"
	"goredis/services"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/gofiber/fiber/v2"
)

type productHandlerRedis struct {
	productService services.ProductService
	redisClient    *redis.Client
}

func NewProductHandlerRedis(productService services.ProductService, redisClient *redis.Client) ProductHandler {
	return productHandlerRedis{
		productService: productService,
		redisClient:    redisClient,
	}
}

func (h productHandlerRedis) GetProducts(c *fiber.Ctx) error {

	key := "handler::GetProducts"

	// Redis Get
	if resJson, err := h.redisClient.Get(context.Background(), key).Result(); err == nil {
		fmt.Println("From Redis")
		c.Set("Content-Type", "application/json")
		return c.SendString(resJson)
	}

	// Service
	products, err := h.productService.GetProducts()
	if err != nil {
		return err
	}

	res := fiber.Map{
		"status":   "OK",
		"products": products,
	}

	// Redis Set
	if data, err := json.Marshal(res); err == nil {
		h.redisClient.Set(context.Background(), key, string(data), time.Second*10)
	}

	fmt.Println("From Database")
	return c.JSON(res)

}
