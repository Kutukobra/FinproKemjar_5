package database

import (
	"database/sql"
)

func NewPostgresDatabase(connectionString string) (*sql.DB, error) {
	driver, err := sql.Open("postgres", connectionString)
	if err != nil {
		return nil, err
	}
	return driver, nil
}
