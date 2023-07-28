package routes

import (
	"beautify/pkg/api/handler"
	"beautify/pkg/api/middleware"

	"github.com/gin-gonic/gin"
)

func UserRoutes(api *gin.RouterGroup, userHandler *handler.UserHandler, productHandler *handler.ProductHandler, paymentHandler *handler.PaymentHandler, couponHandler *handler.CouponHandler, orderHandler *handler.OrderHandler) {

	signup := api.Group("/signup")
	{
		signup.POST("/", userHandler.UserSignup)

	}

	login := api.Group("/login")
	{
		login.POST("/", userHandler.LoginSubmit)
		login.POST("/otp-verify", userHandler.UserOTPVerify)
	}

	// Middleware
	api.Use(middleware.AuthenticateUser)
	{
		api.GET("/", userHandler.Home)
		api.GET("/logout", userHandler.LogoutUser)

		products := api.Group("/products")
		{
			products.GET("/brands", productHandler.GetAllBrands)
			products.GET("/", productHandler.ListProducts)
		}

		profile := api.Group("/profile")
		{
			profile.POST("/address", userHandler.AddAddress)
			profile.PUT("/address", userHandler.UpdateAddress)
			profile.DELETE("/address/:id", userHandler.DeleteAddress)
			profile.GET("/address", userHandler.GetAllAddress)
			profile.GET("/", userHandler.Profile)
		}

		cart := api.Group("/cart")
		{
			cart.POST("/", userHandler.AddToCart)
			cart.GET("/", userHandler.GetcartItems)
			cart.PUT("/", userHandler.UpdateCart)
			cart.DELETE("/", userHandler.DeleteCartItem)
		}

		wishlist := api.Group("/wishlist")
		{
			wishlist.POST("/add", userHandler.AddToWishList)
			wishlist.GET("/get", userHandler.GetWishListItems)
			wishlist.DELETE("/remove", userHandler.DeleteWishListItem)
		}

		coupon := api.Group("/coupons")
		{
			coupon.GET("/", couponHandler.ListAllCoupons)
		}

		order := api.Group("/order")
		{
			order.POST("/createOrder", orderHandler.CreateOrder)
			order.PUT("/updateOrder", orderHandler.UpdateOrder)
			order.GET("/listOrder", orderHandler.ListAllOrders)
			order.DELETE("/cancelOrder", orderHandler.CancelOrder)
			order.POST("/placeOrder", orderHandler.PlaceOrder)
			order.POST("/payment", orderHandler.CheckOut)
		}

		Return := api.Group("/return")
		{
			Return.POST("/", orderHandler.ReturnOrder)
		}
	}

}
