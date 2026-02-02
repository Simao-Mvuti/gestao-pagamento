package middleware

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		authHeader := ctx.GetHeader("Authorization")

		if authHeader == "" {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"estado": "erro", "mensagem": "token inexistente"})
			return
		}

		partes := strings.Split(authHeader, " ")

		if len(partes) != 2 || strings.ToUpper(partes[0]) != "BEARER" {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"estado": "erro", "mensagem": "formato de token inválido"})
			return
		}

		tokenString := partes[1]

		token, err := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {

			if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, jwt.ErrSignatureInvalid
			}
			return apiKey, nil
		})

		if err != nil || !token.Valid {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"estado": "erro", "mensagem": "token inválido ou expirado"})
			return
		}

		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			userID := claims["user_id"]
			ctx.Set("user_id", userID)
		} else {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"estado": "erro", "mensagem": "falha ao processar claims"})
			return
		}

		ctx.Next()
	}
}
