package services

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var secret_key = "IAMBEASTWOW"

func GenerateJwt(username string) (string, error) {
	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": username,
		"iat":      time.Now().Unix(),
		"exp":      time.Now().Add(24 * time.Hour).Unix(),
	})
	tokenString, err := claims.SignedString(secret_key)
	return tokenString, err

}
