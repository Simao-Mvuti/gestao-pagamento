package serve

import "gestao/internal/model"

type ClienteRepository interface {
	CadastrarCliente(cliente model.Cliente) error
	BuscarClientes(userID string) ([]model.Cliente, error)
	BuscarClienteID(ids model.IDs) (model.Cliente, error)
	DeletarCliente(ids model.IDs) error
	AlterarCliente(ids model.IDs, input model.UpdateClienteInput) error
}

type UsuarioRepository interface {
	CriarConta(input model.Usuario) error
	Logar(input model.Usuario) (string, error)
}
