package models

import "time"

type ProductService struct { // Nome da struct pode ser ProductService para englobar ambos
	ID            int       `json:"id"`
	VendorID      int       `json:"vendor_id"`
	CategoryID    int       `json:"category_id"`
	Name          string    `json:"name"`
	Description   string    `json:"description,omitempty"`    // omitempty para campos opcionais
	Price         float64   `json:"price"`                    // Use float64 ou decimal.Decimal se precisar de maior precisão
	ItemType      string    `json:"item_type"`                // 'product' or 'service'
	StockQuantity *int      `json:"stock_quantity,omitempty"` // *int para aceitar NULL
	UnitOfMeasure string    `json:"unit_of_measure,omitempty"`
	Status        string    `json:"status"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
}
