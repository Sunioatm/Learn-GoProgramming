package main

import (
	"goredis/handlers"
	"goredis/repositories"
	"goredis/services"

	"github.com/go-redis/redis/v8"
	"github.com/gofiber/fiber/v2"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	db := initDatabase()
	redisClient := initRedis()
	_ = redisClient
	// productRepo := repositories.NewProductRepositoryDB(db)

	// productRepo := repositories.NewProductRepositoryRedis(db, redisClient)
	productRepoDB := repositories.NewProductRepositoryDB(db)
	productService := services.NewProductService(productRepoDB)
	productHandler := handlers.NewProductHandlerRedis(productService, redisClient)

	app := fiber.New()

	app.Get("/products", productHandler.GetProducts)

	app.Listen("localhost:8000")
}

func initDatabase() *gorm.DB {
	dial := mysql.Open("root:123456@tcp(localhost:3307)/products")
	db, err := gorm.Open(dial, &gorm.Config{})

	if err != nil {
		panic(err)
	}

	return db
}

func initRedis() *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
	})
}
