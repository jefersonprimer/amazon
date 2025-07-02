package models

import "time"

type OrderItem struct {
	ID               int       `json:"id"`
	OrderID          int       `json:"order_id"`
	ProductServiceID int       `json:"product_service_id"`
	VendorID         int       `json:"vendor_id"`
	Quantity         int       `json:"quantity"`
	UnitPrice        float64   `json:"unit_price"`
	Subtotal         float64   `json:"subtotal"`
	CreatedAt        time.Time `json:"created_at"`
}
