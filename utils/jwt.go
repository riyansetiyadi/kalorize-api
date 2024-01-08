package utils

import (
	"time"

	"github.com/dgrijalva/jwt-go"
)

func GenerateJWTToken(fullname, email, Role, key string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"Fullname": fullname,
		"Email":    email,
		"Role":     Role,
		"exp":      time.Now().Add(time.Hour * 1).Unix(),
	})

	tokenString, err := token.SignedString([]byte(key))
	if err != nil {
		return err.Error(), err
	}
	return tokenString, err
}
