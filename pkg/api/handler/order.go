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
// @Summary Create order
// @Description Create order
// @Tags Order
//	@Param			input	body	request.Order	true	"inputs"
// @Failure 400 {object} response.Response{}"Failed to get address id"
// @Failure 400 {object} response.Response{}"Failed to get payment method id"
// @Failure 400 {object} response.Response{}"Failed to get total amount"
// @Failure 400 {object} reponse.Response{}"Failed to create order"
// @Success 200 {object} response.Response{}"Successfully created order. Please complete payment"
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
// @Summary Update order
// @Description Update order
// @Tags Order
// @Param			input	body	request.Order	true	"inputs"
// @Failure 400 {object} response.Response{}"error while getting data from users"
// @Failure 400 {object} response.Response{}"error while updating data"
// @Success 200 {object} response.Response{}"Successfully updated order"
// @Router /order/updateOrder [patch]
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

// ListAllOrder godoc
// @Summary List All order
// @Description List All order
// @Tags Order
//	@Param			input	body	request.Order	true	"inputs"
// @Failure 400 {object} response.Response{}"Missing or invalid inputs"
// @Failure 400 {object} response.Response{}"Missing or invalid inputs"
// @Failure 500 {object} response.Response{}"Something went wrong!""
// @Success 200 {object} response.Response{}"Get Orders successfully"
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

// GetAllOrder godoc
// @Summary Get All order
// @Description Get All order
// @Tags Order
//	@Param			input	body	request.Order	true	"inputs"
// @Failure 400 {object} response.Response{}"Missing or invalid inputs"
// @Failure 400 {object} response.Response{}"Missing or invalid inputs"
// @Failure 500 {object} response.Response{}"Something went wrong!""
// @Success 200 {object} response.Response{}"Get Orders successfully"
// @Router /order/getOrder [get]
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

	//userId := utils.GetUserIdFromContext(c)
	orderList, err := o.OrderService.GetAllOrders(c, page)
	if err != nil {
		response := response.ErrorResponse(500, "Something went wrong!", err.Error(), nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	response := response.SuccessResponse(200, "Get Orders successfully", orderList)
	c.JSON(http.StatusOK, response)
}

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
		// Pagination: utils.Pagination{
		// 	Page:     page,
		// 	PageSize: pagesize,
		// },
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
	pdfFilePath := "/home/ubuntu/go/beautify/salesReport/file.pdf"

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
