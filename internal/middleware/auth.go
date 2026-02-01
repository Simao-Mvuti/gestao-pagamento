package middleware

import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		apiKey := ctx.GetHeader("X-API-KEY")
		chave := os.Getenv("KEY-API")

		if apiKey == "" {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"estado": "erro",
				"mensagem": "chave de acesso não fornecida"})
			return
		}

		if apiKey != chave {
			ctx.AbortWithStatusJSON(http.StatusForbidden, gin.H{
				"estado":   "erro",
				"mensagem": "chave de acesso inválida",
			})
			return
		}

		ctx.Next()
	}
}
