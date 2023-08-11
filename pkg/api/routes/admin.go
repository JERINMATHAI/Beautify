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
			brand.GET("/get", productHandler.GetAllBrands)
			brand.POST("/add", productHandler.AddCategory)
		}

		product := api.Group("/products")
		{
			product.GET("/list", productHandler.ListProducts)
			product.POST("/add", productHandler.AddProduct)
			product.PUT("/update", productHandler.UpdateProduct)
			product.DELETE("/delete", productHandler.DeleteProduct)
			product.POST("/addimage", productHandler.AddImage)

			product.POST("/product-item", productHandler.AddProductItem)
			product.GET("/product-item/:product_id", productHandler.GetProductItem)
		}

		coupons := api.Group("/coupons")
		{
			coupons.GET("/list", couponHandler.ListAllCoupons)
			coupons.POST("/create", couponHandler.CreateNewCoupon)
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
