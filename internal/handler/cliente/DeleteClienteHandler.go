package handler

import (
	"gestao/internal/model"
	serve "gestao/internal/servece"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func DeleteClienteHandler(serve *serve.ClienteService) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		input, ok := ctx.Params.Get("id")
		userID, ok := ctx.Get("ser_id")

		clienteId, err := strconv.Atoi(input)

		if !ok || err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"estado": "erro", "mensagem": "id inválido"})
			return
		}

		if err = serve.DeletarCliente(model.IDs{
			ClienteId: clienteId,
			UserID:    userID.(string),
		}); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"estado": "erro", "mensagem": err.Error()})
			return
		}

		ctx.JSON(http.StatusOK, gin.H{"mensagem": "cliente deletado"})
	}
}
