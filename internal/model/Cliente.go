package model

type Cliente struct {
	ID       int    `json:"id"`
	Nome     string `json:"nome"`
	Contacto string `json:"contacto"`
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
