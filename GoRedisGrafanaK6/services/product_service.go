package services

import "goredis/repositories"

type productService struct {
	productRepo repositories.ProductRepository
}

func NewProductService(productRepo repositories.ProductRepository) ProductService {
	return productService{productRepo: productRepo}
}

func (s productService) GetProducts() (products []Product, err error) {
	productsDB, err := s.productRepo.GetProducts()
	if err != nil {
		return nil, err
	}
	for _, productDB := range productsDB {
		product := Product{
			ID:       productDB.ID,
			Name:     productDB.Name,
			Quantity: productDB.Quantity,
		}
		products = append(products, product)
	}
	return products, nil
}
