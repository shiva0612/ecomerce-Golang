package authenticate

import (
	"fmt"
	"log"
	"net/http"
	"shiva/model"

	"shiva/helper"

	"shiva/database"

	"github.com/gin-gonic/gin"
)

// get user struct from (form/json) using binding
// check if email/username/phone is already registered
// hash password and store hashedpsw in user struct
// creating user in db
func Signup(c *gin.Context) {

	funcName := "Signup"
	user_signup := &model.User_Signup{}

	// get user struct from (form/json) using binding
	helper.Bind(c, user_signup)

	// check if email/username/phone is already registered
	//TODO: replace the error with proper mongo error
	b, err := database.CheckIfUserExists(user_signup.Username, user_signup.Email, user_signup.Phone)
	if b {
		if err.Error() == database.ErrUser_ALREADY_EXIST {
			log.Printf("[%s]: %s", funcName, err.Error())
			respMsg := fmt.Sprintf("user with username,email and phone is already registered")
			c.String(http.StatusBadRequest, respMsg)
			return
		}
		log.Printf("[%s]: %s", funcName, err.Error())
		respMsg := fmt.Sprintf("technical error: please try again later")
		c.String(http.StatusInternalServerError, respMsg)
		return
	}

	// hash password and store hashedpsw in user struct
	hashedpsw := helper.HashPassword(user_signup.Password)
	if hashedpsw == "" {
		log.Printf("[%s]: error while hashing password", funcName)
		respMsg := "please try again later"
		c.String(http.StatusInternalServerError, respMsg)
		return
	}
	user_signup.Password = hashedpsw

	//creating user in db
	err = database.CreateUser(user_signup)
	if err != nil {
		log.Printf("[%s]: error creating user: %s", funcName, err.Error())
		respMsg := fmt.Sprintf("technical error: please try later")

		c.String(http.StatusInternalServerError, respMsg)
	}

	log.Printf("[%s] user created: %s", funcName, user_signup.Username)
	respMsg := fmt.Sprintf("user created: %s, please go to login page ...", user_signup.Username)
	c.String(http.StatusOK, respMsg)

}
