package interfaces

import (
	"beautify/pkg/domain"
	"beautify/pkg/utils"
	"beautify/pkg/utils/request"
	"context"
)

type CouponRepository interface {
	CreateNewCoupon(ctx context.Context, CouponData request.CreateCoupon) error
	GetCouponBycode(ctx context.Context, code string) (coupon domain.Coupon, err error)
	GetCouponById(ctx context.Context, couponId uint) (coupon domain.Coupon, err error)

	GetAllCoupons(ctx context.Context, page request.ReqPagination) (coupon []domain.Coupon, err error)
	//UpdateCoupon(ctx context.Context, couponData request.UpdateCoupon) error
	ApplyCoupon(ctx context.Context, data utils.ApplyCoupon) (AppliedCoupon utils.ApplyCouponResponse, err error)
}
