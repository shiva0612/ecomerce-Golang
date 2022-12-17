package helper

import (
	"log"
	"shiva/config"
	"shiva/model"
	"time"

	"github.com/dgrijalva/jwt-go"
)

func GenerateTokens(user *model.User) (string, string, error) {
	funcName := "GenerateTokens"
	var err error
	TOKEN_TIME := config.Prjconfig.Jwt.TokenTime * time.Minute
	REFRESH_TIME := config.Prjconfig.Jwt.RefreshTime * time.Minute

	tokenTime := time.Now().Add(TOKEN_TIME).Unix()
	refreshTime := time.Now().Add(REFRESH_TIME).Unix()

	userClaims := model.GetUserClaims(user)
	userClaims.StandardClaims = jwt.StandardClaims{
		ExpiresAt: tokenTime,
	}

	refreshClaims := jwt.StandardClaims{
		ExpiresAt: refreshTime,
	}

	Token, _ := jwt.NewWithClaims(jwt.SigningMethodHS256, userClaims).SignedString([]byte(config.Prjconfig.Jwt.TokenKey))
	RefreshToken, err := jwt.NewWithClaims(jwt.SigningMethodHS256, refreshClaims).SignedString([]byte(config.Prjconfig.Jwt.TokenKey))
	if err != nil {
		log.Printf("[%s]: err signing token: %s", funcName, err.Error())
		return "", "", err
	}

	return Token, RefreshToken, nil

}

func GetClaimsFromToken(token string, claims jwt.Claims) error {
	funcName := "getClaimsFromToken"
	token_parsed, err := jwt.ParseWithClaims(token, claims, func(t *jwt.Token) (interface{}, error) {
		return []byte(config.Prjconfig.Jwt.TokenKey), nil
	})

	if err != nil || !token_parsed.Valid {
		log.Printf("[%s]: unauthorized : %s", funcName, err.Error())
		return err
	}
	return nil
}
