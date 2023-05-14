package fixsqlinjection

import (
	"context"
	"fmt"
	"golang-database/database"
	"testing"
)

func TestSqlInjectinSafe(t *testing.T) {
	db := database.GetConnection()
	defer db.Close()

	var context = context.Background()

	usernameAdmin := "admin"
	password := "password"

	rows, err := db.QueryContext(context, `SELECT username FROM admin WHERE username = $1 AND password = $2 LIMIT 1`, usernameAdmin, password)
	// rows, err := db.QueryContext(context, query, username, password)

	if err != nil {
		panic(err)
	}

	var username string
	for rows.Next() {
		err := rows.Scan(&username)

		if err != nil {
			panic(err)
		}
	}

	fmt.Println(username)

	defer rows.Close()

}
