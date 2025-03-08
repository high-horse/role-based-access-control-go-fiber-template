package utils

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt"
)

var jwtKey = []byte("super_secret_key")

type Claims struct {
	Username string `json:"username"`
	Roles []string `json:"roles"`
	Permissions []string `json:"permissions"`
	jwt.StandardClaims
}

func GenerateToken(username string, roles []string, permissions []string) (string, error) {
	expirationTime := time.Now().Add(24 * time.Hour)
	claims := &Claims{
		Username: username,
		Roles: roles,
		Permissions: permissions,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtKey)
}

func ValidateToken(tokenString string) (*Claims, error) {
	claims := &Claims{}

	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})
	
	if err != nil {
		return nil, err
	}
	
	if !token.Valid {
		return nil, errors.New("invalid token")
	}
	return claims, nil
}