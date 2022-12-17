package middleware

import (
	"net/http"
	"shiva/controller/authenticate"
	"shiva/helper"
	"shiva/model"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

const (
	errUnauthorized = "unauthorized access"
)

// check if tokens are present in headers, else abort
// check if tokens are valid and get the claims from token
func Authenticate(c *gin.Context) {

	Token := c.GetHeader(authenticate.TOKEN)

	userClaims := &model.UserClaims{}
	err := helper.GetClaimsFromToken(Token, userClaims)
	if err != nil {
		c.String(http.StatusUnauthorized, errUnauthorized)
		c.Abort()
	}

	if c.Request.URL.Path == "/refresh-token" {
		RefreshToken := c.GetHeader(authenticate.REFRESH_TOKEN)
		refreshClaims := &jwt.StandardClaims{}
		err := helper.GetClaimsFromToken(RefreshToken, refreshClaims)
		if err != nil {
			c.String(http.StatusUnauthorized, errUnauthorized)
			c.Abort()
		}
		resfreshToken(userClaims, refreshClaims)
	}

	c.Set("user_id", userClaims.UserID)
	c.Next()

}

/*
if (token expired or just about to expire) + (refresh token not expired)
*/
func resfreshToken(userClaims *model.UserClaims, refreshClaims *jwt.StandardClaims) {

}
