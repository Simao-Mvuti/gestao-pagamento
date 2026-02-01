package repository

import (
	"database/sql"

	"log"

	_ "github.com/mattn/go-sqlite3"
)

func Connect() *sql.DB {
	db, err := sql.Open("sqlite3", "/home/simao/Desktop/codigo/gestao/store/banco.db")

	if err != nil {
		log.Fatal("Erro ao abrir o banco:", err)
	}

	err = db.Ping()

	if err != nil {
		log.Fatal("Erro ao conectar com o banco:", err)
	}

	return db
}
