package main

import (
	"fmt"
	"golang-roadmap/golang-database/database"
)

func main() {

	connection := database.GetConnection()
	fmt.Println(connection)
}
