package usecase

import (
	"gestao/internal/usecase"
	"testing"
)

func testeSoLetras(t *testing.T) {
	texto := "samuel"
	texto2 := "samuel123"

	if !usecase.SoLetras(texto) {
		t.Error()
	}

	if usecase.SoLetras(texto2) {
		t.Error()
	}
}
