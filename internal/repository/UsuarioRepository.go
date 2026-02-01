package repository

import (
	"database/sql"
	"errors"
	"gestao/internal/model"
	"os"

	"golang.org/x/crypto/bcrypt"
)

type SqliteUsuarioRepository struct {
	DB *sql.DB
}

func NewUsuarioRepository(db *sql.DB) *SqliteUsuarioRepository {
	return &SqliteUsuarioRepository{
		DB: db,
	}
}

func (r *SqliteUsuarioRepository) CriarConta(input model.Usuario) error {
	query := "INSERT INTO usuaios (email,senha) VALUES (?,?) "
	hash, err := bcrypt.GenerateFromPassword([]byte(input.Senha), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	_, err = r.DB.Exec(query, input.Email, string(hash))
	if err != nil {
		return err
	}
	return nil
}

func (r *SqliteUsuarioRepository) Logar(input model.Usuario) (string, error) {
	query := "SELECT senha FROM usuaios WHERE email = ?"
	interno := model.Usuario{}
	row := r.DB.QueryRow(query, input.Email)

	err := row.Scan(&interno.Senha)
	if err != nil {
		return "", errors.New("credenciais inválidas")
	}

	// comparar hash
	err = bcrypt.CompareHashAndPassword([]byte(interno.Senha), []byte(input.Senha))
	if err != nil {
		return "", errors.New("credenciais inválidas")
	}

	// retornar token (pode ser temporário por enquanto)
	return os.Getenv("TOKEN"), nil

}
