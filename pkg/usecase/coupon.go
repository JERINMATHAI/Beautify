package usecase

import (
	"beautify/pkg/domain"
	"beautify/pkg/repository/interfaces"
	service "beautify/pkg/usecase/interfaces"
	"beautify/pkg/utils/request"
	"context"
)

type couponUseCase struct {
	couponRepository interfaces.CouponRepository
}

func NewCouponUseCase(CouponRepo interfaces.CouponRepository) service.CouponService {
	return &couponUseCase{couponRepository: CouponRepo}
}

//Create coupon
func (c *couponUseCase) CreateNewCoupon(ctx context.Context, CouponData request.CreateCoupon) error {
	if err := c.couponRepository.CreateNewCoupon(ctx, CouponData); err != nil {
		return err
	}
	return nil
}

//Get all coupons
func (c *couponUseCase) GetAllCoupons(ctx context.Context, page request.ReqPagination) (coupon []domain.Coupon, err error) {
	if coupon, err = c.couponRepository.GetAllCoupons(ctx, page); err != nil {
		return nil, err
	}
	return coupon, nil
}

// func (c *couponUseCase) UpdateCoupon(ctx context.Context, couponData request.UpdateCoupon) error {
// 	if err := c.couponRepository.UpdateCoupon(ctx, couponData); err != nil {
// 		return err
// 	}
// 	return nil
// }
