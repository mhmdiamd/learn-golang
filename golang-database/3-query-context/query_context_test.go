package querycontext_test

import (
	"context"
	"database/sql"
	"fmt"
	"golang-database/database"
	"testing"
)

// Untuk operasi SQL yang tidak membutuhkan hasil, kita bisa menggunakan perintah Exec, namun jika kita membutuhkan result, seperti SELECT SQL, kita bisa menggunakan function yang berbeda
// Function untuk melakukan query ke database, bisa menggunakan function (DB) QueryContext(context, sql, params)

func RowIteration(rows *sql.Rows) (string, string) {
	var id, name string

	for rows.Next() {

		err := rows.Scan(&id, &name)

		if err != nil {
			panic(err)
		}

	}

	return id, name
}

func TestExecQuery(t *testing.T) {
	db := database.GetConnection()
	ctx := context.Background()

	defer db.Close()

	query := "SELECT * FROM customers"
	rows, err := db.QueryContext(ctx, query)

	if err != nil {
		panic(err)
	}

	// Hasil Query function adalah sebuah data structs sql.Rows
	// Rows digunakan untuk melakukan iterasi terhadap hasil dari query
	// Kita bisa menggunakan function (Rows) Next() (boolean) untuk melakukan iterasi terhadap data hasil query, jika return data false, artinya sudah tidak ada data lagi didalam result
	// Untuk membaca tiap data, kita bisa menggunakan (Rows) Scan(columns...)
	// Dan jangan lupa, setelah menggunakan Rows, jangan lupa untuk menutupnya menggunakan (Rows) Close()

	id, name := RowIteration(rows)

	fmt.Println("id : ", id)
	fmt.Println("name : ", name)

	defer rows.Close()
}
