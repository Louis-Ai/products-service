package delivery

import "testing"

func TestNewDeliverySelector(t *testing.T) {
	t.Run("Should return UPSService for UPS", func(t *testing.T) {
		provider := "ups"
		price := 0.01

		service, err := NewDeliverySelector(provider, price)

		if err != nil {
			t.Errorf("Unexpected error creating delivery service, got:%v", err)
		}

		_, ok := service.(*UPSService)
		if !ok {
			t.Errorf("Expected UPSService to be true, got: %v", ok)
		}
	})
}
