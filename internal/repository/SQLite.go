package repository

import (
	"database/sql"
	"os"

	"log"

	_ "github.com/mattn/go-sqlite3"
)

func Connect() *sql.DB {
	pathBanco := os.Getenv("DB_PATH")

	if pathBanco == "" {
		log.Fatal("DB_PATH vazia")
	}

	db, err := sql.Open("sqlite3", pathBanco)

	if err != nil {
		log.Fatal("Erro ao abrir o banco:", err)
	}

	err = db.Ping()

	if err != nil {
		log.Fatal("Erro ao conectar com o banco:", err)
	}

	return db
}
