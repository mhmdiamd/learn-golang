package dbinsert_test

import (
	"context"
	"fmt"
	"golang-database/database"
	_ "golang-database/database"
	"testing"
)

/*
	- sql.DB di Golang sebenarnya bukanlah sebuah koneksi ke database
	- Melainkan sebuah pool ke database, atau dikenal dengan konsep Database Pooling
	- Di dalam sql.DB, Golang melakukan management koneksi ke database secara otomatis. Hal ini menjadikan kita tidak perlu melakukan management koneksi database secara manual
	- Dengan kemampuan database pooling ini, kita bisa menentukan jumlah minimal dan maksimal koneksi yang dibuat oleh Golang, sehingga tidak membanjiri koneksi ke database, karena biasanya ada batas maksimal koneksi yang bisa ditangani oleh database yang kita gunakan
*/

func TestExecSql(t *testing.T) {
	db := database.GetConnection()
	context := context.Background()

	defer db.Close()

	query := "INSERT INTO customers(id, name) VALUES('ilham', 'Ilham')"
	_, err := db.ExecContext(context, query)

	if err != nil {
		panic(err)
	}

	fmt.Println("Success insert to database")

}
