package tipedatacolumn_test

import (
	"context"
	"database/sql"
	"fmt"
	"golang-database/database"
	"testing"
	"time"
)

/*
	Sebelumnya kita hanya membuat table dengan tipe data di kolom nya berupa VARCHAR
	Untuk VARCHAR di database, biasanya kita gunakan String di Golang
	Apa representasinya di Golang, misal tipe data timestamp, date dan lain-lain

	Varchar, Char -> string
	int, bigint -> int32, int64
	float, numeric/double -> float32, float64
	boolean -> bool
	date, datetime, time, timestamp -> time.Time
*/

type User struct {
	id, name, email        string
	balance                int
	rating                 float32
	birth_date, created_at time.Time
	married                bool
}

func RowsIteration(rows *sql.Rows) User {
	var id, name, email string
	var balance int
	var rating float32
	var birth_date, created_at time.Time
	var married bool

	for rows.Next() {

		err := rows.Scan(&id, &name, &email, &balance, &rating, &birth_date, &created_at, &married)

		if err != nil {
			panic(err)
		}

	}

	return User{
		id:         id,
		name:       name,
		email:      email,
		balance:    balance,
		rating:     rating,
		birth_date: birth_date,
		created_at: created_at,
		married:    married,
	}
}

func TestInsertNewColumn(t *testing.T) {
	db := database.GetConnection()
	ctx := context.Background()

	// query := "INSERT INTO customers(id, name, email, balance, rating, birth_date, married) VALUES('ilham','Ilham Darmawan', 'am@gmail.com', 100000, 5.0, '1999-9-9', true)"
	// _, err := db.ExecContext(ctx, query)

	query := "SELECT * FROM customers WHERE id='ilham'"
	rows, err := db.QueryContext(ctx, query)

	if err != nil {
		panic(err)
	}

	user := RowsIteration(rows)

	fmt.Println(
		"id", user.id,
		"name", user.name,
		"balance", user.balance,
		"rating", user.rating,
		"birth_date", user.birth_date,
		"createdAt", user.created_at,
		"married", user.married,
	)

	// fmt.Println("Success insert data customers")
}
