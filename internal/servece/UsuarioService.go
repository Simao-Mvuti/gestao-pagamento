package serve

import (
	"gestao/internal/model"
)

type UsuarioService struct {
	Re UsuarioRepository
}

func NewUsuarioService(usuario UsuarioRepository) *UsuarioService {
	return &UsuarioService{
		Re: usuario,
	}
}

func (r *UsuarioService) CriarConta(input model.Usuario) error {
	return r.Re.CriarConta(input)
}

func (r *UsuarioService) Logar(input model.Usuario) (string, error) {
	return r.Re.Logar(input)
}
