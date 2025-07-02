package models

import (
	"time"
)

type Category struct {
	ID               uint      `gorm:"primaryKey;autoIncrement" json:"id"`
	Name             string    `gorm:"type:varchar(100);unique;not null" json:"name"`
	Description      string    `gorm:"type:text" json:"description"`
	ParentCategoryID *uint     `gorm:"column:parent_category_id" json:"parent_category_id"`
	ParentCategory   *Category `gorm:"foreignKey:ParentCategoryID" json:"parent_category,omitempty"`
	CreatedAt        time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt        time.Time `gorm:"autoUpdateTime" json:"updated_at"`
}
