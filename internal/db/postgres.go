package db

import (
	"database/sql"
	"os"

	"github.com/joho/godotenv"
)

func NewPostgresDB() (*sql.DB, error) {
	_ = godotenv.Load(".env")
	dsn := os.Getenv("DATABASE_URL")
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		return nil, err
	}
	return db, nil
}
