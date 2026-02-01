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
		clienteRoute.DELETE("/:id", handler.DeleteClienteHandler(repository))
		clienteRoute.PUT("/:id", handler.PutClienteHandler(repository))
		clienteRoute.GET("/:id", handler.GetsClienteHandler(repository))
		clienteRoute.POST("/", handler.PostClienteHandler(repository))
		clienteRoute.GET("/", handler.GetsClientesHandler(repository))

	}

	c.Run(":8080")
}
