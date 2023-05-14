package preparestatement_test

import (
	"context"
	"fmt"
	"golang-database/database"
	"strconv"
	"testing"
)

/*
	- Jika menggunakan Query dan Exec dengan parameter, kita tidak bisa menjamin bahwa koneksi yang digunakan akan sama, oleh karena itu, bisa jadi prepare statement akan selalu dibuat berkali-kali walaupun kita menggunakan SQL yang sama
	- Untuk membuat Prepare Statement, kita bisa menggunakan function (DB) Prepare(context, sql)
	- Prepare Statement direpresentasikan dalam struct database/sql.Stmt
	- Sama seperti resource sql lainnya, Stmt harus di Close() jika sudah tidak digunakan lagi
*/

func TestPrepareStatement0(t *testing.T) {
	var context = context.Background()

	db := database.GetConnection()
	defer db.Close()

	stmt, err := db.PrepareContext(context, "INSERT INTO admin VALUES($1, $2, $3)")

	if err != nil {
		panic(err)
	}

	defer stmt.Close()

	for i := 0; i < 10; i++ {

		username := "admin" + strconv.Itoa(i)
		password := "password"

		result, err := stmt.ExecContext(context, strconv.Itoa(i), username, password)

		if err != nil {
			panic(err)
		}

		lastInsertId, _ := result.LastInsertId()

		fmt.Println(lastInsertId)

	}

}
