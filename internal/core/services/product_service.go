package services

import (
	"fmt"
	"log/slog"

	"github.com/Louis-Ai/products-service/internal/core/domain"
	"github.com/Louis-Ai/products-service/internal/core/ports"
)

type productService struct {
	repository ports.ProductRepository
	delivery   ports.DeliveryService
	logger     *slog.Logger
}

func NewProductService(repository ports.ProductRepository, delivery ports.DeliveryService, logger *slog.Logger) ports.ProductService {
	return &productService{
		repository: repository,
		delivery:   delivery,
		logger:     logger,
	}
}

func (ps *productService) GetPricedProducts() ([]domain.PricedProduct, error) {
	products, err := ps.repository.GetProductList()
	if err != nil {
		ps.logger.Error("Failed to retrieve products from repository", "error", err)
		return nil, err
	}

	var pricedProducts []domain.PricedProduct

	for _, p := range products {
		deliveryPrice, err := ps.delivery.Calculate(p)
		if err != nil {
			ps.logger.Error("Failed to calculate delivery price", "name", p.Name, "error", err)
			continue
		}

		totalPrice := p.Price + deliveryPrice

		pricedProduct := domain.PricedProduct{
			Name:          p.Name,
			DeliveryPrice: fmt.Sprintf("%.2f", deliveryPrice),
			ProductPrice:  fmt.Sprintf("%.2f", p.Price),
			TotalPrice:    fmt.Sprintf("%.2f", totalPrice),
		}

		pricedProducts = append(pricedProducts, pricedProduct)
	}

	ps.logger.Info("Calculated product prices", "provider", ps.delivery.ProviderName())
	return pricedProducts, nil
}
