package models

import "time"

type Order struct {
	ID              int       `json:"id"`
	BuyerUserID     int       `json:"buyer_user_id"`
	OrderDate       time.Time `json:"order_date"`
	TotalAmount     float64   `json:"total_amount"`
	OrderStatus     string    `json:"order_status"`
	DeliveryAddress string    `json:"delivery_address"`
	PaymentMethod   string    `json:"payment_method,omitempty"`
	PaymentStatus   string    `json:"payment_status"`
	Notes           string    `json:"notes,omitempty"`
	UpdatedAt       time.Time `json:"updated_at"`
}
