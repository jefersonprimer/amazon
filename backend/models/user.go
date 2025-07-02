package models

import "time"

type User struct {
	ID           int       `json:"id"`
	FullName     string    `json:"full_name"`
	Email        string    `json:"email"`
	PasswordHash string    `json:"-"` // Oculta o hash da senha em respostas JSON
	UserType     string    `json:"user_type"`
	PhoneNumber  string    `json:"phone_number,omitempty"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}
