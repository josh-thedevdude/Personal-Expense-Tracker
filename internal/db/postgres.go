package db

import (
	"database/sql"
	"os"
)

func NewPostgresDB() (*sql.DB, error) {
	connStr := os.Getenv("DB_URL")

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}
