package handler

import (
	"gestao/internal/middleware"
	"gestao/internal/model"
	serve "gestao/internal/servece"
	"net/http"

	"github.com/gin-gonic/gin"
)

func PostClienteHandler(service *serve.ClienteService) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		cliente := model.ClienteDTO{}
		userID, ok := ctx.Get(middleware.USER_ID)

		if !ok || userID == "" {
			ctx.JSON(http.StatusInternalServerError, gin.H{"estado": "erro", "mensagem": "id inválido"})
			return
		}

		if err := ctx.ShouldBindJSON(&cliente); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"estado": "erro", "mensagem": err.Error()})
			return
		}

		if err := service.CadastrarCliente(model.Cliente{
			Nome:     cliente.Nome,
			Contacto: cliente.Contacto,
			UserID:   userID.(string),
		}); err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"estado": "erro", "mensagem": err.Error()})
			return
		}

		ctx.JSON(http.StatusCreated, gin.H{"estado": "sucesso", "mensagem": "cadastro feio"})
	}
}
