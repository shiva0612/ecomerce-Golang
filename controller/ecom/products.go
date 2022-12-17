package ecom

import (
	"log"
	"net/http"
	"shiva/database"
	"shiva/model"

	"github.com/gin-gonic/gin"
)

/*
only admin can add products check usertype
from the token
*/
func AddProducts(c *gin.Context) {

	products := make([]model.Product, 0)
	err := c.ShouldBindJSON(&products)
	if err != nil {
		log.Println("error adding product: while binding json: ", err.Error())
		c.String(http.StatusBadRequest, "give proper json")
		return
	}

	err = database.AddProducts(products)
	if err != nil {
		c.String(http.StatusInternalServerError, "technical error:try later")
		return
	}

	c.String(http.StatusOK, "products added")

}

/*
lh:port
lh:port/?category="books"
lh:port/?category="books"&search="theory"
lh:port/?search="theory"
*/
func ProductsList(c *gin.Context) {

	qp := map[string]string{}
	c.BindQuery(&qp)

	products, err := database.ProductsList(qp)
	if err != nil {
		c.String(http.StatusInternalServerError, "technical error:try later")
		return
	}

	c.JSON(http.StatusOK, products)
}
