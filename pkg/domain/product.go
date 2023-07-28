package domain

import (
	"gorm.io/gorm"
)

// Category struct
type Category struct {
	ID           uint       `json:"id" gorm:"primaryKey;not null;autoIncrement"`
	CategoryName string     `json:"category_name" gorm:"unique;not null"`
	Products     []*Product `json:"-" gorm:"foreignKey:CategoryID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	//Parent       *Category  `json:"-" gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	//Children     []*Category `json:"-" gorm:"foreignKey:ParentID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}

// Product struct
type Product struct {
	gorm.Model
	ID            uint      `json:"id" gorm:"primaryKey;not null;autoIncrement"`
	Name          string    `json:"product_name" gorm:"not null;size:50"`
	Description   string    `json:"description" gorm:"not null;size:500"`
	CategoryID    uint      `json:"brand_id" gorm:"index;not null"`
	Category      *Category `json:"-" gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Price         uint      `json:"price" gorm:"not null"`
	DiscountPrice uint      `json:"discount_price" gorm:"default:null"`
	Image         string    `json:"image" gorm:"not null"`
	// CreatedAt     time.Time      `json:"created_at" gorm:"not null"`
	// UpdatedAt     time.Time      `json:"updated_at" gorm:"default:null"`
	// Items         []*ProductItem `json:"-" gorm:"foreignKey:ProductID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}

// ProductItem struct
type ProductItem struct {
	gorm.Model
	ID            uint   `json:"id" gorm:"primaryKey;not null;autoIncrement"`
	ProductID     uint   `json:"product_id" gorm:"index;"`
	QtyInStock    uint   `json:"qty_in_stock" gorm:"not null"`
	StockStatus   bool   `json:"stock_status" gorm:"not null;default:true;type:boolean;"`
	Price         uint   `json:"price" gorm:"not null"`
	SKU           string `json:"sku" gorm:"unique;not null"`
	DiscountPrice uint   `json:"discount_price" gorm:"default:null"`
	// CreatedAt     time.Time `json:"created_at" gorm:"not null"`
	// UpdatedAt     time.Time `json:"updated_at" gorm:"default:null"`
	// //Configurations []*ProductConfig `json:"product_config" gorm:"many2many:product_configurations;"`
	//Images []*ProductImage `json:"-" gorm:"foreignKey:ProductItemID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}

type ProductImage struct {
	ID            uint        `json:"id" gorm:"primaryKey;not null;autoIncrement"`
	ProductItemID uint        `json:"product_item_id" gorm:"not null"`
	ProductItem   ProductItem `json:"-" gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL"`
	Image         string      `json:"image" gorm:"not null"`
}
