package handler

import (
	"beautify/pkg/domain"
	service "beautify/pkg/usecase/interfaces"
	"beautify/pkg/utils"
	"beautify/pkg/utils/request"
	"beautify/pkg/utils/response"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/jung-kurt/gofpdf"
)

type OrderHandler struct {
	OrderService service.OrderService
}

func NewOrderHandler(orderUseCase service.OrderService) *OrderHandler {
	return &OrderHandler{
		OrderService: orderUseCase,
	}
}

// CreateOrder godoc
// @Summary Create an order
// @Description Create an order with specified address and payment method
// @Tags Order
// @Param address_id query int true "Address ID"
// @Param paymentmethod_id query int true "Payment Method ID"
// @Success 200 {object} response.Response{} "Successfully created order. Please complete payment"
// @Failure 400 {object} response.Response{} "Failed to get address id or Failed to get payment method id or Failed to get total amount or Failed to create order"
// @Router /order/createOrder [post]
func (o *OrderHandler) CreateOrder(c *gin.Context) {
	var order domain.Order
	addressID, err := strconv.Atoi(c.Query("address_id"))
	if err != nil {
		response := response.ErrorResponse(400, "Failed to get address id", err.Error(), order)
		c.JSON(400, response)
		return
	}
	PaymentMetodId, err := strconv.Atoi(c.Query("paymentmethod_id"))
	if err != nil {
		response := response.ErrorResponse(400, "Failed to get payment method id", err.Error(), order)
		c.JSON(400, response)
		return
	}

	userId := utils.GetUserIdFromContext(c)
	totalAmount, err := o.OrderService.GetTotalAmount(c, userId)
	if err != nil {
		response := response.ErrorResponse(400, "Failed to get total amount", err.Error(), order)
		c.JSON(400, response)
		return
	}
	order.Total_Amount = totalAmount
	order.Address_Id = addressID
	order.PaymentMethodID = PaymentMetodId
	order.Payment_Status = "Pending"
	order.Order_Status = "Order Created"
	order.DeliveryStatus = "Pending"
	order.User_Id = userId

	orderResp, err := o.OrderService.CreateOrder(c, order)
	if err != nil {
		response := response.ErrorResponse(400, "Failed to create order", err.Error(), "Try Again")
		c.JSON(400, response)
		return
	}
	response := response.SuccessResponse(200, "Successfully created order. Please complete payment", orderResp)
	c.JSON(200, response)
}

// UpdateOrder godoc
// @Summary Update order details
// @Description Update order details based on provided data
// @Tags Order
// @Param body body request.UpdateOrder true "Update Order Details"
// @Success 200 {object} response.Response{} "Successfully updated order"
// @Failure 400 {object} response.Response{} "Error while getting data from users or Error while updating data"
// @Router /order/updateOrder [put]
func (o *OrderHandler) UpdateOrder(c *gin.Context) {
	var UpdateOrderDetails request.UpdateOrder
	if err := c.ShouldBindJSON(&UpdateOrderDetails); err != nil {
		response := response.ErrorResponse(400, "error while getting data from users", err.Error(), UpdateOrderDetails)
		c.JSON(400, response)
		return
	}
	uporder, err := o.OrderService.UpdateOrderDetails(c, UpdateOrderDetails)
	if err != nil {
		response := response.ErrorResponse(400, "error while updating data", err.Error(), UpdateOrderDetails)
		c.JSON(400, response)
		return
	}
	response := response.SuccessResponse(200, "Successfully updated order", uporder)
	c.JSON(200, response)
}

// ListAllOrders godoc
// @Summary List all orders
// @Description List all orders for the authenticated user
// @Tags Order
// @Param page_number query uint true "Page Number"
// @Param count query uint true "Count of items per page"
// @Success 200 {object} response.Response{} "Get Orders successfully"
// @Failure 400 {object} response.Response{} "Missing or invalid inputs"
// @Failure 500 {object} response.Response{} "Something went wrong!"
// @Router /order/listOrder [get]
func (o *OrderHandler) ListAllOrders(c *gin.Context) {
	var page request.ReqPagination
	count, err0 := utils.StringToUint(c.Query("count"))
	if err0 != nil {
		response := response.ErrorResponse(400, "Missing or invalid inputs", err0.Error(), nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	page_number, err1 := utils.StringToUint(c.Query("page_number"))
	if err1 != nil {
		response := response.ErrorResponse(400, "Missing or invalid inputs", err0.Error(), nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	page.PageNumber = page_number
	page.Count = count
	userId := utils.GetUserIdFromContext(c)
	orderList, err := o.OrderService.ListAllOrders(c, page, userId)
	if err != nil {
		response := response.ErrorResponse(500, "Something went wrong!", err.Error(), nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	response := response.SuccessResponse(200, "Get Orders successfully", orderList)
	c.JSON(http.StatusOK, response)
}

// GetAllOrders godoc
// @Summary Get all orders
// @Description Get a list of all orders
// @Tags Order
// @Param page_number query uint true "Page Number"
// @Param count query uint true "Count of items per page"
// @Success 200 {object} response.Response{} "Get Orders successfully"
// @Failure 400 {object} response.Response{} "Missing or invalid inputs"
// @Failure 500 {object} response.Response{} "Something went wrong!"
// @Router /order/all [get]
func (o *OrderHandler) GetAllOrders(c *gin.Context) {
	var page request.ReqPagination
	count, err0 := utils.StringToUint(c.Query("count"))
	if err0 != nil {
		response := response.ErrorResponse(400, "Missing or invalid inputs", err0.Error(), nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	page_number, err1 := utils.StringToUint(c.Query("page_number"))
	if err1 != nil {
		response := response.ErrorResponse(400, "Missing or invalid inputs", err0.Error(), nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	page.PageNumber = page_number
	page.Count = count
	orderList, err := o.OrderService.GetAllOrders(c, page)
	if err != nil {
		response := response.ErrorResponse(500, "Something went wrong!", err.Error(), nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	response := response.SuccessResponse(200, "Get Orders successfully", orderList)
	c.JSON(http.StatusOK, response)
}

// CancelOrder godoc
// @Summary Cancel an order
// @Description Cancel an order based on the provided order ID
// @Tags Order
// @Param order_id query int true "Order ID"
// @Success 200 {object} response.Response{} "Successfully deleted order"
// @Failure 400 {object} response.Response{} "Please add id as params or Can't delete order"
// @Router /order/cancelOrder [delete]
func (o *OrderHandler) CancelOrder(c *gin.Context) {
	order_id, err := strconv.Atoi(c.Query("order_id"))
	if err != nil {
		response := response.ErrorResponse(400, "Please add id as params", err.Error(), order_id)
		c.JSON(400, response)
		return
	}
	err1 := o.OrderService.DeleteOrder(c, uint(order_id))
	if err1 != nil {
		response := response.ErrorResponse(400, "can't delete order", err.Error(), "")
		c.JSON(400, response)
		return
	}
	response := response.SuccessResponse(200, "successfully deleted order")
	c.JSON(200, response)
}

// PlaceOrder godoc
// @Summary Place an order
// @Description Place an order and apply a coupon if available
// @Tags Order
// @Param order_id query int true "Order ID"
// @Param coupon_id query int false "Coupon ID (optional)"
// @Success 200 {object} response.Response{} "Successfully confirmed order, complete payment process on delivery or placed order with complete payment details"
// @Failure 400 {object} response.Response{} "Invalid coupon or Add more quantity or Failed to place order"
// @Router /order/placeOrder [post]
func (o *OrderHandler) PlaceOrder(c *gin.Context) {
	var placeorder request.PlaceOrderRequest
	var order domain.Order
	order_id, _ := strconv.Atoi(c.Query("order_id"))
	coupon_id, _ := strconv.Atoi(c.Query("coupon_id"))

	placeorder.CouponId = coupon_id
	placeorder.OrderId = order_id
	order.Order_Id = uint(order_id)
	order.Applied_Coupon_id = uint(coupon_id)
	order.OrderDate = time.Now()
	couponResp, err := o.OrderService.ValidateCoupon(c, order.Applied_Coupon_id)
	if err != nil {
		response := response.ErrorResponse(400, "Invalid coupon", err.Error(), "try with a valid coupon")
		c.JSON(400, response)
		return
	} else {
		totalamnt, err := o.OrderService.ApplyDiscount(c, couponResp, uint(order_id))
		if err != nil {
			response := response.ErrorResponse(400, "Add more quantity", err.Error(), "try again")
			c.JSON(400, response)
			return
		}
		order.Total_Amount = float64(totalamnt)
	}
	paymentResp, err := o.OrderService.PlaceOrder(c, order)
	if err != nil {
		response := response.ErrorResponse(400, "failed to place order", err.Error(), "")
		c.JSON(400, response)
		return
	}
	if paymentResp.PaymentMethodId == "1" {
		response := response.SuccessResponse(200, "Successfully confirmed order complete payment process on delivery", paymentResp)
		c.JSON(200, response)
		return
	}
	response := response.SuccessResponse(200, "Successfully  placed order complete payment details", paymentResp)
	c.JSON(200, response)
}

// CheckOut godoc
// @Summary Process order checkout
// @Description Process order checkout and generate payment information
// @Tags Order
// @Param order_id query int true "Order ID"
// @Success 200 {object} response.Response{} "Successfully confirmed order or generated payment information"
// @Failure 400 {object} response.Response{} "Please add order_id as params or Failed to find payment method or Failed to place order or error while getting id from cookie or error while getting total amount or error while getting details or failed to create razorpay order"
// @Router /order/payment [post]
func (o *OrderHandler) CheckOut(c *gin.Context) {
	var razorPay request.RazorPayReq
	order_id, err := strconv.Atoi(c.Query("order_id"))
	if err != nil {
		response := response.ErrorResponse(400, "Please add order_id  as params", err.Error(), "")
		c.JSON(400, response)
		return
	}
	payment_method_id, err := o.OrderService.FindPaymentMethodIdByOrderId(c, uint(order_id))
	if err != nil {
		response := response.ErrorResponse(400, "Failed to find payment method", err.Error(), "")
		c.JSON(400, response)
		return
	}
	if payment_method_id == 1 {
		orderREsp, err := o.OrderService.UpdateOrderStatus(c, uint(order_id))
		if err != nil {
			response := response.ErrorResponse(400, "Failed to place order", err.Error(), "")
			c.JSON(400, response)
			return
		}
		response := response.SuccessResponse(200, "Successfully  confirmed order", orderREsp)
		c.JSON(200, response)
		return
	} else {
		//userId := utils.GetUserIdFromContext(c)
		//orderList, err := o.OrderService.ListAllOrders(c, page, userId)
		//id, err := middlware.GetId(c, "User_Authorization")
		userId := utils.GetUserIdFromContext(c)

		if err != nil {
			response := response.ErrorResponse(400, "error while getting id from cookie", err.Error(), userId)
			c.JSON(400, response)
			return
		}
		totalAmount, err := o.OrderService.FindTotalAmountByOrderId(c, uint(order_id))
		if err != nil {
			response := response.ErrorResponse(400, "error while getting total amount", err.Error(), userId)
			c.JSON(400, response)
			return
		}
		razorPay.Total_Amount = totalAmount
		phnEmail, err := o.OrderService.FindPhoneEmailByUserId(c, int(userId))
		if err != nil {
			response := response.ErrorResponse(400, "error while getting details", err.Error(), userId)
			c.JSON(400, response)
			return
		}
		razorPay.Email = phnEmail.Email
		razorPay.Phone = phnEmail.Phone

		razorpayOrder, err := o.OrderService.GetRazorpayOrder(c, uint(userId), razorPay)
		if err != nil {
			response := response.ErrorResponse(400, "failed to create razorpay order ", err.Error(), nil)
			c.JSON(400, response)
			return
		}
		c.HTML(200, "payment.html", razorpayOrder)
		o.OrderService.UpdateStatusRazorpay(c, uint(order_id))
	}

}

// ReturnOrder godoc
// @Summary Request to return an order
// @Description Request to return an order based on the provided order ID and reason
// @Tags Order
// @Param orderId query int true "Order ID"
// @Param reason query string true "Reason for return"
// @Success 200 {object} response.Response{} "Successfully requested to return products"
// @Failure 400 {object} response.Response{} "Please add order id as params or Error while getting id from cookie or Invalid order_id or Failed to find refund amount or Failed to return order"
// @Router /return/product [post]
func (o *OrderHandler) ReturnOrder(c *gin.Context) {
	var returnOrder domain.OrderReturn
	order_id, err := strconv.Atoi(c.Query("orderId"))
	if err != nil {
		response := response.ErrorResponse(400, "Please add order id as params", err.Error(), "")
		c.JSON(400, response)
		return
	}
	userId := utils.GetUserIdFromContext(c)
	if err != nil {
		response := response.ErrorResponse(400, "error while getting id from cookie", err.Error(), " ")
		c.JSON(400, response)
		return
	}
	err1 := o.OrderService.VerifyOrderID(c, uint(userId), uint(order_id))
	if err1 != nil {
		response := response.ErrorResponse(400, "invalid order_id", err1.Error(), userId)
		c.JSON(400, response)
		return
	}

	returnOrder.OrderID = uint(order_id)
	returnOrder.RequestDate = time.Now()
	returnOrder.ReturnReason = c.Query("reason")
	returnOrder.ReturnStatus = "Return Requested"
	//finding total amount by orderid
	total_amount, err := o.OrderService.FindTotalAmountByOrderId(c, uint(order_id))
	if err != nil {
		response := response.ErrorResponse(400, "Failed to find refund amount", err.Error(), "")
		c.JSON(400, response)
		return
	}
	returnOrder.RefundAmount = total_amount
	returnResp, err := o.OrderService.ReturnRequest(c, returnOrder)
	if err != nil {
		response := response.ErrorResponse(400, "failed to return order", err.Error(), "")
		c.JSON(400, response)
		return
	}

	response := response.SuccessResponse(200, "successfully requsted to return products", returnResp)
	c.JSON(200, response)

}

// SalesReport godoc
// @Summary Generate sales report PDF
// @Description Generate a sales report in PDF format based on provided filters
// @Tags Order
// @Param count query uint true "Count of items per page"
// @Param page_number query uint true "Page Number"
// @Param startDate query string true "Start Date (YYYY-MM-DD)"
// @Param endDate query string true "End Date (YYYY-MM-DD)"
// @Success 200 {object} response.Response{} "Successfully generated pdf"
// @Failure 400 {object} response.Response{} "Invalid inputs or Please add start date as params or Please add end date as params or There is no sales report on this period"
// @Failure 500 {object} response.Response{} "Failed to generate PDF"
// @Router /admin/dashboard/salesReport [get]
func (o *OrderHandler) SalesReport(c *gin.Context) {
	count, err1 := utils.StringToUint(c.Query("count"))
	if err1 != nil {
		response := response.ErrorResponse(400, "invalid inputs", err1.Error(), nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	pageNumber, err2 := utils.StringToUint(c.Query("page_number"))
	if err2 != nil {
		response := response.ErrorResponse(400, "invalid inputs", err1.Error(), nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	pagination := request.ReqPagination{
		PageNumber: pageNumber,
		Count:      count,
	}
	sDate, err := utils.StringToTime(c.Query("startDate"))
	if err != nil {
		response := response.ErrorResponse(400, "Please add start date as params", err.Error(), "")
		c.JSON(400, response)
		return
	}
	eDate, err := utils.StringToTime(c.Query("endDate"))
	if err != nil {
		response := response.ErrorResponse(400, "Please add end date as params", err.Error(), "")
		c.JSON(400, response)
		return
	}
	salesData := request.ReqSalesReport{
		StartDate: sDate,
		EndDate:   eDate,
	}
	salesReport, _ := o.OrderService.SalesReport(c, pagination, salesData)
	if salesReport == nil {
		response := response.ErrorResponse(400, "There is no sales report on this period", " ", " ")
		c.JSON(400, response)
		return
	}
	fmt.Println(salesReport)
	// Create a new PDF document
	pdf := gofpdf.New("P", "mm", "A4", "")

	// Add a new page
	pdf.AddPage()

	// Set the font and font size
	pdf.SetFont("Arial", "i", 12)

	// Add the report title
	pdf.CellFormat(0, 15, "Sales Report", "", 0, "C", false, 0, "")
	pdf.Ln(10)
	// Add the sales report data to the PDF
	i := 1
	for _, sale := range salesReport {

		pdf.CellFormat(0, 15, fmt.Sprint(i)+".", "", 0, "L", false, 0, "")
		pdf.Ln(10)
		pdf.Cell(0, 10, fmt.Sprintf("User ID: %d", sale.UserID))
		pdf.Ln(10)
		pdf.Cell(0, 10, fmt.Sprintf("Name: %s", sale.Name))
		pdf.Ln(10)
		pdf.Cell(0, 10, fmt.Sprintf("Email: %s", sale.Email))
		pdf.Ln(10)
		pdf.Cell(0, 10, fmt.Sprintf("Order Date: %v", sale.OrderDate))
		pdf.Ln(10)
		pdf.Cell(0, 10, fmt.Sprintf("TotalPrice: %v", sale.OrderTotalPrice))
		pdf.Ln(10)
		pdf.Cell(0, 10, fmt.Sprintf("Order Status: %s", sale.OrderStatus))
		pdf.Ln(10)
		pdf.Cell(0, 10, fmt.Sprintf("Payment status: %v", sale.PaymentStatus))
		pdf.Ln(10)
		// pdf.Cell(0, 10, fmt.Sprintf("Payment Type: %v", sale.PaymentType))
		// pdf.Ln(10)

		// Move to the next line
		pdf.Ln(10)
		i++
	}

	// Generate a temporary file path for the PDF
	pdfFilePath := "/home/ubuntu/Beautify/salesReport/file.pdf"

	// Save the PDF to the temporary file path
	err = pdf.OutputFileAndClose(pdfFilePath)
	if err != nil {
		response := response.ErrorResponse(500, "Failed to generate PDF", err.Error(), "")
		c.JSON(500, response)
		return
	}

	// Set the appropriate headers for the file download
	c.Header("Content-Disposition", "attachment; filename=sales_report.pdf")
	c.Header("Content-Type", "application/pdf")

	// Serve the PDF file for download
	c.File(pdfFilePath)

	response := response.SuccessResponse(200, "Successfully generated pdf", " ")
	c.JSON(200, response)
}

func (o *OrderHandler) CreateUserWallet(c *gin.Context) {
	userID := utils.GetUserIdFromContext(c)

	err := o.OrderService.CreateUserWallet(userID)
	if err != nil {
		response := response.ErrorResponse(500, "Failed", err.Error(), nil)
		c.JSON(http.StatusInternalServerError, response)
		return
	}

	response := response.SuccessResponse(200, "Success, wallet initialized", nil, nil)
	c.JSON(http.StatusOK, response)
}
