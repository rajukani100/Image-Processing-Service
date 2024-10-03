package services

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var secret_key = []byte("IAMBEASTWOW")

func GenerateJwt(username *string) (string, error) {
	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": *username,
		"iat":      time.Now().Unix(),
		"exp":      time.Now().Add(24 * time.Hour).Unix(),
	})
	tokenString, err := claims.SignedString(secret_key)
	if err != nil {
		return "", err // signing error
	}
	return tokenString, nil //success

}

func VerifyToken(tokenString string) error {
	token, err := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
		return secret_key, nil
	})
	if err != nil {
		return err
	}
	if !token.Valid {
		return errors.New("invalid token")
	}
	return nil
}
