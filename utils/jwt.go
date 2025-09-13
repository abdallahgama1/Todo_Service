// utils/jwt.go

package utils

import (
	"time"
	"github.com/golang-jwt/jwt/v5"
)

var jwtKey = []byte("my_super_secret_key")

type Claims struct {
	UserID uint `json:"userId"`
	jwt.RegisteredClaims
}

func GenerateToken(userId uint) (string, error) {
	expirationTime := time.Now().Add(24 * time.Hour)
	claims := &Claims{
		UserID: userId,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtKey)
}