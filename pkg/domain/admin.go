package domain

import (
	"gorm.io/gorm"
)

type Admin struct {
	gorm.Model
	ID       uint   `json:"-" gorm:"primaryKey;not null"`
	UserName string `json:"user_name" gorm:"not null" binding:"omitempty,min=4,max=15"`
	Email    string `json:"email" gorm:"not null" binding:"omitempty,email"`
	Password string `json:"Password" gorm:"not null" binding:"required,min=3,max=30"`
	// CreatedAt time.Time `json:"-" gorm:"not null"`
	// UpdatedAt time.Time `json:"-"`
}
