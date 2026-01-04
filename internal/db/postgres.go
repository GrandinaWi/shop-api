package db

import "database/sql"

func NewPostgresDB() (*sql.DB, error) {
	dsn := "postgres://shop:shop@localhost:5432/postgresshop?sslmode=disable"
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		return nil, err
	}
	return db, nil
}
