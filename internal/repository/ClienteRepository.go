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
	query := `
		INSERT INTO clientes (nome, contacto, user_id)
		VALUES (?, ?, ?)
	`

	_, err := r.DB.Exec(
		query,
		cliente.Nome,
		cliente.Contacto,
		cliente.UserID,
	)

	return err
}

func (r *SqliteClienteRepository) BuscarClienteID(ids model.IDs) (model.Cliente, error) {

	query := `
		SELECT id, nome, contacto, user_id
		FROM clientes
		WHERE id = ? AND user_id = ?
	`

	var cliente model.Cliente

	row := r.DB.QueryRow(
		query,
		ids.ClienteId,
		ids.UserID,
	)

	err := row.Scan(
		&cliente.ID,
		&cliente.Nome,
		&cliente.Contacto,
		&cliente.UserID,
	)

	if err != nil {
		return model.Cliente{}, err
	}

	return cliente, nil
}
func (r *SqliteClienteRepository) BuscarClientes(userID string) ([]model.Cliente, error) {

	query := `
		SELECT id, nome, contacto, user_id
		FROM clientes
		WHERE user_id = ?
	`

	rows, err := r.DB.Query(query, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var clientes []model.Cliente

	for rows.Next() {
		var c model.Cliente

		err := rows.Scan(
			&c.ID,
			&c.Nome,
			&c.Contacto,
			&c.UserID,
		)

		if err != nil {
			return nil, err
		}

		clientes = append(clientes, c)
	}

	return clientes, nil
}

func (r *SqliteClienteRepository) AlterarCliente(
	ids model.IDs,
	input model.UpdateClienteInput,
) error {

	query := `
		UPDATE clientes
		SET contacto = ?
		WHERE id = ? AND user_id = ?
	`

	_, err := r.DB.Exec(
		query,
		input.Contacto,
		ids.ClienteId,
		ids.UserID,
	)

	return err
}

func (r *SqliteClienteRepository) DeletarCliente(ids model.IDs) error {

	query := `
		DELETE FROM clientes
		WHERE id = ? AND user_id = ?
	`

	_, err := r.DB.Exec(
		query,
		ids.ClienteId,
		ids.UserID,
	)

	return err
}
