package domain

import "gorm.io/gorm"

type Product struct {
	gorm.Model
	Name        string  `json:"name", gorm:"index;"`
	Description string  `json:"description"`
	CategoryID  uint    `json:"category_id"`
	ImageURL    string  `json:"image_url"`
	Price       float64 `json:"price"`
	Stock       int     `json:"stock"`
}
