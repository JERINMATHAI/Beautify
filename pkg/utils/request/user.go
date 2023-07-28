package request

import "time"

type Address struct {
	ID           uint      `json:"address_id"`
	UserID       uint      `json:"-"`
	House        string    `json:"house"`
	AddressLine1 string    `json:"address_line_1"`
	AddressLine2 string    `json:"address_line_2"`
	City         string    `json:"city"`
	State        string    `json:"state"`
	PinCode      string    `json:"pin_code"`
	Country      string    `json:"country"`
	IsDefault    bool      `json:"is_default"`
	CreatedAt    time.Time `json:"-"`
	UpdatedAt    time.Time `json:"-"`
}

type AddressPatchReq struct {
	ID           uint      `json:"address_id"`
	UserID       uint      `json:"-"`
	House        string    `json:"house"`
	AddressLine1 string    `json:"address_line_1"`
	AddressLine2 string    `json:"address_line_2"`
	City         string    `json:"city"`
	State        string    `json:"state"`
	PinCode      string    `json:"pin_code"`
	Country      string    `json:"country"`
	IsDefault    bool      `json:"is_default"`
	UpdatedAt    time.Time `json:"-"`
}

type AddToCartReq struct {
	UserID         uint    `json:"user_id"`
	ProductID      uint    `json:"product_id" binding:"required"`
	Quantity       uint    `json:"quantity" binding:"required"`
	Price          float64 `json:"-"`
	Discount_price uint    `json:"-"`
}

type UpdateCartReq struct {
	UserID    uint `json:"-"`
	ProductID uint `json:"product_id" binding:"required"`
	Quantity  uint `json:"quantity" binding:"required"`
}

type DeleteCartItemReq struct {
	UserID    uint `json:"-"`
	ProductID uint `json:"product_id" binding:"required"`
}

type AddToWishReq struct {
	UserID         uint    `json:"user_id"`
	ProductID      uint    `json:"product_id" binding:"required"`
	Price          float64 `json:"-"`
	Discount_price uint    `json:"-"`
}

type UpdateWishReq struct {
	UserID    uint `json:"-"`
	ProductID uint `json:"product_id" binding:"required"`
}

type DeleteWishItemReq struct {
	UserID    uint `json:"-"`
	ProductID uint `json:"product_id" binding:"required"`
}
