package delivery

import (
	"fmt"
	"strings"

	"github.com/Louis-Ai/products-service/internal/core/ports"
)

func NewDeliverySelector(provider string, pricePerGram float64) (ports.DeliveryService, error) {
	switch strings.ToUpper(provider) {
	case "UPS":
		return NewUPSService(pricePerGram), nil
	case "ROYAL-MAIL":
		return NewRoyalMailService(pricePerGram), nil
	default:
		return nil, fmt.Errorf("delivery provider does not match any available: %v", provider)
	}
}
