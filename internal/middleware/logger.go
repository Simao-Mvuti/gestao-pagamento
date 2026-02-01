package middleware

import (
	"log"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

func Logger() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// Antes do handler
		data := time.Now()
		metodo := strings.ToUpper(ctx.Request.Method)
		path := ctx.Request.URL.Path

		log.Printf("[%s] %s %s - Iniciando...", data.Format("2006-01-02 15:04:05"), metodo, path)

		// Chama o handler
		ctx.Next()

		// Depois do handler (opcional: status)
		status := ctx.Writer.Status()
		duracao := time.Since(data)
		log.Printf("[%s] %s %s | Status: %d | Duração: %s", data.Format("2006-01-02 15:04:05"), metodo, path, status, duracao)
	}
}
