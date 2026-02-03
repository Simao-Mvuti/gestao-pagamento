package usecase

import (
	"gestao/internal/model"
	"gestao/internal/usecase"
	"testing"
)

func testValidarCliente(t *testing.T) {
	validar := func(t *testing.T, resposta bool, esperado bool) {
		t.Helper()

		if resposta != esperado {
			t.Errorf("Resposta %t, Mas esperava %t", resposta, esperado)
		}

	}

	t.Run("Testando com dados validos", func(t *testing.T) {
		resposta := usecase.ValidarCliente(model.Cliente{
			Nome:     "Simao",
			Contacto: "926387504",
		})

		esperado := true
		validar(t, resposta, esperado)

	})

	t.Run("Testando com dados invalidos", func(t *testing.T) {
		resposta := usecase.ValidarCliente(model.Cliente{
			Nome:     "Si@mao?",
			Contacto: "92fv87504",
		})

		esperado := false
		validar(t, resposta, esperado)

	})
}
