package handler

import (
	"gestao/internal/model"
	serve "gestao/internal/servece"
	"net/http"

	"github.com/gin-gonic/gin"
)

func PostClienteHandler(service serve.Repository) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		cliente := model.Cliente{}

		if err := ctx.ShouldBindJSON(&cliente); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"estado": "erro", "mensagem": err.Error()})
			return
		}

		if err := service.CadastrarCliente(cliente); err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"estado": "erro", "mensagem": err.Error()})
			return
		}

		ctx.JSON(http.StatusCreated, gin.H{"estado": "sucesso", "mensagem": "cadastro feio"})
	}
}
