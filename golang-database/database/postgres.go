package database

import (
	"database/sql"
	"time"

	_ "github.com/lib/pq"
)

func GetConnection() *sql.DB {
	connStr := "user=postgres password=ilham dbname=golang_belajar_db port=5432 host=localhost"
	db, err := sql.Open("postgres", connStr)

	if err != nil {
		panic(err)
	}

	db.SetMaxIdleConns(10)
	db.SetMaxOpenConns(50)
	db.SetConnMaxIdleTime(5 * time.Minute)
	db.SetConnMaxLifetime(60 * time.Minute)

	return db
}
