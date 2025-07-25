package delivery

import "github.com/Louis-Ai/products-service/internal/core/domain"

type UPSService struct {
	pricePerGram float64
}

func NewUPSService(pricePergram float64) *UPSService {
	return &UPSService{
		pricePerGram: pricePergram,
	}
}

func (s *UPSService) Calculate(product domain.Product) (float64, error) {
	deliveryPrice := s.pricePerGram * product.Weight
	return deliveryPrice, nil
}

func (s *UPSService) ProviderName() string {
	return "UPS"
}
