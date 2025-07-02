package models

import "time"

type Vendor struct {
	ID           int       `json:"id"`
	UserID       int       `json:"user_id"`
	CompanyName  string    `json:"company_name"`
	Description  string    `json:"description,omitempty"`
	LogoURL      string    `json:"logo_url,omitempty"`
	Address      string    `json:"address"`
	ContactPhone string    `json:"contact_phone,omitempty"`
	ContactEmail string    `json:"contact_email,omitempty"`
	Status       string    `json:"status"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}
