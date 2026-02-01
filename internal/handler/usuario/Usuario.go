package handler

import (
	"gestao/internal/model"
	serve "gestao/internal/servece"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Usuario(serve *serve.UsuarioService) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var usuario model.Usuario
		if err := ctx.ShouldBindJSON(&usuario); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"estado": "erro", "mensagem": err.Error()})
			return
		}

		if err := serve.CriarConta(usuario); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"estado": "erro", "mensagem": err.Error()})
			return
		}

		ctx.JSON(http.StatusOK, gin.H{"estado": "sucesso", "mensagem": "usuario ciado"})

	}
}

func Login(serve *serve.UsuarioService) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var usuario model.Usuario
		if err := ctx.ShouldBindJSON(&usuario); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"estado": "erro", "mensagem": err.Error()})
			return
		}

		token, err := serve.Logar(usuario)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"estado": "erro", "mensagem": err.Error()})
			return
		}

		ctx.JSON(http.StatusOK, gin.H{"estado": "sucesso", "mensagem": "guarda bem o token", "token": token})
	}
}
