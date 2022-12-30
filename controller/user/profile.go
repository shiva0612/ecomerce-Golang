package user

import (
	"log"
	"net/http"
	"shiva/database"
	"shiva/model"

	"github.com/gin-gonic/gin"
)

func Profile(c *gin.Context) {
	userID := c.GetString("user_id")
	user, err := database.GetProfile(userID)
	if err != nil {
		c.String(http.StatusInternalServerError, "technical error:try later")
		return
	}
	c.JSON(http.StatusOK, user)
}

func AddAddress(c *gin.Context) {

	address := make([]model.Address, 0)
	err := c.ShouldBindJSON(&address)
	if err != nil {
		c.String(http.StatusBadRequest, "please give proper json")
		return
	}

	userID := c.GetString("user_id")
	err = database.AddAddress(userID, address)
	if err != nil {
		log.Println("error wgile adding address: ", err.Error())
		c.String(http.StatusInternalServerError, "technical error, try later")
		return
	}

	c.String(http.StatusOK, "address added")
}

func GetAddress(c *gin.Context) {

	userID := c.GetString("user_id")

	add, err := database.GetAddress(userID)
	if err != nil {
		c.String(http.StatusInternalServerError, "technical error, try later")
		return
	}
	c.JSON(http.StatusOK, add)
}

/*
userID from token is fetched
users.orders and refer order again and give all order details
*/
func ViewOrders(c *gin.Context) {
	userID := c.GetString("user_id")

	orders, err := database.ViewOrders(userID)
	if err != nil {
		c.String(http.StatusInternalServerError, "technical error:try later")
		return
	}

	c.JSON(http.StatusOK, orders)
}
