package ports

import "github.com/Louis-Ai/products-service/internal/core/domain"

// Driven
type ProductRepository interface {
	GetProductList() ([]domain.Product, error)
}

type DeliveryService interface {
	Calculate(product domain.Product) (float64, error)
	ProviderName() string
}

// Driver
type ProductService interface {
	GetPricedProducts() ([]domain.PricedProduct, error)
}
