package serve

import "gestao/internal/model"

type Repository interface {
	CadastrarCliente(cliente model.Cliente) error
}
