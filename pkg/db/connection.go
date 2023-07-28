package db

import (
	"beautify/pkg/config"
	"beautify/pkg/domain"
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectDatabase(cfg config.Config) (*gorm.DB, error) {

	psqlInfo := fmt.Sprintf("host= %s user= %s password=%s dbname=%s port=%s sslmode=disable", cfg.DBHost, cfg.DBUser, cfg.DBPassword, cfg.DBName, cfg.DBPort)

	db, err := gorm.Open(postgres.Open(psqlInfo), &gorm.Config{
		SkipDefaultTransaction: true,
	})

	if err != nil {
		log.Fatal("Failed to connect with database")
		return nil, err
	}

	//Auto migrate

	db.AutoMigrate(
		domain.Admin{},
		domain.Users{},
		domain.Address{},
		// Product tables
		domain.Category{},
		domain.Product{},
		domain.ProductItem{},
		domain.ProductImage{},
		//Cart
		domain.Cart{},
		domain.CartItems{},
		//WishList
		domain.WishList{},
		domain.WishListItems{},
		//Payment
		domain.PaymentDetails{},
		domain.PaymentMethod{},
		domain.PaymentStatus{},
		// Order tables
		domain.Order{},
		domain.OrderReturn{},
		// Coupon
		domain.Coupon{},
	)
	if err != nil {
		log.Fatal("DB Migration failed")
		return nil, nil
	}
	fmt.Println("DB migration success")
	return db, nil
}
