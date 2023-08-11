package repository

import (
	"beautify/pkg/domain"
	"beautify/pkg/repository/interfaces"
	"beautify/pkg/utils"
	"beautify/pkg/utils/request"
	"context"
	"errors"
	"time"

	"gorm.io/gorm"
)

type couponDatabase struct {
	DB *gorm.DB
}

func NewCouponRepository(db *gorm.DB) interfaces.CouponRepository {
	return &couponDatabase{DB: db}
}

// Create a coupon
func (c *couponDatabase) CreateNewCoupon(ctx context.Context, CouponData request.CreateCoupon) error {

	query := `INSERT INTO coupons(code,min_order_value,discount_percent,discount_max_amount,valid_till)
	VALUES($1, $2, $3, $4, $5)`
	if err := c.DB.Exec(query, CouponData.Code, CouponData.MinOrderValue,
		CouponData.DiscountPercent, CouponData.DiscountMaxAmount, CouponData.ValidTill).Error; err != nil {
		return err
	}

	return nil
}

//Get all coupons
func (c *couponDatabase) GetAllCoupons(ctx context.Context, page request.ReqPagination) (coupon []domain.Coupon, err error) {
	limit := page.Count
	offset := (page.PageNumber - 1) * limit
	query := `SELECT * FROM coupons ORDER BY id DESC LIMIT ? OFFSET ?`
	if err := c.DB.Raw(query, limit, offset).Scan(&coupon).Error; err != nil {
		return coupon, err
	}
	return coupon, nil
}

// Fetch by coupon code
func (c *couponDatabase) GetCouponBycode(ctx context.Context, code string) (coupon domain.Coupon, err error) {
	query := `SELECT * FROM coupons WHERE code = ?`
	if err := c.DB.Raw(query, code).Scan(&coupon).Error; err != nil {
		return coupon, err
	}
	return coupon, nil
}

//Get coupon by Id
func (c *couponDatabase) GetCouponById(ctx context.Context, couponId uint) (coupon domain.Coupon, err error) {
	query := `SELECT * FROM coupons WHERE id = ?`
	if err := c.DB.Raw(query, couponId).Scan(&coupon).Error; err != nil {
		return coupon, err
	}
	return coupon, nil
}

//Apply coupon to product
func (c *couponDatabase) ApplyCoupon(ctx context.Context, data utils.ApplyCoupon) (AppliedCoupon utils.ApplyCouponResponse, err error) {

	// Get coupon and validate
	couponData, err := c.GetCouponBycode(ctx, data.CouponCode)
	if err != nil {
		return AppliedCoupon, err
	}
	// if couponData.ID == 0 {
	// 	return AppliedCoupon, errors.New("Invalid Coupon")
	// }
	if couponData.ValidTill.Before(time.Now()) {
		return AppliedCoupon, errors.New("Coupon Expired")
	}
	if data.TotalPrice < couponData.MinOrderValue {
		return AppliedCoupon, errors.New("Unable to apply coupon. Minimum order value not reached")
	}
	AppliedCoupon.CouponDiscount = data.TotalPrice * couponData.DiscountPercent / 100
	if AppliedCoupon.CouponDiscount > couponData.DiscountMaxAmount {
		AppliedCoupon.CouponDiscount = couponData.DiscountMaxAmount
	}
	AppliedCoupon.FinalPrice = data.TotalPrice - AppliedCoupon.CouponDiscount
	AppliedCoupon.CouponId = couponData.ID
	AppliedCoupon.CouponCode = couponData.Code

	// check coupon is valid or not
	return AppliedCoupon, nil

}
