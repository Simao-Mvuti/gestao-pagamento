package usecase

import (
	"gestao/internal/model"
	"strings"
	"unicode"
)

func ValidarCliente(cliente model.Cliente) bool {
	nome := strings.TrimSpace(cliente.Nome)
	contacto := strings.TrimSpace(cliente.Contacto)

	if nome == "" || contacto == "" {
		return false
	}

	if !SoLetras(nome) {
		return false
	}

	return true
}

func SoLetras(texto string) bool {
	novoTexto := []rune(texto)

	for _, valor := range novoTexto {
		if !unicode.IsLetter(valor) {
			return false
		}
	}

	return true
}
