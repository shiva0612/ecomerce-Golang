package helper

import (
	"log"

	"golang.org/x/crypto/bcrypt"
)

// if err in hashing return "". so, check if empty and proceed
func HashPassword(psw string) string {
	bytes, err := bcrypt.GenerateFromPassword([]byte(psw), bcrypt.MinCost)
	if err != nil {
		log.Println("error while generating hash for password: ", err.Error())
		return ""
	}
	return string(bytes)
}

func VerifyPassword(psw, hashedpsw string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedpsw), []byte(psw))
	if err != nil {
		log.Println("password mismatch")
		return false
	}
	return true
}
