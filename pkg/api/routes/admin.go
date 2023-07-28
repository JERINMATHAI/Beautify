package routes

import (
	"beautify/pkg/api/handler"
	"beautify/pkg/api/middleware"

	"github.com/gin-gonic/gin"
)

func AdminRoutes(api *gin.RouterGroup, adminHandler *handler.AdminHandler, productHandler *handler.ProductHandler, couponHandler *handler.CouponHandler, orderHandler *handler.OrderHandler, paymentHandler *handler.PaymentHandler) {

	signup := api.Group("/signup")
	{
		signup.POST("/", adminHandler.AdminSignUp)
	}

	login := api.Group("/login")
	{
		login.POST("/", adminHandler.AdminLogin)
	}

	// Middleware
	api.Use(middleware.AuthenticateAdmin)
	{
		api.GET("/", adminHandler.AdminHome)
		api.GET("/logout", adminHandler.LogoutAdmin)

		user := api.Group("/users")
		{
			user.GET("/", adminHandler.ListUsers)
			user.PATCH("/block", adminHandler.BlockUser)
		}

		brand := api.Group("/brands")
		{
			brand.GET("/", productHandler.GetAllBrands)
			brand.POST("/", productHandler.AddCategory)
		}

		product := api.Group("/products")
		{
			product.GET("/", productHandler.ListProducts)
			product.POST("/", productHandler.AddProduct)
			product.PUT("/", productHandler.UpdateProduct)
			product.DELETE("/", productHandler.DeleteProduct)
			product.POST("/product-item", productHandler.AddProductItem)
			product.GET("/product-item/:product_id", productHandler.GetProductItem)
		}

		coupons := api.Group("/coupons")
		{
			coupons.GET("/", couponHandler.ListAllCoupons)
			coupons.POST("/", couponHandler.CreateNewCoupon)
		}
		paymentmethod := api.Group("/paymentmethod")
		{
			paymentmethod.POST("/add", paymentHandler.AddpaymentMethod)
			paymentmethod.GET("/view", paymentHandler.GetPaymentMethods)
			paymentmethod.PUT("/update", paymentHandler.UpdatePaymentMethod)
			paymentmethod.DELETE("/delete", paymentHandler.DeleteMethod)
		}
		order := api.Group("/order")
		{
			order.GET("/listOrder", orderHandler.ListAllOrders)
		}
		dashboard := api.Group("/dashboard")
		{
			dashboard.GET("/salesReport", orderHandler.SalesReport)
		}
	}

}
