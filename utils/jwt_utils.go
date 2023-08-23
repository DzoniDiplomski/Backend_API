package utils

import (
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
)

func GenerateJWT(id int, role string) string {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":   id,
		"role": role,
		"exp":  time.Now().Add(time.Hour * 24).Unix(),
	})

	tokenString, _ := token.SignedString([]byte(os.Getenv("JWT_SECRET_KEY")))
	return tokenString
}

func CheckJWTValidity(jwtString string) bool {
	token, err := jwt.Parse(jwtString, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("JWT_SECRET_KEY")), nil
	})

	if err != nil {
		return false
	}

	if !token.Valid {
		return false
	}

	return true
}
