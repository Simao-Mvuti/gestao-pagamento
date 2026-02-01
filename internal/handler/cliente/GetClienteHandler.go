package handler

import (
	serve "gestao/internal/servece"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetsClientesHandler(serve *serve.ClienteService) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		clientes, err := serve.BuscarClientes()

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
		id, err := strconv.Atoi(input)

		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"estado": "erro", "mensagem": err.Error()})
			return
		}

		if !ok || id < 0 {
			ctx.JSON(http.StatusInternalServerError, gin.H{"estado": "erro", "mensagem": err.Error()})
			return
		}

		cliente, err := serve.BuscarClienteID(id)
		if err != nil {
			ctx.JSON(http.StatusNotFound, gin.H{"estado": "erro", "mensagem": err.Error()})
			return
		}

		ctx.JSON(http.StatusOK, cliente)
	}
}
