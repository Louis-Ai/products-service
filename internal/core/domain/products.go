package domain

type Product struct {
	Name   string  `json:"name"`
	Weight float64 `json:"weight"`
	Price  float64 `json:"price"`
}

type PricedProduct struct {
	Name          string `json:"name"`
	DeliveryPrice string `json:"delivery_price"`
	ProductPrice  string `json:"product_price"`
	TotalPrice    string `json:"total_price"`
}
