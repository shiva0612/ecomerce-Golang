package helper

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

// binds data to user from (form / json )
func Bind(c *gin.Context, obj any) {
	funcName := "bind"
	switch c.ContentType() {
	case "application/json":
		err := c.ShouldBindJSON(obj)
		if err != nil {
			log.Printf("[%s]: Error while binding json(BadRequest): %s", funcName, err.Error())
			usermsg := fmt.Sprintf("invalid input: err = %s", err.Error())
			c.String(http.StatusBadRequest, usermsg)
			return
		}
	case "application/multipart/form-data":
		err := c.ShouldBind(obj)
		if err != nil {
			log.Printf("[%s]: Error while binding form(BadRequest): %s", funcName, err.Error())
			usermsg := fmt.Sprintf("invalid input: err = %s", err.Error())
			c.String(http.StatusBadRequest, usermsg)
			return
		}

	}
}
