package main

import (
	"gestao/internal/handler"
	"gestao/internal/repository"

	"github.com/gin-gonic/gin"
)

func main() {
	c := gin.Default()
	db := repository.Connect()
	defer db.Close()

	repository := repository.NewClienteRepository(db)
	clienteRoute := c.Group("/cliente")
	{
		clienteRoute.POST("/", handler.PostClienteHandler(repository))
	}

	c.Run(":8080")
}
