package openconnection_test

import (
	"database/sql"
	"testing"

	_ "github.com/lib/pq"
)

func TestOpenConnection(t *testing.T) {
	connStr := "postgres://postgres:ilham@localhost:5432/golang_belajar_db?sslmode=verify-full"
	db, err := sql.Open("postgres", connStr)

	if err != nil {
		panic(err)
	}

	defer db.Close()
}
