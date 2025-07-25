package services

import (
	"io"
	"log/slog"
	"reflect"
	"testing"

	"github.com/Louis-Ai/products-service/internal/core/domain"
)

type mockProductRepository struct {
	products []domain.Product
	err      error
}

func (m *mockProductRepository) GetProductList() ([]domain.Product, error) {
	return m.products, m.err
}

type mockDeliveryService struct {
	price        float64
	err          error
	providerName string
}

func (m *mockDeliveryService) Calculate(product domain.Product) (float64, error) {
	return m.price, m.err
}

func (m *mockDeliveryService) ProviderName() string {
	return m.providerName
}

func TestProductService_GetPricedProducts(t *testing.T) {
	testLogger := slog.New(slog.NewTextHandler(io.Discard, nil))

	testProduct := domain.Product{
		Name:   "Item 1",
		Weight: 90,
		Price:  110.00,
	}

	//successfully get priced product
	t.Run("Return priced product", func(t *testing.T) {
		mockRepository := &mockProductRepository{
			products: []domain.Product{testProduct},
		}

		mockDelivery := &mockDeliveryService{
			price:        0.90,
			providerName: "mock-deliveries",
		}

		service := NewProductService(mockRepository, mockDelivery, testLogger)

		pricedProducts, err := service.GetPricedProducts()

		if err != nil {
			t.Fatalf("Error not expected, got: %v", err)
		}

		if len(pricedProducts) != 1 {
			t.Fatalf("Expected 1 product with pricing, got: %v", len(pricedProducts))
		}

		expectedData := domain.PricedProduct{
			Name:          "Item 1",
			ProductPrice:  "110.00",
			DeliveryPrice: "0.90",
			TotalPrice:    "110.90",
		}

		if !reflect.DeepEqual(expectedData, pricedProducts[0]) {
			t.Errorf("Expected: %+v, got: %+v", expectedData, pricedProducts[0])
		}

	})

	//Repository failure
	t.Run("Return an error when repository fails", func(t *testing.T) {

	})

	//Delivery service failure
	t.Run("Return an error when delivery service fails", func(t *testing.T) {

	})
}
