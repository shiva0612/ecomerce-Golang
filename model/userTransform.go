package model

import (
	"github.com/jinzhu/copier"
)

func User_from_userSingup(user_signup *User_Signup) *User {
	user := &User{}
	copier.Copy(user, user_signup)
	return user
}

func GetUserClaims(user *User) *UserClaims {
	userClaims := &UserClaims{}
	copier.Copy(userClaims, user)

	return userClaims
}
