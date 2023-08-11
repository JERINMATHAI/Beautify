package handler

import (
	"beautify/pkg/api/auth"
	"beautify/pkg/domain"
	"beautify/pkg/utils"
	"beautify/pkg/utils/request"
	"beautify/pkg/utils/response"
	"net/http"

	service "beautify/pkg/usecase/interfaces"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/copier"
)

type AdminHandler struct {
	adminUseCase service.AdminService
}

func NewAdminHandler(adminService service.AdminService) *AdminHandler {
	return &AdminHandler{adminUseCase: adminService}

}

//	Admin SignUp  godoc
//	@Summary		Admin SignUp
//	@id				AdminSignUp
//	@Description	Create new Admin account
//	@Tags			Admin
//	@Param			input	body	request.SignupUserData	true	"inputs"
//	@Accept			json
//	@Failure		400	{object}	response.Response{}		"Missing or Invalid entry"
//	@Failure		400	{object}	response.Response{}		"Failed to create Admin account"
//	@Success		200	{object}	response.Response{}	"Admin account created successfully"
//	@Router			/admin/signup [post]
func (a *AdminHandler) AdminSignUp(c *gin.Context) {
	var body domain.Admin
	if err := c.ShouldBindJSON(&body); err != nil {
		response := response.ErrorResponse(400, "Missing or invalid entry", err.Error(), body)
		c.JSON(http.StatusBadRequest, response)
	}
	if _, err := a.adminUseCase.Signup(c, body); err != nil {
		response := response.ErrorResponse(400, "Failed to create Admin account", err.Error(), body)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	response := response.SuccessResponse(200, "Admin account created successfully", nil)
	c.JSON(http.StatusOK, response)
}

//	Admin Login  godoc
//	@Summary		Admin Login
//	@id				AdminLogin
//	@Description	Login to Admin account
//	@Tags			Admin
//	@Param			input	body	request.LoginData	true	"inputs"
//	@Accept			json
//	@Failure		400	{object}	response.Response{}		"Missing or Invalid entry"
//	@Failure		400	{object}	response.Response{}		"Failed to login"
//  @Failure		500 {object}	response.Response{}		"Generate JWT failure"
//	@Success		200	{object}	response.Response{}	"Successfully logged in"
//	@Router			/admin/login [post]
func (a *AdminHandler) AdminLogin(c *gin.Context) {
	//Bind login data
	var body request.LoginData
	if err := c.ShouldBindJSON(&body); err != nil {
		response := response.ErrorResponse(400, "Missing or invalid entry", err.Error(), body)
		c.JSON(http.StatusBadRequest, response)
	}
	// validate login data
	var admin domain.Admin
	copier.Copy(&admin, body)
	admin, err := a.adminUseCase.Login(c, admin)
	if err != nil {
		response := response.ErrorResponse(400, "Failed to login", err.Error(), admin)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	// Setup JWT
	if !auth.JwtCookieSetup(c, "admin-auth", admin.ID) {
		response := response.ErrorResponse(500, "Generate JWT failure", err.Error(), nil)
		c.JSON(http.StatusBadRequest, response)
	}
	// Success response
	response := response.SuccessResponse(200, "Successfully logged in", nil)
	c.JSON(http.StatusOK, response)
}

//	Admin Home  godoc
//	@Summary		Admin Home
//	@id				AdminHome
//	@Description	Welcome to Admin Home
//	@Tags			Admin
//	@Success		200	{object}	response.Response{}	"Successfully logged in"
//	@Router			/admin/ [get]
func (a *AdminHandler) AdminHome(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"StatusCode": 200,
		"message":    "Welcome to Admin Home",
	})
}

func (a *AdminHandler) LogoutAdmin(c *gin.Context) {
	c.SetCookie("admin-auth", "", -1, "", "", false, true)
	response := response.SuccessResponse(http.StatusOK, "Log out successful", nil)
	c.JSON(http.StatusOK, response)
}

// ListUsers godoc
// @Summary Get a list of users
// @Description Get a paginated list of users.
// @Tags Users
// @Param count query int false "Number of users to fetch per page"
// @Param page_number query int false "Page number"
//  @Success		200 {object}	response.Response{} "List user successful"
// @Failure 400 {object} response.Response{} "Missing or invalid inputs"
// @Failure 500 {object} response.Response{} "Failed to get all users"
// @Router /users [get]
func (a *AdminHandler) ListUsers(c *gin.Context) {

	count, err1 := utils.StringToUint(c.Query("count"))
	if err1 != nil {
		response := response.ErrorResponse(400, "Missing or invalid inputs", err1.Error(), nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	pageNumber, err2 := utils.StringToUint(c.Query("page_number"))

	if err2 != nil {
		response := response.ErrorResponse(400, "Missing or invalid inputs", err1.Error(), nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	pagination := request.ReqPagination{
		PageNumber: pageNumber,
		Count:      count,
	}

	users, err := a.adminUseCase.GetAllUser(c, pagination)
	if err != nil {
		respone := response.ErrorResponse(500, "Failed to get all users", err.Error(), nil)
		c.JSON(http.StatusInternalServerError, respone)
		return
	}

	// check there is no user
	if len(users) == 0 {
		response := response.SuccessResponse(200, "Oops!...No user to show", nil)
		c.JSON(http.StatusOK, response)
		return
	}

	response := response.SuccessResponse(200, "List user successful", users)
	c.JSON(http.StatusOK, response)

}

//	Admin Block Users  godoc
//	@Summary		Admin Block Users
//	@id				BlockUser
//	@Description	Block users in admin side
//	@Tags			Admin
//	@Param			input	body	request.Block	true	"inputs"
//	@Accept			json
//	@Failure		400	{object}	response.Response{}		"Invalid inputs"
//	@Failure		400	{object}	response.Response{}		"Failed to change user block_status"
//  @Success		200 {object}	response.Response{}		"Successfully changed user block_status"
//	@Router			/admin/users/block [patch]
func (a *AdminHandler) BlockUser(ctx *gin.Context) {
	var body request.Block
	if err := ctx.ShouldBindJSON(&body); err != nil {
		response := response.ErrorResponse(400, "Invalid input", err.Error(), body)
		ctx.JSON(http.StatusBadRequest, response)
		return
	}
	// if successfully blocked or unblock user then response 200
	err := a.adminUseCase.BlockUser(ctx, body.UserID)
	if err != nil {
		response := response.ErrorResponse(400, "Failed to change user block_status", err.Error(), nil)
		ctx.JSON(http.StatusBadRequest, response)
		return
	}
	response := response.SuccessResponse(200, "Successfully changed user block_status", body.UserID)
	ctx.JSON(http.StatusOK, response)
}
