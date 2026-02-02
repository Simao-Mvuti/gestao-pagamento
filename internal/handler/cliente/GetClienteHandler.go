package handler

import (
	"gestao/internal/middleware"
	"gestao/internal/model"
	serve "gestao/internal/servece"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetsClientesHandler(serve *serve.ClienteService) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		userID, ok := ctx.Get(middleware.USER_ID)
		if !ok {
			ctx.JSON(http.StatusNotFound, gin.H{"estado": "erro", "mensagem": "id invalido"})
			return
		}

		clientes, err := serve.BuscarClientes(userID.(string))

		if err != nil {
			ctx.JSON(http.StatusNotFound, gin.H{"estado": "erro", "mensagem": err.Error()})
			return
		}

		ctx.JSON(http.StatusOK, clientes)
	}
}

func GetsClienteHandler(serve *serve.ClienteService) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		input, ok := ctx.Params.Get("id")
		userID, ok := ctx.Get(middleware.USER_ID)

		clienteId, err := strconv.Atoi(input)

		if !ok || err != nil || clienteId < 0 {
			ctx.JSON(http.StatusNotFound, gin.H{"estado": "erro", "mensagem": "id invalido"})
			return
		}

		cliente, err := serve.BuscarClienteID(model.IDs{
			ClienteId: clienteId,
			UserID:    userID.(string),
		})
		if err != nil {
			ctx.JSON(http.StatusNotFound, gin.H{"estado": "erro", "mensagem": err.Error()})
			return
		}

		ctx.JSON(http.StatusOK, cliente)
	}
}
