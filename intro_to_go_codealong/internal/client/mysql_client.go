package client

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

func GetMySQLDB() (*sql.DB, error) {
	db, err := sql.Open("mysql", "docker:docker@tcp(127.0.0.1)/introtogo")
	if err != nil {
		return nil, fmt.Errorf("error opening MySQL pool: %w", err)
	}

	return db, nil
}
