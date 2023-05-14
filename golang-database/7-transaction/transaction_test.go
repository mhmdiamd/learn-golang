package transaction_test

import (
	"context"
	"fmt"
	"golang-database/database"
	"strconv"
	"testing"
)

func TestTransaction(t *testing.T) {
	db := database.GetConnection()

	tx, err := db.Begin()

	if err != nil {
		panic(err)
	}

	context := context.Background()
	// Proses Datbase Transaction
	sqlQuery := "INSERT INTO admin VALUES($1, $2, $3)"

	for i := 0; i < 10; i++ {

		usernameAdmin := "admin" + strconv.Itoa(i)
		passwordAdmin := "admin"
		_, err := tx.ExecContext(context, sqlQuery, strconv.Itoa(i), usernameAdmin, passwordAdmin)

		if err != nil {
			panic(err)
		}
	}

	// End Proses database transaction
	tx.Commit()

	fmt.Println("Success Insert to database")
}
