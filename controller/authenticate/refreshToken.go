package authenticate

import (
	"github.com/gin-gonic/gin"
)

/*
if (token expired or just about to expire) + (refresh token not expired)
*/
func RefreshToken(c *gin.Context) {

	c.String(200, "ok")
}
