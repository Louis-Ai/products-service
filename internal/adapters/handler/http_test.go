package handler

import (
	"io"
	"log/slog"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/Louis-Ai/products-service/internal/core/domain"
)

type mockProductService struct {
	products []domain.PricedProduct
	err      error
}

func (m *mockProductService) GetPricedProducts() ([]domain.PricedProduct, error) {
	return m.products, m.err
}

func TestHTTP_GetProducts(t *testing.T) {
	testLogger := slog.New(slog.NewTextHandler(io.Discard, nil))

	t.Run("200 - Successfully returned product list", func(t *testing.T) {
		mockService := &mockProductService{
			products: []domain.PricedProduct{
				{
					Name:          "Test 1",
					DeliveryPrice: "10.00",
					ProductPrice:  "200.00",
					TotalPrice:    "210.00",
				},
			},
		}

		handler := NewHTTPHandler(mockService, testLogger)

		request := httptest.NewRequest(http.MethodGet, "/products", nil)
		response := httptest.NewRecorder()

		handler.GetProducts(response, request)

		if response.Code != http.StatusOK {
			t.Fatalf("Expected status code %v, got %v", http.StatusOK, response.Code)
		}

		expectedBody := `[{"name":"Test 1","delivery_price":"10.00","product_price":"200.00","total_price":"210.00"}]`

		if expectedBody != strings.TrimSpace(response.Body.String()) {
			t.Errorf("Expected: %v, got: %v", expectedBody, strings.TrimSpace(response.Body.String()))
		}

	})
}
