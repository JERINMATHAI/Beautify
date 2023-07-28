package interfaces

import (
	"beautify/pkg/domain"
	"beautify/pkg/utils/request"
	"context"
)

type PaymentService interface {
	AddPaymentMethod(c context.Context, payment domain.PaymentMethod) (domain.PaymentMethod, error)
	GetPaymentMethods(ctx context.Context, page request.ReqPagination) (payment []domain.PaymentMethod, err error)
	UpdatePaymentMethod(c context.Context, payment domain.PaymentMethod) (domain.PaymentMethod, error)
	DeleteMethod(c context.Context, id uint) error
}
