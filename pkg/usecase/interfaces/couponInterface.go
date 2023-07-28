package interfaces

import (
	"beautify/pkg/domain"
	"beautify/pkg/utils/request"
	"context"
)

type CouponService interface {
	CreateNewCoupon(ctx context.Context, CouponData request.CreateCoupon) error
	GetAllCoupons(ctx context.Context, page request.ReqPagination) (coupon []domain.Coupon, err error)
	//UpdateCoupon(ctx context.Context, couponData request.UpdateCoupon) error
}
