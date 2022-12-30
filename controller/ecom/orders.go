package ecom

import (
	"net/http"
	"shiva/database"
	"shiva/model"

	"github.com/gin-gonic/gin"
)

/*
selected (address, pymt option, email, phone) as input
click on place order
orderID as response
*/
func PlaceOrder(c *gin.Context) {

	userID := c.GetString("user_id")
	checkoutCart := new(model.CheckoutCart)
	err := c.ShouldBindJSON(&checkoutCart)
	if err != nil {
		c.String(http.StatusBadRequest, "give proper json")
		return
	}
	orderID, err := database.PlaceOrder(userID, *checkoutCart)
	if err != nil {
		c.String(http.StatusInternalServerError, "technical error:try later")
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"order_id": orderID,
	})

}
