package routes

import (
	"shiva/controller/authenticate"
	"shiva/controller/user"

	"github.com/gin-gonic/gin"
)

func UserRoutes(router *gin.Engine) {

	router.POST("/refresh-token", authenticate.RefreshToken)
	router.GET("/logout", authenticate.Logout)
	router.POST("/deleteme", authenticate.Delete)

	router.GET("/profile", user.Profile)
	router.GET("/orders", user.ViewOrders)
	router.POST("/address/add", user.AddAddress)
	router.GET("address/view", user.GetAddress)

}
