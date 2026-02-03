package usecase

import (
	"gestao/internal/usecase"
	"testing"
)

func TestSoLetras(t *testing.T) {
	verificarResultado := func(t *testing.T, resultado bool, esperado bool) {
		t.Helper()
		if resultado != esperado {
			t.Errorf("Resultado %t ,Esperado %t", resultado, esperado)
		}
	}

	t.Run("Testanto com Letras apenas", func(t *testing.T) {
		resultado := usecase.SoLetras("samuel")
		esperado := true
		verificarResultado(t, resultado, esperado)
	})

	t.Run("Testando Com letras e numero", func(t *testing.T) {
		resultado := usecase.SoLetras("samuel123")
		esperado := false
		verificarResultado(t, resultado, esperado)
	})

	t.Run("Com simbolos e letras", func(t *testing.T) {
		resultado := usecase.SoLetras("S@muel?")
		esperado := false
		verificarResultado(t, resultado, esperado)
	})
}
