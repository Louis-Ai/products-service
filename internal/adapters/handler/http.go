package handler

import (
	"encoding/json"
	"log/slog"
	"net/http"

	"github.com/Louis-Ai/products-service/internal/core/ports"
)

type HTTPHandler struct {
	productService ports.ProductService
	logger         *slog.Logger
}

func NewHTTPHandler(productService ports.ProductService, logger *slog.Logger) *HTTPHandler {
	return &HTTPHandler{
		productService: productService,
		logger:         logger,
	}
}

func (h *HTTPHandler) GetProducts(w http.ResponseWriter, r *http.Request) {
	pricedProducts, err := h.productService.GetPricedProducts()
	if err != nil {
		h.logger.Error("Service failed retrieving products pricelist", "error", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	err = json.NewEncoder(w).Encode(pricedProducts)
	if err != nil {
		h.logger.Error("Failed to encode json response", "error", err)
	}
}
