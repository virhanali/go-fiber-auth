package models

import "gorm.io/gorm"

// Product struct
type Product struct {
	gorm.Model  `json:"-"`
	Title       string `gorm:"not null" json:"title"`
	Description string `gorm:"not null" json:"description"`
	Amount      int    `gorm:"not null" json:"amount"`
}

type ProductRequest struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Amount      int    `json:"amount"`
}
