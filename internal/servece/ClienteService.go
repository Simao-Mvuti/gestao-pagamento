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

func (re *ClienteService) BuscarClientes(userID string) ([]model.ClienteDTO, error) {
	var clientesDto []model.ClienteDTO
	clientes, err := re.Re.BuscarClientes(userID)
	if err != nil {
		return []model.ClienteDTO{}, err
	}

	for _, cliente := range clientes {
		clientesDto = append(clientesDto, model.ToClienteDTO(cliente))
	}

	return clientesDto, nil
}

func (re *ClienteService) BuscarClienteID(ids model.IDs) (model.ClienteDTO, error) {
	cliente, err := re.Re.BuscarClienteID(ids)
	if err != nil {
		return model.ClienteDTO{}, err
	}

	return model.ToClienteDTO(cliente), nil

}

func (re *ClienteService) DeletarCliente(ids model.IDs) error {
	return re.Re.DeletarCliente(ids)
}

func (re *ClienteService) AlterarCliente(ids model.IDs, input model.UpdateClienteInput) error {
	return re.Re.AlterarCliente(ids, input)
}
