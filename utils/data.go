package utils

import (
	"fmt"
	"log"

	"github.com/dgrijalva/jwt-go"
)

func ParseDataEmail(bearerToken string) (email string, err error) {
	token, err := jwt.Parse(bearerToken, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte("kalorize"), nil
	})

	if err != nil {
		log.Printf("Error parsing JWT token: %v", err)
		return "", err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		emailClaim := claims["Email"]
		if emailClaim == nil {
			err := fmt.Errorf("email claim is missing in JWT token")
			log.Printf("Error: %v", err)
			return "", err
		}
		email = emailClaim.(string)
	}
	return email, err
}

func ParseDataFullname(bearerToken string) (email string, err error) {
	token, err := jwt.Parse(bearerToken, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte("kalorize"), nil
	})

	if err != nil {
		log.Printf("Error parsing JWT token: %v", err)
		return "", err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		emailClaim := claims["Fullname"]
		if emailClaim == nil {
			err := fmt.Errorf("email claim is missing in JWT token")
			log.Printf("Error: %v", err)
			return "", err
		}
		email = emailClaim.(string)
	}
	return email, err
}
