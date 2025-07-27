package delivery

import "github.com/Louis-Ai/products-service/internal/core/domain"

type RoyalMailService struct {
	pricePerGram float64
}

func NewRoyalMailService(pricePerGram float64) *RoyalMailService {
	return &RoyalMailService{
		pricePerGram: pricePerGram,
	}
}

func (s *RoyalMailService) Calculate(product domain.Product) (float64, error) {
	if product.Weight <= 2000 {
		return 4.50, nil
	}
	return 15.00, nil
}

func (s *RoyalMailService) ProviderName() string {
	return "Royal Mail"
}
