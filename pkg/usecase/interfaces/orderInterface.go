package interfaces

import (
	"beautify/pkg/domain"
	"beautify/pkg/utils/request"
	"beautify/pkg/utils/response"
	"context"
)

type OrderService interface {
	GetTotalAmount(c context.Context, userid uint) (float64, error)
	CreateOrder(c context.Context, order domain.Order) (response.OrderResponse, error)
	UpdateOrderDetails(c context.Context, uporder request.UpdateOrder) (response.OrderResponse, error)
	ListAllOrders(c context.Context, page request.ReqPagination, userId uint) (orders []response.OrderResponse, err error)
	GetAllOrders(c context.Context, page request.ReqPagination) (orders []response.OrderResponse, err error)
	DeleteOrder(c context.Context, order_id uint) error
	PlaceOrder(c context.Context, order domain.Order) (response.PaymentResponse, error)
	FindPaymentMethodIdByOrderId(c context.Context, order_id uint) (uint, error)
	ValidateCoupon(c context.Context, CouponId uint) (response.CouponResponse, error)
	ApplyDiscount(c context.Context, CouponResponse response.CouponResponse, order_id uint) (int, error)
	FindTotalAmountByOrderId(c context.Context, order_id uint) (float64, error)
	FindPhoneEmailByUserId(c context.Context, usr_id int) (response.PhoneEmailResp, error)
	UpdateOrderStatus(c context.Context, order_id uint) (response.OrderResponse, error)
	GetRazorpayOrder(c context.Context, userID uint, razorPay request.RazorPayReq) (response.ResRazorpayOrder, error)
	UpdateStatusRazorpay(c context.Context, order_id uint) (response.OrderResponse, error)

	ReturnRequest(c context.Context, returnOrder domain.OrderReturn) (response.ReturnResponse, error)
	VerifyOrderID(c context.Context, id uint, orderid uint) error

	SalesReport(c context.Context, page request.ReqPagination, salesData request.ReqSalesReport) ([]response.SalesReport, error)
	CreateUserWallet(userID uint) error

	GetAllPendingReturnRequest(c context.Context, page request.ReqPagination) (ReturnRequests []response.ReturnRequests, err error)
	ClearWalletHistory(ctx context.Context, userId uint) error
}
