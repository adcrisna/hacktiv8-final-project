package helpers

import (
	"errors"

	"github.com/dgrijalva/jwt-go"
)

const (
	SECRET_KEY = "chrisna-scret"
)

func GenerateToken(id int, email string) string {
	claims := jwt.MapClaims{
		"id":    id,
		"email": email,
	}
	parseToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, _ := parseToken.SignedString([]byte(SECRET_KEY))
	return signedToken
}

func VerifyToken(tokenString string) (interface{}, error) {
	errResponse := errors.New("Token-Invalid")
	token, _ := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errResponse
		}
		return []byte(SECRET_KEY), nil
	})
	r, ok := token.Claims.(jwt.MapClaims)
	if !ok && !token.Valid {
		return nil, errResponse
	}
	return r, nil
}
