package main

import (
	handlerCliente "gestao/internal/handler/cliente"
	handlerUsuario "gestao/internal/handler/usuario"
	"gestao/internal/middleware"
	"gestao/internal/repository"
	serve "gestao/internal/servece"

	"github.com/gin-gonic/gin"
)

func main() {
	c := gin.Default()

	db := repository.Connect()
	defer db.Close()

	// Usuário
	usuarioRepo := repository.NewUsuarioRepository(db)
	usuarioService := serve.NewUsuarioService(usuarioRepo) // retorna ponteiro

	// Cliente
	clienteRepo := repository.NewClienteRepository(db)
	clienteService := serve.NewClienteService(clienteRepo) // retorna ponteiro

	// Rotas
	c.POST("/usuario", handlerUsuario.Usuario(usuarioService))
	c.POST("/login", handlerUsuario.Login(usuarioService))

	clienteRoute := c.Group("/cliente")
	clienteRoute.Use(middleware.AuthMiddleware())
	{
		clienteRoute.DELETE("/:id", handlerCliente.DeleteClienteHandler(clienteService))
		clienteRoute.PUT("/:id", handlerCliente.PutClienteHandler(clienteService))
		clienteRoute.GET("/:id", handlerCliente.GetsClienteHandler(clienteService))
		clienteRoute.POST("/", handlerCliente.PostClienteHandler(clienteService))
		clienteRoute.GET("/", handlerCliente.GetsClientesHandler(clienteService))
	}

	c.Run(":8080")
}
