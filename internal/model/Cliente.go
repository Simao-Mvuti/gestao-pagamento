package model

type Cliente struct {
	ID       int    `json:"id"`
	Nome     string `json:"nome"`
	Contacto string `json:"contacto"`
	UserID   string `json:"user_id"`
}

type ClienteDTO struct {
	Nome     string `json:"nome" binding:"required,min=9`
	Contacto string `json:"contacto" binding:"required,min=2`
}

type IDs struct {
	UserID    string
	ClienteId int
}

func ToClienteDTO(cliente Cliente) ClienteDTO {
	return ClienteDTO{
		Nome:     cliente.Nome,
		Contacto: cliente.Contacto,
	}
}

//Adicionar opcao de ativo e desativado para excluir

type UpdateClienteInput struct {
	Contacto string `json:"contacto"`
}

type Usuario struct {
	ID    string `json:"user_id"`
	Email string `json:"email"`
	Senha string `json"senha"`
}

type UsuarioLogin struct {
	Email string `json:"email"`
	Senha string `json:"senha"`
}
