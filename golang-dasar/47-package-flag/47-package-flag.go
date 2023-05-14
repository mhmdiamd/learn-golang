package main

// Package flag berisikan fungsionalitas untuk memparsing command line argument

import (
	"flag"
	"fmt"
)

func main() {
	host := flag.String("host", "localhost", "Put your database host")
	username := flag.String("username", "root", "Put your database username")
	password := flag.String("password", "root", "Put your database password")

	flag.Parse()

	fmt.Println(*host, *username, *password)
}
