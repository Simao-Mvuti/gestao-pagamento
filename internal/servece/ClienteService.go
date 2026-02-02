package serve

import (
	"errors"
	"gestao/internal/model"
	"gestao/internal/usecase"
)

type ClienteService struct {
	Re ClienteRepository
}

func NewClienteService(cliente ClienteRepository) *ClienteService {
	return &ClienteService{
		Re: cliente,
	}
}

func (re *ClienteService) CadastrarCliente(cliente model.Cliente) error {
	if usecase.ValidarCliente(cliente) {
		return errors.New("Dados Invalidos")
	}
	return re.Re.CadastrarCliente(cliente)
}

func (re *ClienteService) BuscarClientes(userID string) ([]model.Cliente, error) {
	return re.Re.BuscarClientes(userID)
}

func (re *ClienteService) BuscarClienteID(ids model.IDs) (model.Cliente, error) {
	return re.Re.BuscarClienteID(ids)
}

func (re *ClienteService) DeletarCliente(ids model.IDs) error {
	return re.Re.DeletarCliente(ids)
}

func (re *ClienteService) AlterarCliente(ids model.IDs, input model.UpdateClienteInput) error {
	return re.Re.AlterarCliente(ids, input)
}
