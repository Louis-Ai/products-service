package repository

import (
	"encoding/json"
	"os"

	"github.com/Louis-Ai/products-service/internal/core/domain"
	"github.com/Louis-Ai/products-service/internal/core/ports"
)

type jsonRepository struct {
	filepath string
}

func NewJSONRepository(filepath string) ports.ProductRepository {
	return &jsonRepository{
		filepath: filepath,
	}
}

func (r *jsonRepository) GetProductList() ([]domain.Product, error) {
	file, err := os.Open(r.filepath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var products []domain.Product

	decoder := json.NewDecoder(file)
	err = decoder.Decode(&products)
	if err != nil {
		return nil, err
	}

	return products, nil
}
