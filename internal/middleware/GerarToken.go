package middleware

import (
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var apiKey = []byte(os.Getenv("KEY_TOKEN"))

type claims struct {
	UserId string `json:"user_id"`
	jwt.RegisteredClaims
}

func GerarToken(userId string) (string, error) {
	userClaims := claims{
		UserId: userId,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
			Issuer:    "gestao-clientes-api",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, userClaims)

	return token.SignedString(apiKey)
}
