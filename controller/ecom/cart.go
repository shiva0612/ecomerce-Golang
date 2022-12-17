package ecom

import (
	"net/http"
	"shiva/database"
	"strconv"

	"github.com/gin-gonic/gin"
)

/*
productID and quantity will be provided
userID is fetched from context
calculate sum of product
added to user.cart in collection
*/
func AddProductToCart(c *gin.Context) {
	userID := c.GetString("user_id")
	productID, _ := c.GetQuery("product_id")
	quantityS, _ := c.GetQuery("quantity")
	if productID == "" || quantityS == "" {
		c.String(http.StatusBadRequest, "product_id or quantity is empty")
		return
	}
	quantity, err := strconv.Atoi(quantityS)
	if err != nil {
		c.String(http.StatusBadRequest, "quantity is not a number")
		return
	}
	if quantity < 0 {
		c.String(http.StatusBadRequest, "quantity must be positive")
		return
	}

	err = database.AddProductToCart(userID, productID, quantity)
	if err != nil {
		c.String(http.StatusInternalServerError, "technical error:try later")
		return
	}
	c.String(http.StatusOK, "added to cart")

}

/*
productID and quantity will be provided
userID is fetched from token
removed from user.cart in collection
*/
func RemoveProductFromCart(c *gin.Context) {
	userID := c.GetString("user_id")
	productID, _ := c.GetQuery("product_id")
	if productID == "" {
		c.String(http.StatusBadRequest, "product_id  is empty")
		return
	}

	err := database.RemoveProductFromCart(userID, productID)
	if err != nil {
		c.String(http.StatusInternalServerError, "technical error:try later")
		return
	}
	c.String(http.StatusOK, "item removed from cart")

}

/*
userID from token will be fetched
user.cart will be provided
*/
func ViewCart(c *gin.Context) {
	userID := c.GetString("user_id")
	cartItems, err := database.ViewCart(userID)
	if err != nil {
		c.String(http.StatusInternalServerError, "technical error:try later")
		return
	}

	c.JSON(http.StatusOK, cartItems)

}

/*
checkout cart = get addresslist, pymnt option,
choose email and phone to be used
and click on place order
*/
func CheckoutCart(c *gin.Context) {

	userID := c.GetString("user_id")
	address, err := database.CheckoutCart(userID)
	if err != nil {
		c.String(http.StatusInternalServerError, "technical error:try later")
		return
	}

	c.JSON(http.StatusOK, address)

}
