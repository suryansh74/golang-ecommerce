package domain

import "gorm.io/gorm"

type Category struct {
	gorm.Model
	Name         string    `json:"name", gorm:"index;"`
	ParentID     uint      `json:"parent_id"`
	ImageURL     string    `json:"image_url"`
	Products     []Product `json:"products"`
	DisplayOrder int       `json:"display_order"`
}
