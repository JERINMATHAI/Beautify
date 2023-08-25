package handler

import (
	"beautify/pkg/api/auth"
	"beautify/pkg/domain"
	"beautify/pkg/usecase/interfaces"
	"beautify/pkg/utils"
	"beautify/pkg/utils/request"
	"beautify/pkg/utils/response"
	"beautify/pkg/verify"
	"errors"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/copier"
)

type UserHandler struct {
	userService interfaces.UserService
}

func NewUserHandler(userUsecase interfaces.UserService) *UserHandler {
	return &UserHandler{userService: userUsecase}
}

//	User Signup  godoc
//	@Summary		User Signup
//	@id				User Signup
//	@Description	Create new user account
//	@Tags			User
//	@Param			input	body	request.SignupUserData	true	"inputs"
//	@Accept			json
//	@Failure		400	{object}	response.Response{}		"Missing or Invalid entry"
//	@Success		200	{object}	response.Response{}	"Successfully logged in"
//	@Router			/signup [post]
func (u *UserHandler) UserSignup(c *gin.Context) {
	var body request.SignupUserData
	if err := c.ShouldBindJSON(&body); err != nil {
		response := response.ErrorResponse(400, "Missing or invalid entry", err.Error(), body)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	var user domain.Users
	if err := copier.Copy(&user, body); err != nil {
		fmt.Println("Copy failed")
	}
	usr, err := u.userService.SignUp(c, user)
	if err != nil {
		response := response.ErrorResponse(400, "User already exist", err.Error(), body)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	response := response.SuccessResponse(200, "Account created successfuly", usr)
	c.JSON(http.StatusOK, response)
}

//	User Login  godoc
//	@Summary		User login
//	@id				User login
//	@Description	Login to user account
//	@Tags			User
//	@Param			input	body	request.LoginData	true	"inputs"
//	@Accept			json
//	@Failure		400	{object}	response.Response{}		"Missing or Invalid entry"
//  @Failure 		400 {object}	response.Response{} 	"Failed to login"
//	@Failure		500	{object}	response.Response{}		"Generate JWT failure"
//	@Success		200 {object}	response.Response{}		"OTP send to your mobile number!"
//	@Router			/login [post]
func (u *UserHandler) LoginSubmit(c *gin.Context) {
	var body request.LoginData
	if err := c.ShouldBindJSON(&body); err != nil {
		response := response.ErrorResponse(400, "Missing or invalid entry", err.Error(), body)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	if body.Email == "" && body.Password == "" && body.UserName == "" {
		_ = errors.New("Please enter user_name and password")
		response := "Field should not be empty"
		c.JSON(http.StatusBadRequest, response)
		return
	}
	var user domain.Users
	copier.Copy(&user, body)
	// validate login data
	user, err := u.userService.Login(c, user)
	if err != nil {
		response := response.ErrorResponse(400, "Failed to login", err.Error(), user)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	// Setup JWT
	if !auth.JwtCookieSetup(c, "user-auth", user.ID) {
		response := response.ErrorResponse(500, "Generate JWT failure", err.Error(), nil)
		c.JSON(http.StatusBadRequest, response)
	}
	// Success response
	response := response.SuccessResponse(200, "OTP send to your mobile number!", user.Phone)
	c.JSON(http.StatusOK, response)
}

//	User OTP Verification  godoc
//	@Summary		User OTP Verification
//	@id				User OTP Verification
//	@Description	OTP Verification to user account
//	@Tags			User
//	@Param			input	body	request.OTPVerify	true	"inputs"
//	@Accept			json
//	@Failure		400	{object}	response.Response{}		"Missing or Invalid entry"
//  @Failure 		500 {object}	response.Response{} 	"User not registered"
//	@Failure		500	{object}	response.Response{}		"Failed to login"
//	@Success		200 {object}	response.Response{}		"Successfully logged in"
//	@Router			/login/otp-verify [post]
func (u *UserHandler) UserOTPVerify(c *gin.Context) {
	var body request.OTPVerify
	if err := c.ShouldBindJSON(&body); err != nil {
		response := response.ErrorResponse(400, "Missing or invalid entry", err.Error(), body)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	var user = domain.Users{
		ID: body.UserID,
	}
	usr, err := u.userService.OTPLogin(c, user)
	if err != nil {
		response := response.ErrorResponse(500, "User not registered", err.Error(), user)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	// Verify OTP
	err = verify.TwilioVerifyOTP("+91"+usr.Phone, body.OTP)
	if err != nil {
		response := gin.H{"error": err.Error()}
		c.JSON(http.StatusBadRequest, response)
		return
	}
	// Setup JWT
	ok := auth.JwtCookieSetup(c, "user-auth", usr.ID)
	if !ok {
		response := response.ErrorResponse(500, "Failed to login", "", nil)
		c.JSON(http.StatusInternalServerError, response)
		return

	}
	response := response.SuccessResponse(200, "Successfuly logged in!", nil)
	c.JSON(http.StatusOK, response)
}

//	User Home  godoc
//	@Summary		User Home
//	@id				User Home
//	@Description	Home page for user
//	@Tags			User
//	@Accept			json
//	@Success		200 {object}	response.Response{}		"Welcome to home !"
//	@Router			/ [get]
func (u *UserHandler) Home(c *gin.Context) {
	response := response.SuccessResponse(200, "Welcome to home !", nil)
	c.JSON(http.StatusOK, response)
}

//	User Logout  godoc
//	@Summary		User Logout
//	@id				User Logout
//	@Description	Logout from user account
//	@Tags			User
//	@Accept			json
//	@Success		200 {object}	response.Response{}		"Log out successful"
//	@Router			/logout [get]
func (u *UserHandler) LogoutUser(c *gin.Context) {
	c.SetCookie("user-auth", "", -1, "", "", false, true)
	response := response.SuccessResponse(http.StatusOK, "Log out successful", nil)
	c.JSON(http.StatusOK, response)
}

//	Add user addresss  godoc
//	@Summary		Add user addresss
//	@id				Add user addresss
//	@Description	Add the addresss of user
//	@Tags			User
//	@Param			input	body	request.Address	true	"inputs"
//	@Accept			json
//	@Failure		400	{object}	response.Response{}		"Missing or Invalid entry"
//  @Failure 		500 {object}	response.Response{} 	"Something went wrong"
//	@Success		200 {object}	response.Response{}		"Address saved successfully"
//	@Router			/profile/add-address [post]
func (u *UserHandler) AddAddress(c *gin.Context) {
	var body request.Address
	userId := utils.GetUserIdFromContext(c)
	body.UserID = userId
	if err := c.ShouldBindJSON(&body); err != nil {
		response := response.ErrorResponse(400, "Missing or invalid entry", err.Error(), body)
		c.JSON(400, response)
		return
	}
	if err := u.userService.Addaddress(c, body); err != nil {
		response := response.ErrorResponse(500, "Something went wrong!", err.Error(), body)
		c.JSON(500, response)
		return
	}
	response := response.SuccessResponse(200, "Address saved successfully", body)
	c.JSON(200, response)

}

//	Update user addresss  godoc
//	@Summary		update addresss
//	@id				update user addresss
//	@Description	update the addresss of user
//	@Tags			User
//	@Param			input	body	request.AddressPatchReq	true	"inputs"
//	@Accept			json
//	@Failure		400	{object}	response.Response{}		"Missing or Invalid entry"
//  @Failure 		500 {object}	response.Response{} 	"Something went wrong"
//	@Success		200 {object}	response.Response{}		"Address updated successfully"
//	@Router			/profile/edit-address [put]
func (u *UserHandler) UpdateAddress(c *gin.Context) {
	userId := utils.GetUserIdFromContext(c)
	var body request.AddressPatchReq
	body.UserID = userId
	if err := c.ShouldBindJSON(&body); err != nil {
		response := response.ErrorResponse(400, "Missing or invalid entry", err.Error(), body)
		c.JSON(400, response)
		return
	}
	if err := u.userService.UpdateAddress(c, body); err != nil {
		response := response.ErrorResponse(500, "Something went wrong!", err.Error(), body)
		c.JSON(500, response)
		return
	}
	response := response.SuccessResponse(200, "Address updated successfuly", nil)
	c.JSON(200, response)
}

//	Delete user addresss  godoc
//	@Summary		Delete user addresss
//	@id				Delete user addresss
//	@Description	Delete the addresss of user
//	@Tags			User
//	@Param			input	body	request.Address	true	"inputs"
//	@Accept			json
//	@Failure		400	{object}	response.Response{}		"Missing or Invalid entry"
//  @Failure 		500 {object}	response.Response{} 	"Something went wrong"
//	@Success		200 {object}	response.Response{}		"Address deleted successfully"
//	@Router			/profile/delete-address [delete]
func (u *UserHandler) DeleteAddress(c *gin.Context) {
	userId := utils.GetUserIdFromContext(c)
	addressId, err := utils.StringToUint(c.Param("id"))
	if err != nil {
		response := response.ErrorResponse(400, "Missing or invalid entry", err.Error(), nil)
		c.JSON(400, response)
		return
	}
	if err := u.userService.DeleteAddress(c, userId, addressId); err != nil {
		response := response.ErrorResponse(500, "Something went wrong!", err.Error(), nil)
		c.JSON(500, response)
		return
	}
	response := response.SuccessResponse(200, "Address deleted successfuly", nil)
	c.JSON(200, response)
}

//	Get all user addresss  godoc
//	@Summary		Get all user addresss
//	@id				Get all user addresss
//	@Description	Get all the addresss of user
//	@Tags			User
//	@Param			input	body	request.Address	true	"inputs"
//	@Accept			json
//	@Failure		500	{object}	response.Response{}		"User not detected"
//  @Failure 		500 {object}	response.Response{} 	"Something went wrong"
//	@Success		200 {object}	response.Response{}		"Get all addresses successfully"
//	@Router			/profile/getaddress [post]
func (u *UserHandler) GetAllAddress(c *gin.Context) {
	userId := utils.GetUserIdFromContext(c)
	if userId == 0 {
		response := response.ErrorResponse(500, "No user detected!", "", nil)
		c.IndentedJSON(400, response)
		return
	}
	address, err := u.userService.GetAllAddress(c, userId)
	if err != nil {
		response := response.ErrorResponse(500, "Something went wrong!", err.Error(), nil)
		c.IndentedJSON(500, response)
		return
	}
	response := response.SuccessResponse(200, "Get all address successful", address)
	c.IndentedJSON(200, response)
}

// Profile godoc
// @Summary Get user profile
// @Description Retrieve user profile details from the database
// @Tags User
// @Param Authorization header string true "token"
// @Success 200 {object} response.Response{} "Successfuly got profile"
// @Failure 500 {object} response.Response{} "Something went wrong!"
// @Router /profile [get]
func (u *UserHandler) Profile(c *gin.Context) {
	userId := utils.GetUserIdFromContext(c)

	user, err := u.userService.Profile(c, userId)
	if err != nil {
		response := response.ErrorResponse(500, "Something went wrong!", err.Error(), nil)
		c.JSON(500, response)
		return
	}
	response := response.SuccessResponse(200, "Successfuly got profile", user)
	c.JSON(200, response)

}

// AddToCart godoc
// @Summary Add product to cart
// @Description Add a product item to the user's cart
// @Tags Cart
// @Param Authorization header string true "Bearer token"
// @Param body body request.AddToCartReq true "Product details to be added to cart"
// @Success 200 {object} response.Response{} "Successfuly added product item to cart"
// @Failure 400 {object} response.Response{} "Invalid input or failed to add product item to cart"
// @Router /cart/add [post]
func (u *UserHandler) AddToCart(c *gin.Context) {
	var body request.AddToCartReq

	if err := c.ShouldBindJSON(&body); err != nil {
		response := response.ErrorResponse(400, "invalid input", err.Error(), body.ProductID)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	// get userId from context
	body.UserID = utils.GetUserIdFromContext(c)
	if body.UserID == 0 {
		c.JSON(400, "No user id on context")
		return
	}
	if err := u.userService.SaveCartItem(c, body); err != nil {
		response := response.ErrorResponse(http.StatusBadRequest, "Failed to add product item in cart", err.Error(), nil)
		c.JSON(400, response)
		return
	}
	response := response.SuccessResponse(200, "Successfuly added product item to cart ", body)
	c.JSON(200, response)

}

// GetcartItems godoc
// @Summary Get user's cart items
// @Description Retrieve cart items of the user from the database
// @Tags User
// @Param Authorization header string true "token"
// @Param count query int true "Number of items to retrieve"
// @Param page_number query int true "Page number for pagination"
// @Success 200 {object} response.Response{} "Get Cart Items successful"
// @Failure 400 {object} response.Response{} "Missing or invalid inputs"
// @Failure 500 {object} response.Response{} "Something went wrong!"
// @Router /cart/get [get]
func (u *UserHandler) GetcartItems(c *gin.Context) {
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
	cartItems, err := u.userService.GetCartItemsbyCartId(c, page, userId)
	if err != nil {
		response := response.ErrorResponse(500, "Something went wrong!", err.Error(), nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	response := response.SuccessResponse(200, "Get Cart Items successful", cartItems)
	c.JSON(http.StatusOK, response)
}

// UpdateCart godoc
// @Summary Update user's cart
// @Description Update cart items of the user in the database
// @Tags User
// @Param Authorization header string true "token"
// @Param input body request.UpdateCartReq true "Cart update details"
// @Success 200 {object} response.Response{} "Successfuly updated cart"
// @Failure 400 {object} response.Response{} "invalid input"
// @Failure 400 {object} response.Response{} "No user id on context"
// @Failure 500 {object} response.Response{} "Something went wrong!"
// @Router /cart/update [put]
func (u *UserHandler) UpdateCart(c *gin.Context) {
	var body request.UpdateCartReq

	if err := c.ShouldBindJSON(&body); err != nil {
		response := response.ErrorResponse(400, "invalid input", err.Error(), body)
		c.JSON(400, response)
		return
	}
	// get userId from context
	body.UserID = utils.GetUserIdFromContext(c)
	if body.UserID == 0 {
		response := response.ErrorResponse(400, "No user id on context", "", nil)
		c.JSON(400, response)
		return
	}
	if err := u.userService.UpdateCart(c, body); err != nil {
		response := response.ErrorResponse(500, "Something went wrong!", err.Error(), body)
		c.JSON(500, response)
		return
	}
	response := response.SuccessResponse(200, "Successfuly updated cart", body)
	c.JSON(200, response)

}

func (u *UserHandler) DeleteCartItem(c *gin.Context) {
	var body request.DeleteCartItemReq
	if err := c.ShouldBindJSON(&body); err != nil {
		response := response.ErrorResponse(400, "invalid input", err.Error(), body)
		c.JSON(400, response)
		return
	}
	// get userId from context
	body.UserID = utils.GetUserIdFromContext(c)
	if err := u.userService.RemoveCartItem(c, body); err != nil {
		response := response.ErrorResponse(500, "Something went wrong!", err.Error(), body)
		c.JSON(500, response)
		return
	}
	response := response.SuccessResponse(200, "Successfuly removed item from cart", body)
	c.JSON(200, response)
}

func (u *UserHandler) AddToWishList(c *gin.Context) {
	var body request.AddToWishReq
	if err := c.ShouldBindJSON(&body); err != nil {
		response := response.ErrorResponse(400, "invalid input", err.Error(), body.ProductID)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	// get userId from context
	body.UserID = utils.GetUserIdFromContext(c)
	if body.UserID == 0 {
		c.JSON(400, "No user id on context")
		return
	}
	if err := u.userService.SaveWishListItem(c, body); err != nil {
		response := response.ErrorResponse(http.StatusBadRequest, "Failed to add product item in wish list", err.Error(), nil)
		c.JSON(400, response)
		return
	}
	response := response.SuccessResponse(200, "Successfuly added product item to wish list ", body)
	c.JSON(200, response)

}

func (u *UserHandler) GetWishListItems(c *gin.Context) {
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
	wishItems, err := u.userService.GetWishListItemsbyCartId(c, page, userId)
	if err != nil {
		response := response.ErrorResponse(500, "Something went wrong!", err.Error(), nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	response := response.SuccessResponse(200, "Get Wish List Items successful", wishItems)
	c.JSON(http.StatusOK, response)
}

func (u *UserHandler) DeleteWishListItem(c *gin.Context) {
	var body request.DeleteWishItemReq
	if err := c.ShouldBindJSON(&body); err != nil {
		response := response.ErrorResponse(400, "invalid input", err.Error(), body)
		c.JSON(400, response)
		return
	}
	// get userId from context
	body.UserID = utils.GetUserIdFromContext(c)
	if err := u.userService.RemoveWishListItem(c, body); err != nil {
		response := response.ErrorResponse(500, "Something went wrong!", err.Error(), body)
		c.JSON(500, response)
		return
	}
	response := response.SuccessResponse(200, "Successfuly removed item from wish list", body)
	c.JSON(200, response)
}

func (u *UserHandler) GetWalletHistory(c *gin.Context) {
	userId := utils.GetUserIdFromContext(c)
	history, err := u.userService.GetWalletHistory(c, userId)
	if err != nil {
		response := response.ErrorResponse(500, "Something went wrong!", err.Error(), nil)
		c.IndentedJSON(500, response)
		return
	}
	response := response.SuccessResponse(200, "Get wallet history successful", history)
	c.IndentedJSON(200, response)
}
