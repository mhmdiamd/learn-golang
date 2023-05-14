package main

import (
	"fmt"
	"os"
)

/*
Package os berisikan fungsionalitas untuk mengakses fitur sistem operasi secara independen
(bisa digunakan  disemua sistem operasi)
*/

func main() {
	args := os.Args

	fmt.Println(args)

	hostname, err := os.Hostname()

	if err == nil {
		fmt.Println(hostname)
	} else {
		fmt.Println("Error", err.Error())
	}
}
