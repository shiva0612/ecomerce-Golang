package authenticate

import (
	"fmt"
	"log"
	"net/http"
	"shiva/database"
	"shiva/helper"
	"shiva/model"

	"github.com/gin-gonic/gin"
)

// get the credentials from form/body
// bind to struct
// get the user from db with the username
// validate the psw
// generate tokens and store in header
func Login(c *gin.Context) {
	funcName := "Login"
	user_login := &model.User_Login{}

	//get user obj
	helper.Bind(c, user_login)

	// get the user from db with the username
	user, err := database.GetUser(user_login.Username)
	if err != nil {
		var respCode int
		var respMsg string

		switch err.Error() {
		case database.ErrUSER_NOT_FOUND:
			log.Printf("[%s]: user does not exist", funcName)
			respCode = http.StatusBadRequest
			respMsg = fmt.Sprintf(database.ErrUSER_NOT_FOUND)
		default:
			log.Printf("[%s]: error while getting user from db: %s", funcName, err.Error())
			respCode = http.StatusInternalServerError
			respMsg = fmt.Sprintf("Technical error: please try again later")
		}
		c.String(respCode, respMsg)
		return
	}

	// validate the user.psw with db.user.psw
	if !helper.VerifyPassword(user_login.Password, user.Password) {
		log.Printf("[%s]: invalid password", funcName)
		respMsg := "invalid password"
		c.String(http.StatusUnauthorized, respMsg)
	}

	// generate tokens and store in header
	Token, RefreshToken, err := helper.GenerateTokens(user)
	if err != nil {
		log.Printf("[%s]: error generating token: %s", funcName, err.Error())
		respMsg := fmt.Sprintf("technical error: please try later")

		c.String(http.StatusInternalServerError, respMsg)
		return
	}

	//setting tokens in headers
	c.Header(TOKEN, Token)
	c.Header(REFRESH_TOKEN, RefreshToken)

	c.String(http.StatusOK, "Login successful...")

}
