package service

import (
	"fiber-rest-api/model/domain"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

type TokenMetaData struct {
	Expires int64 `json:"expires"`
}

func GenerateNewAccessToken(secretKey, expiresIn string, payload *domain.User) (string, error) {
	var (
		timeDur        time.Duration
		expirationTime int64
	)

	expire, _ := strconv.Atoi(expiresIn)
	timeDur = time.Duration(expire)
	expirationTime = time.Now().Add(time.Minute * timeDur).Unix()

	Jwt := jwt.New(jwt.SigningMethodHS512)

	//set Claims
	claims := Jwt.Claims.(jwt.MapClaims)
	claims["id"] = payload.Id
	claims["name"] = payload.Username
	claims["email"] = payload.Email
	claims["exp"] = expirationTime

	tokenString, err := Jwt.SignedString([]byte(secretKey))
	if err != nil {
		panic(err)
	}

	return tokenString, nil
}
