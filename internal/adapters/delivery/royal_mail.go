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
	deliveryPrice := s.pricePerGram * product.Weight

	return deliveryPrice, nil
}

func (s *RoyalMailService) ProviderName() string {
	return "Royal Mail"
}
