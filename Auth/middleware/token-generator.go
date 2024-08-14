package middleware

import (
	"time"

	"github.com/golang-jwt/jwt"
)

func GenerateToken(userID string) (string, error) {
	claims := jwt.MapClaims{}
	claims["user_id"] = userID
	claims["exp"] = time.Now().Add(time.Hour * 24).Unix() // Token 24 saat geçerli
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(secretKey)
}
