package authenticate

import (
	"net/http"
	"shiva/database"

	"github.com/gin-gonic/gin"
)

// remove the tokens from headers
func Logout(c *gin.Context) {
	c.Header(TOKEN, "")
	c.Header(REFRESH_TOKEN, "")
	c.String(http.StatusOK, "you are logged out")
}

func Delete(c *gin.Context) {
	userID := c.GetString("user_id")

	err := database.DeleteUser(userID)
	if err != nil {
		c.String(http.StatusInternalServerError, "error deleting user")
		return
	}
	c.String(http.StatusOK, "deleted user")
}
