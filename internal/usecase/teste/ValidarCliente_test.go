package usecase

import (
	"gestao/internal/model"
	"gestao/internal/usecase"
	"testing"
)

func testValidarCliente(t *testing.T) {
	clienteUm := model.Cliente{
		Nome:     "si1",
		Contacto: "926387504",
	}

	clienteDois := model.Cliente{
		Nome:     "Simao",
		Contacto: "926387504",
	}

	res := usecase.ValidarCliente(clienteUm)
	resDois := usecase.ValidarCliente(clienteDois)

	if !res {
		t.Error()
	}

	if !resDois {
		t.Error()
	}
}
