package handler

import (
	"gestao/internal/middleware"
	"gestao/internal/model"
	serve "gestao/internal/servece"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func PutClienteHandler(serve *serve.ClienteService) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id, ok := ctx.Params.Get("id")
		userID, ok := ctx.Get(middleware.USER_ID)

		clienteId, err := strconv.Atoi(id)

		if !ok || err != nil || clienteId < 0 {
			ctx.JSON(http.StatusInternalServerError, gin.H{"estado": "erro", "mensagem": "id inválido"})
			return
		}

		var contacto model.UpdateClienteInput
		if err := ctx.ShouldBindJSON(&contacto); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"estado": "erro", "mensagem": err.Error()})
			return
		}

		if err = serve.AlterarCliente(model.IDs{
			ClienteId: clienteId,
			UserID:    userID.(string),
		}, contacto); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"estado": "erro", "mensagem": err.Error()})
			return
		}

		ctx.JSON(http.StatusOK, gin.H{"mensagem": "cliente alterado"})
	}
}
