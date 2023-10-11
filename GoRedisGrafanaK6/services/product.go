package services

type Product struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Quantity int    `json:"quantity"`
}

type ProductService interface {
	GetProducts() ([]Product, error)
}
