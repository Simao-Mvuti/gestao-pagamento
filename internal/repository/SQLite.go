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

func (r *ClienteRepository) BuscarClienteID(index int) (model.Cliente, error) {
	query := "SELECT id,nome,contacto FROM clientes WHERE id = ?"
	var cliente model.Cliente
	rows := r.DB.QueryRow(query, index)

	err := rows.Scan(
		&cliente.ID,
		&cliente.Nome,
		&cliente.Contacto,
	)

	if err != nil {
		return model.Cliente{}, err
	}

	return cliente, nil

}

func (r *ClienteRepository) BuscarClientes() ([]model.Cliente, error) {
	query := "SELECT id, nome, contacto FROM clientes"

	rows, err := r.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	clientes := []model.Cliente{}

	for rows.Next() {
		var c model.Cliente

		err := rows.Scan(
			&c.ID,
			&c.Nome,
			&c.Contacto,
		)
		if err != nil {
			return nil, err
		}

		clientes = append(clientes, c)
	}

	return clientes, nil
}

func (r *ClienteRepository) AlterarCliente(index int, input model.UpdateClienteInput) error {
	query := "UPDATE clientes SET contacto = ? WHERE id = ?"
	if _, err := r.DB.Exec(query, input.Contacto, index); err != nil {
		return err
	}

	return nil
}

func (r *ClienteRepository) DeletarCliente(index int) error {
	query := "DELETE FROM clientes WHERE id = ?"
	if _, err := r.DB.Exec(query, index); err != nil {
		return err
	}

	return nil
}
