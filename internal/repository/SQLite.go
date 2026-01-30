package repository

import (
	"database/sql"
	"gestao/internal/model"

	"log"

	_ "github.com/mattn/go-sqlite3"
)

func Connect() *sql.DB {
	db, err := sql.Open("sqlite3", "/home/simao/Desktop/codigo/gestao/store/clientes.db")

	if err != nil {
		log.Fatal("Erro ao abrir o banco:", err)
	}

	err = db.Ping()

	if err != nil {
		log.Fatal("Erro ao conectar com o banco:", err)
	}

	return db
}

type ClienteRepository struct {
	DB *sql.DB
}

func NewClienteRepository(db *sql.DB) *ClienteRepository {
	return &ClienteRepository{
		DB: db,
	}
}

func (r *ClienteRepository) CadastrarCliente(cliente model.Cliente) error {
	query := "INSERT INTO clientes (nome, contacto) VALUES (?,?)"
	_, err := r.DB.Exec(query, cliente.Nome, cliente.Contacto)
	if err != nil {
		return err
	}
	return nil
}
