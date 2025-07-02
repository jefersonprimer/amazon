package models

import (
	"time"
)

type ProductImage struct {
	ID               uint      `gorm:"primaryKey;autoIncrement" json:"id"`
	ProductServiceID uint      `gorm:"not null;index" json:"product_service_id"`
	ImageURL         string    `gorm:"type:varchar(255);not null" json:"image_url"`
	IsMain           bool      `gorm:"default:false" json:"is_main"`
	DisplayOrder     int       `gorm:"default:0" json:"display_order"`
	CreatedAt        time.Time `gorm:"autoCreateTime" json:"created_at"`

	ProductService ProductService `gorm:"foreignKey:ProductServiceID;constraint:OnDelete:CASCADE;" json:"-"`
}
