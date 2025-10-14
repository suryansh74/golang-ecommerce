package domain

import (
	"gorm.io/gorm"
	"time"
)

type User struct {
	gorm.Model
	FirstName string    `json:"first_name"`
	LastName  string    `json:"last_name"`
	Phone     string    `json:"phone"`
	Email     string    `json:"email" gorm:"index;unique;not null"`
	Password  string    `json:"password"`
	Code      int       `json:"code"`
	Expiry    time.Time `json:"expiry"`
	Verified  bool      `json:"verified" gorm:"default:false"`
	UserType  string    `json:"user_type" gorm:"default:buyer"`
}
