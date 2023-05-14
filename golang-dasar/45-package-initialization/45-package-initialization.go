package main

import (
	"fmt"
	"golang-dasar/45-package-initialization/database"
)

func main() {

	connection := database.GetDatabase()
	fmt.Println(connection)
}
