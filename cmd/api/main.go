package main

import (
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"strconv"

	"github.com/Louis-Ai/products-service/internal/adapters/delivery"
	"github.com/Louis-Ai/products-service/internal/adapters/handler"
	"github.com/Louis-Ai/products-service/internal/adapters/repository"
	"github.com/Louis-Ai/products-service/internal/core/services"
	"github.com/joho/godotenv"
)

func main() {
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))

	err := godotenv.Load()
	if err != nil {
		logger.Warn("No .env file found. Using environment variables")
	}

	port := os.Getenv("PORT")
	if port == "" {
		logger.Info("No port provided, running on default 8080")
		port = "8080"
	}

	priceString := os.Getenv("DELIVERY_PRICE")
	if priceString == "" {
		logger.Error("DELIVERY_PRICE environment variable unset")
		os.Exit(1)
	}

	pricePerGram, err := strconv.ParseFloat(priceString, 64)
	if err != nil {
		logger.Error("Invalid DELIVERY_PRICE: %v", err)
		os.Exit(1)
	}

	deliveryProvider := os.Getenv("DELIVERY_PROVIDER")
	if deliveryProvider == "" {
		logger.Error("DELIVERY_PROVIDER environment variable unset")
	}

	deliveryService, err := delivery.NewDeliverySelector(deliveryProvider, pricePerGram)
	if err != nil {
		logger.Error("Error setting up delivery service: %v", err)
		os.Exit(1)
	}

	productRepository := repository.NewJSONRepository("products.json")

	productService := services.NewProductService(productRepository, deliveryService, logger)

	httpHandler := handler.NewHTTPHandler(productService, logger)

	http.HandleFunc("/products", httpHandler.GetProducts)

	addr := fmt.Sprintf(":%s", port)

	http.ListenAndServe(addr, nil)
}
