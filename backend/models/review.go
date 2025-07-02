package models

import "time"

type Review struct {
	ID               int       `json:"id" gorm:"primaryKey;autoIncrement"`
	UserID           int       `json:"user_id" gorm:"not null"`
	ProductServiceID *int      `json:"product_service_id" gorm:"default:null"`
	VendorID         *int      `json:"vendor_id" gorm:"default:null"`
	Rating           int       `json:"rating" gorm:"not null;check:rating >= 1 AND rating <= 5"`
	Comment          *string   `json:"comment"`
	ReviewDate       time.Time `json:"review_date" gorm:"autoCreateTime"`
}
