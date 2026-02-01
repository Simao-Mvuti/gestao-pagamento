package handler

import (
	serve "gestao/internal/servece"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func DeleteClienteHandler(serve *serve.ClienteService) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		input, ok := ctx.Params.Get("id")
		id, err := strconv.Atoi(input)
		if !ok || err != nil || id < 0 {
			ctx.JSON(http.StatusBadRequest, gin.H{"estado": "erro", "mensagem": err.Error()})
			return
		}

		if err = serve.DeletarCliente(id); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"estado": "erro", "mensagem": err.Error()})
			return
		}

		ctx.JSON(http.StatusOK, gin.H{"mensagem": "cliente deletado"})
	}
}
