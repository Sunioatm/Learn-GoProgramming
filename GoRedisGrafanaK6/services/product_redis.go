package services

import (
	"context"
	"encoding/json"
	"fmt"
	"goredis/repositories"
	"time"

	"github.com/go-redis/redis/v8"
)

type productServiceRedis struct {
	productRepo repositories.ProductRepository
	redisClient *redis.Client
}

func NewProductServiceRedis(productRepo repositories.ProductRepository, redisClient *redis.Client) ProductService {
	return productServiceRedis{
		productRepo: productRepo,
		redisClient: redisClient,
	}
}

func (s productServiceRedis) GetProducts() (products []Product, err error) {

	key := "service::GetProducts"

	// Redis GET
	if productJson, err := s.redisClient.Get(context.Background(), key).Result(); err == nil {
		if json.Unmarshal([]byte(productJson), &products) == nil {
			fmt.Println("From Redis")
			return products, nil
		}
	}

	// Repository
	productsDB, err := s.productRepo.GetProducts()
	if err != nil {
		return nil, err
	}
	for _, productDB := range productsDB {
		products = append(products, Product{
			ID:       productDB.ID,
			Name:     productDB.Name,
			Quantity: productDB.Quantity,
		})
	}

	// Redis SET
	if data, err := json.Marshal(products); err == nil {
		s.redisClient.Set(context.Background(), key, string(data), time.Second*10)
	}

	fmt.Println("From Database")
	return products, nil
}
