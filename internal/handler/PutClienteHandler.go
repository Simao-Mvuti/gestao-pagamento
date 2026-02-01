package handler

import (
	"gestao/internal/model"
	serve "gestao/internal/servece"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func PutClienteHandler(serve serve.Repository) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		input, ok := ctx.Params.Get("id")
		var contacto model.UpdateClienteInput
		if err := ctx.ShouldBindJSON(&contacto); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"estado": "erro", "mensagem": err.Error()})
			return
		}
		id, err := strconv.Atoi(input)
		if !ok || err != nil || id < 0 {
			ctx.JSON(http.StatusBadRequest, gin.H{"estado": "erro", "mensagem": err.Error()})
			return
		}

		if err = serve.AlterarCliente(id, contacto); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"estado": "erro", "mensagem": err.Error()})
			return
		}

		ctx.JSON(http.StatusOK, gin.H{"mensagem": "cliente alterado"})
	}
}
