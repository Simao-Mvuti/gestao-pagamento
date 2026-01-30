package model

type Cliente struct {
	Nome     string `json:"nome" binding:"requered",min=3`
	Contacto string `json:"contacto" binding:"requered",min=5`
}
