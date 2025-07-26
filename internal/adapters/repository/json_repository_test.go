package repository

import (
	"os"
	"path/filepath"
	"reflect"
	"testing"

	"github.com/Louis-Ai/products-service/internal/core/domain"
)

func TestJSONRepo_GetProductList(t *testing.T) {
	t.Run("Successfully retrieves products from a json file", func(t *testing.T) {
		fileContent := `[{"name": "Test Product 1", "price": 99.99, "weight": 1000}]`

		tempDirectory := t.TempDir()
		tempFile := filepath.Join(tempDirectory, "products.json")

		os.WriteFile(tempFile, []byte(fileContent), 0644)

		repository := NewJSONRepository(tempFile)

		products, err := repository.GetProductList()

		if err != nil {
			t.Errorf("Unexpected error, got:%v", err)
		}

		expectedData := domain.Product{
			Name:   "Test Product 1",
			Price:  99.99,
			Weight: 1000,
		}

		if !reflect.DeepEqual(expectedData, products[0]) {
			t.Errorf("Expected: %+v, got: %+v", expectedData, products[0])
		}

	})
}
