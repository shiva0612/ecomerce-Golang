package routes

import (
	"shiva/controller/ecom"

	"github.com/gin-gonic/gin"
)

func EcomRoutes(router *gin.Engine) {
	router.POST("/admin/addproducts", ecom.AddProducts)

	router.GET("/", ecom.ProductsList)
	router.POST("/cart/add", ecom.AddProductToCart)
	router.POST("/cart/remove", ecom.RemoveProductFromCart)
	router.POST("/cart/view", ecom.ViewCart)
	router.GET("/cart/checkout", ecom.CheckoutCart)
	router.POST("/cart/placeorder", ecom.PlaceOrder)
	router.GET("/orders", ecom.ViewOrders)
}
