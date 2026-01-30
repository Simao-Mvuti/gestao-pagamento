package serve

import (
	"errors"
	"gestao/internal/model"
	"gestao/internal/usecase"
)

func ClienteService(re Repository, cliente model.Cliente) error {
	if usecase.ValidarCliente(cliente) {
		return errors.New("Dados Invalidos")
	}
	return re.CadastrarCliente(cliente)
}
