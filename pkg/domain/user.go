package domain

import (
	"time"

	"gorm.io/gorm"
)

type Users struct {
	gorm.Model
	ID          uint   `json:"id" gorm:"primaryKey;unique;autoIncrement"`
	UserName    string `json:"user_name" gorm:"not null" binding:"required,min=3,max=15"`
	FirstName   string `json:"first_name" gorm:"not null" binding:"required,min=2,max=20"`
	LastName    string `json:"last_name" gorm:"not null" binding:"required,min=2,max=20"`
	Age         uint   `json:"age" gorm:"not null" binding:"required,numeric"`
	Email       string `json:"email" gorm:"unique;not null" binding:"required,email"`
	Phone       string `json:"phone" gorm:"unique;not null" binding:"required,min=10,max=10"`
	Password    string `json:"password" gorm:"not null" binding:"required,eqfield=ConfirmPassword"`
	BlockStatus bool   `json:"block_status" gorm:"not null;default:false"`
}

type Address struct {
	gorm.Model
	ID           uint   `json:"id" gorm:"primaryKey;unique;autoIncrement"`
	UserID       uint   `json:"-"`
	House        string `json:"house" gorm:"not null"`
	AddressLine1 string `json:"address_line_1" gorm:"not null" binding:"required,min=2,max=40"`
	AddressLine2 string `json:"address_line_2" gorm:"not null" binding:"required,min=2,max=40"`
	City         string `json:"city" gorm:"not null" binding:"required,min=2,max=20"`
	State        string `json:"state" gorm:"not null" binding:"required,min=2,max=20"`
	PinCode      string `json:"pin_code" gorm:"not null" binding:"required,min=2,max=10"`
	Country      string `json:"country" gorm:"not null" binding:"required,min=2,max=20"`
	IsDefault    bool   `gorm:"not null"`
}

type CartItems struct {
	gorm.Model
	ID          uint           `gorm:"primaryKey"`
	CartID      uint           `gorm:"not null"`
	ProductID   uint           `gorm:"not null"`
	Quantity    uint           `gorm:"not null"`
	StockStatus bool           `gorm:"not null;default:true"`
	Price       float64        `gorm:"not null"`
	letedAt     gorm.DeletedAt `gorm:"index"`
}

type Cart struct {
	ID     uint    `gorm:"primaryKey"`
	UserID uint    `gorm:"not null"`
	Total  float64 `gorm:"default:0"`
}

type WishListItems struct {
	gorm.Model
	ID          uint    `gorm:"primaryKey"`
	WishID      uint    `gorm:"not null"`
	ProductID   uint    `gorm:"not null"`
	Quantity    uint    `gorm:"not null"`
	StockStatus bool    `gorm:"not null;default:true"`
	Price       float64 `gorm:"not null"`
}

type WishList struct {
	ID     uint    `gorm:"primaryKey"`
	UserID uint    `gorm:"not null"`
	Total  float64 `gorm:"default:0"`
}

type Wallet struct {
	ID        uint    `gorm:"primaryKey"`
	UserID    uint    `gorm:"not null"`
	Balance   float64 `gorm:"not null"`
	Remark    string
	UpdatedAt time.Time
	CreatedAt time.Time
}
