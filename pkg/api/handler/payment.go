package handler

import (
	"beautify/pkg/domain"
	service "beautify/pkg/usecase/interfaces"
	"beautify/pkg/utils"
	"beautify/pkg/utils/request"
	"beautify/pkg/utils/response"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type PaymentHandler struct {
	PaymentService service.PaymentService
}

func NewPaymentHandler(payUseCase service.PaymentService) *PaymentHandler {
	return &PaymentHandler{
		PaymentService: payUseCase,
	}
}

// AddPaymentMethod godoc
// @Summary Add payment method
// @Description Add payment method
// @Tags Payment
//	@Param			input	body	domain.PaymentMethod	true	"inputs"
// @Failure 400 {object} response.Response{} "Error while fetching data from user"
// @Failure 400 {object} response.Response{}"Can't add payment method"
// @Success 200 {object} response.Response{} "Successfully added payment method"
// @Router /admin/payment/add [post]
func (p *PaymentHandler) AddpaymentMethod(c *gin.Context) {
	var payment domain.PaymentMethod
	if err := c.ShouldBindJSON(&payment); err != nil {
		response := response.ErrorResponse(400, "Error while fetching data from user", err.Error(), payment)
		c.JSON(400, response)
		return
	}
	paymentresp, err1 := p.PaymentService.AddPaymentMethod(c, payment)
	if err1 != nil {
		response := response.ErrorResponse(400, "Can't add payment method", err1.Error(), paymentresp)
		c.JSON(400, response)
		return
	}
	response := response.SuccessResponse(200, "Successfully added payment method", paymentresp)
	c.JSON(200, response)
}

// GetPaymentMethod godoc
// @Summary Get payment method
// @Description Get payment method
// @Tags Payment
// @Param			input	body	domain.PaymentMethod	true	"inputs"
// @Failure 400 {object} response.Response{} "invalid inputs"
// @Failure 400 {object} response.Response{}"invalid inputs"
// @Failure 500 {object} response.Response{}" internal server error"
// @Success 200 {object} response.Response{} "List of payment methods"
// @Router /admin/payment/add [get]
func (p *PaymentHandler) GetPaymentMethods(ctx *gin.Context) {
	count, err1 := utils.StringToUint(ctx.Query("count"))
	if err1 != nil {
		response := response.ErrorResponse(400, "invalid inputs", err1.Error(), nil)
		ctx.JSON(http.StatusBadRequest, response)
		return
	}
	pageNumber, err2 := utils.StringToUint(ctx.Query("page_number"))

	if err2 != nil {
		response := response.ErrorResponse(400, "invalid inputs", err1.Error(), nil)
		ctx.JSON(http.StatusBadRequest, response)
		return
	}
	pagination := request.ReqPagination{
		PageNumber: pageNumber,
		Count:      count,
	}
	payment, err := p.PaymentService.GetPaymentMethods(ctx, pagination)
	if err != nil {
		response := response.ErrorResponse(500, "Internal server error", err.Error(), nil)
		ctx.JSON(500, response)
		return
	}

	response := response.SuccessResponse(200, "List of payment methods", payment)
	ctx.JSON(200, response)
}

// DeletePaymentMethod godoc
// @Summary Delete payment method
// @Description Delete payment method
// @Tags Payment
// @Param			input	body	domain.PaymentMethod	true	"inputs"
// @Failure 400 {object} response.Response{} "Please add id as params"
// @Failure 400 {object} response.Response{}"can't delete payment method"
// @Success 200 {object} response.Response{} "successfully deleted method"
// @Router /admin/payment/delete [delete]
func (p *PaymentHandler) DeleteMethod(c *gin.Context) {
	methodID, err := strconv.Atoi(c.Query("methodID"))
	if err != nil {
		response := response.ErrorResponse(400, "Please add id as params", err.Error(), methodID)
		c.JSON(400, response)
		return
	}
	err1 := p.PaymentService.DeleteMethod(c, uint(methodID))
	if err1 != nil {
		response := response.ErrorResponse(400, "can't delete payment method", err.Error(), "")
		c.JSON(400, response)
		return
	}
	response := response.SuccessResponse(200, "successfully deleted method")
	c.JSON(200, response)
}

// UpdatePaymentMethod godoc
// @Summary Update payment method
// @Description Update payment method
// @Tags Payment
// @Param			input	body	domain.PaymentMethod	true	"inputs"
// @Failure 400 {object} response.Response{} "Error while getting data from admin side"
// @Failure 400 {object} response.Response{}"Can't update data"
// @Success 200 {object} response.Response{} "successfully updated method"
// @Router /admin/payment/update [patch]
func (p *PaymentHandler) UpdatePaymentMethod(c *gin.Context) {
	var payment domain.PaymentMethod
	if err := c.BindJSON(&payment); err != nil {
		response := response.ErrorResponse(400, "Error while getting data from admin side", err.Error(), payment)
		c.JSON(400, response)
		return
	}
	paymentresp, err := p.PaymentService.UpdatePaymentMethod(c, payment)
	if err != nil {
		response := response.ErrorResponse(400, "Can't update data", err.Error(), "")
		c.JSON(400, response)
		return
	}

	response := response.SuccessResponse(200, "successfully updated product", paymentresp)
	c.JSON(200, response)
}
