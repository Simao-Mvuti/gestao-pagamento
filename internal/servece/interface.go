package serve

import "gestao/internal/model"

type ClienteRepository interface {
	CadastrarCliente(cliente model.Cliente) error
	BuscarClientes() ([]model.Cliente, error)
	BuscarClienteID(index int) (model.Cliente, error)
	DeletarCliente(index int) error
	AlterarCliente(index int, input model.UpdateClienteInput) error
}

type UsuarioRepository interface {
	CriarConta(input model.Usuario) error
	Logar(input model.Usuario) (string, error)
}
