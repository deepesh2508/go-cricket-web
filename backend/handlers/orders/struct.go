package orders

type CreateOrderReq struct {
	UserID          int    `json:"user_id"`
	PaymentMethodID int    `json:"payment_method_id"`
	ShippingAddress string `json:"shipping_address"`
}

type CreateOrderRes struct {
	Message string `json:"message"`
}

type OrderItem struct {
	ProductID  int     `json:"product_id"`
	Quantity   int     `json:"quantity"`
	Price      float64 `json:"price"`
	TotalPrice float64 `json:"total_price"`
}

type BuyNowReq struct {
	UserID          int    `json:"user_id"`
	ProductID       int    `json:"product_id"`
	Quantity        int    `json:"quantity"`
	PaymentMethodID int    `json:"payment_method_id"`
	ShippingAddress string `json:"shipping_address"`
}

type BuyNowRes struct {
	Message string `json:"message"`
}
