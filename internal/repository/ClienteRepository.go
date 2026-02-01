package repository

import (
	"database/sql"
	"gestao/internal/model"
)

type SqliteClienteRepository struct {
	DB *sql.DB
}

func NewClienteRepository(db *sql.DB) *SqliteClienteRepository {
	return &SqliteClienteRepository{
		DB: db,
	}
}

func (r *SqliteClienteRepository) CadastrarCliente(cliente model.Cliente) error {
	query := "INSERT INTO clientes (nome, contacto) VALUES (?,?)"
	_, err := r.DB.Exec(query, cliente.Nome, cliente.Contacto)
	if err != nil {
		return err
	}
	return nil
}

func (r *SqliteClienteRepository) BuscarClienteID(index int) (model.Cliente, error) {
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

func (r *SqliteClienteRepository) BuscarClientes() ([]model.Cliente, error) {
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

func (r *SqliteClienteRepository) AlterarCliente(index int, input model.UpdateClienteInput) error {
	query := "UPDATE clientes SET contacto = ? WHERE id = ?"
	if _, err := r.DB.Exec(query, input.Contacto, index); err != nil {
		return err
	}

	return nil
}

func (r *SqliteClienteRepository) DeletarCliente(index int) error {
	query := "DELETE FROM clientes WHERE id = ?"
	if _, err := r.DB.Exec(query, index); err != nil {
		return err
	}

	return nil
}
