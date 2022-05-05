package jwt

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/google/uuid"
)

var jwtKey = []byte("4AYDipIpozCCL51ETa7RksNTh/t6yvnQTKGW9ECkF/Loje8xzJH3xyQxgovNB4t3oS5iH+iDm5mlygssd6Feaw==")

func MakeToken() (string, error) {
	c := jwt.StandardClaims{
		Issuer:    "app",
		Subject:   "user",
		Audience:  "all",
		IssuedAt:  time.Now().Unix(),
		ExpiresAt: time.Now().Add(24 * 7 * time.Hour).Unix(),
		Id:        uuid.New().String(),
	}
	tokenObj := jwt.NewWithClaims(jwt.SigningMethodHS256, c)

	tokenStr, err := tokenObj.SignedString(jwtKey)
	return tokenStr, err
}

func ParseToken(tokenString string) (*jwt.StandardClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &jwt.StandardClaims{}, func(token *jwt.Token) (i interface{}, err error) {
		return jwtKey, nil
	})
	if err != nil {
		return nil, err
	}
	if claims, ok := token.Claims.(*jwt.StandardClaims); ok && token.Valid {
		return claims, nil
	}
	return nil, errors.New("invalid token")
}
