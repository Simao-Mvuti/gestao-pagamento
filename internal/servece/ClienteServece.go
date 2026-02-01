package serve

import (
	"errors"
	"gestao/internal/model"
	"gestao/internal/usecase"
)

type ClienteService struct {
	Re Repository
}

func (re ClienteService) CadastrarCliente(cliente model.Cliente) error {
	if usecase.ValidarCliente(cliente) {
		return errors.New("Dados Invalidos")
	}
	return re.Re.CadastrarCliente(cliente)
}

func (re ClienteService) BuscarClientes() ([]model.Cliente, error) {
	return re.Re.BuscarClientes()
}

func (re ClienteService) BuscarClienteID(index int) (model.Cliente, error) {
	return re.Re.BuscarClienteID(index)
}

func (re ClienteService) DeletarCliente(index int) error {
	return re.Re.DeletarCliente(index)
}

func (re ClienteService) AlterarCliente(index int, input model.UpdateClienteInput) error {
	return re.Re.AlterarCliente(index, input)
}
