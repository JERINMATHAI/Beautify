package usecase

import (
	"beautify/pkg/config"
	"beautify/pkg/domain"
	"beautify/pkg/repository/interfaces"
	service "beautify/pkg/usecase/interfaces"
	"beautify/pkg/utils"
	"beautify/pkg/utils/request"
	"beautify/pkg/utils/response"
	"context"
	"errors"
	"fmt"
)

// type OrderUseCase struct {
// 	OrderRepository interfaces.OrderRepository
// 	UserRepository  interfaces.UserRepository
// 	PayMentRepo     interfaces.PaymentRepository
// 	CouponRepo      interfaces.CouponRepository
// }

type OrderUseCase struct {
	OrderRepository interfaces.OrderRepository
}

// func NewOrderUseCase(repo interfaces.OrderRepository, UserRepo interfaces.UserRepository,
// 	payMentRepo interfaces.PaymentRepository, couponRepo interfaces.CouponRepository) service.OrderService {
// 	return &OrderUseCase{
// 		OrderRepository: repo,
// 		UserRepository:  UserRepo,
// 		PayMentRepo:     payMentRepo,
// 		CouponRepo:      couponRepo}
// }

func NewOrderUseCase(repo interfaces.OrderRepository) service.OrderService {
	return &OrderUseCase{
		OrderRepository: repo,
	}
}

func (o *OrderUseCase) GetTotalAmount(c context.Context, userid uint) (float64, error) {
	var total_amount float64
	total_amount = 0
	cart, err := o.OrderRepository.GetTotalAmount(c, int(userid))
	if err != nil {
		return 0, err
	}

	for _, c := range cart {
		total_amount = total_amount + float64(c.Total)
	}
	return total_amount, nil
}

func (o *OrderUseCase) CreateOrder(c context.Context, order domain.Order) (response.OrderResponse, error) {
	//Checking whether the payment id exist
	_, err := o.OrderRepository.FindPaymentMethodById(c, uint(order.PaymentMethodID))

	if err != nil {
		return response.OrderResponse{}, errors.New("payment method doesn't exists")
	}
	orderresp, err := o.OrderRepository.CreateOrder(c, order)
	if err != nil {
		return response.OrderResponse{}, err
	}
	return orderresp, nil
}
func (o *OrderUseCase) UpdateOrderDetails(c context.Context, uporder request.UpdateOrder) (response.OrderResponse, error) {
	//Checking whether the payment id exist
	_, err := o.OrderRepository.FindPaymentMethodById(c, uporder.PaymentMethodID)

	if err != nil {
		return response.OrderResponse{}, errors.New("payment method doesn't exists")
	}
	orderup, err := o.OrderRepository.UpdateOrderDetails(c, uporder)
	if err != nil {
		return response.OrderResponse{}, err
	}
	return orderup, nil
}

//List order for user
func (o *OrderUseCase) ListAllOrders(c context.Context, page request.ReqPagination, userId uint) (orders []response.OrderResponse, err error) {

	return o.OrderRepository.ListAllOrders(c, page, userId)

}

//List order for admin
func (o *OrderUseCase) GetAllOrders(c context.Context, page request.ReqPagination) (orders []response.OrderResponse, err error) {

	return o.OrderRepository.GetAllOrders(c, page)

}

func (o *OrderUseCase) DeleteOrder(c context.Context, order_id uint) error {
	err := o.OrderRepository.DeleteOrder(c, order_id)
	if err != nil {
		return err
	}
	return nil
}

func (o *OrderUseCase) PlaceOrder(c context.Context, order domain.Order) (response.PaymentResponse, error) {
	method_Id, err := o.OrderRepository.FindPaymentMethodIdByOrderId(c, order.Order_Id)
	if err != nil {
		return response.PaymentResponse{}, err
	}
	if method_Id == 1 {
		order.Order_Status = "order confirmed"
	} else {
		order.Order_Status = "order confirmed payment pending"
	}
	paymentresp, err := o.OrderRepository.PlaceOrder(c, order)
	if err != nil {
		return response.PaymentResponse{}, err
	}
	return paymentresp, nil
}

func (o *OrderUseCase) FindPaymentMethodIdByOrderId(c context.Context, order_id uint) (uint, error) {
	method_id, err := o.OrderRepository.FindPaymentMethodIdByOrderId(c, order_id)
	if err != nil {
		return 0, err
	}
	return method_id, nil
}

func (o *OrderUseCase) ValidateCoupon(c context.Context, CouponId uint) (response.CouponResponse, error) {
	err := o.OrderRepository.FindCouponById(c, CouponId)
	if err != nil {
		return response.CouponResponse{}, err
	}
	couponResp, err := o.OrderRepository.ValidateCoupon(c, CouponId)
	if err != nil {
		return response.CouponResponse{}, err
	}
	return couponResp, nil
}

func (o *OrderUseCase) ApplyDiscount(c context.Context, CouponResponse response.CouponResponse, order_id uint) (int, error) {
	order, err := o.OrderRepository.ApplyDiscount(c, order_id)
	if err != nil {
		return 0, nil
	}

	totalamnt := order.Total_Amount - float64(CouponResponse.Discount)

	return int(totalamnt), nil

}

func (o *OrderUseCase) UpdateOrderStatus(c context.Context, order_id uint) (response.OrderResponse, error) {
	order_status := "Order confirmed"
	orderResp, err := o.OrderRepository.UpdateOrderStatus(c, order_id, order_status)
	if err != nil {
		return response.OrderResponse{}, err
	}
	return orderResp, nil
}

func (o *OrderUseCase) FindTotalAmountByOrderId(c context.Context, order_id uint) (float64, error) {
	totalAmount, err := o.OrderRepository.FindTotalAmountByOrderId(c, order_id)
	if err != nil {
		return 0, err
	}
	return totalAmount, nil
}

func (o *OrderUseCase) FindPhoneEmailByUserId(c context.Context, usr_id int) (response.PhoneEmailResp, error) {
	phnEmail, err := o.OrderRepository.FindPhoneEmailByUserId(c, usr_id)
	if err != nil {
		return response.PhoneEmailResp{}, err
	}
	return phnEmail, nil
}

func (o *OrderUseCase) GetRazorpayOrder(c context.Context, userID uint, razorPay request.RazorPayReq) (response.ResRazorpayOrder, error) {
	var razorpayOrder response.ResRazorpayOrder

	//razorpay amount is caluculate on pisa for india so make the actual price into paisa
	razorPayAmount := uint(razorPay.Total_Amount * 100)

	razopayOrderId, err := utils.GenerateRazorpayOrder(razorPayAmount, "test reciept")
	if err != nil {
		return razorpayOrder, err
	}
	fmt.Println(razopayOrderId)
	// set all details on razopay order
	razorpayOrder.AmountToPay = uint(razorPay.Total_Amount)

	razorpayOrder.RazorpayKey, _ = config.GetRazorPayConfig()

	razorpayOrder.UserID = userID
	razorpayOrder.RazorpayOrderID = razopayOrderId

	razorpayOrder.Email = razorPay.Email
	razorpayOrder.Phone = razorPay.Phone

	return razorpayOrder, nil
}

func (o *OrderUseCase) UpdateStatusRazorpay(c context.Context, order_id uint) (response.OrderResponse, error) {
	order_status := "Order confirmed"
	payment_status := "Payment Done"
	delivery_status := "Order delivered successfully"
	orderResp, err := o.OrderRepository.UpdateStatusRazorpay(c, order_id, order_status, payment_status, delivery_status)
	if err != nil {
		return response.OrderResponse{}, err
	}
	return orderResp, nil
}

func (o *OrderUseCase) ReturnRequest(c context.Context, returnOrder domain.OrderReturn) (response.ReturnResponse, error) {
	returnResp, err := o.OrderRepository.ReturnRequest(c, returnOrder)
	if err != nil {
		return response.ReturnResponse{}, err
	}
	return returnResp, nil
}

func (o *OrderUseCase) VerifyOrderID(c context.Context, id uint, orderid uint) error {
	err := o.OrderRepository.VerifyOrderID(c, id, orderid)
	if err != nil {
		return err
	}
	return nil
}
func (o *OrderUseCase) SalesReport(c context.Context, page request.ReqPagination, salesData request.ReqSalesReport) ([]response.SalesReport, error) {
	salesReport, err := o.OrderRepository.SalesReport(c, page, salesData)
	if err != nil {
		return []response.SalesReport{}, err
	}
	return salesReport, nil
}

func (o *OrderUseCase) CreateUserWallet(userID uint) error {
	wallet, err := o.OrderRepository.FindUserWallet(userID)
	if err != nil {
		return fmt.Errorf("Failed to check user wallet :%s", err)
	}

	if wallet.ID != 0 {
		return fmt.Errorf("User already have wallet")
	}
	Wallet, err := o.OrderRepository.InitializeNewWallet(userID)
	if err != nil {
		return fmt.Errorf("Failed to initialize wallet for user %d : %s", userID, err)
	}
	if Wallet.ID == 0 {
		return fmt.Errorf("Failed to verify new wallet for user id  %d", userID)
	}
	return nil
}
