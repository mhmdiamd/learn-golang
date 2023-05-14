package main

import "fmt"

type Customer struct {
	name, address string
	age           int
}

func main() {
	var ilham Customer

	ilham.name = "Ilham"
	ilham.address = "Pd salak"
	ilham.age = 12

	fmt.Println(ilham)

	// Using Struct literal

	var ilham2 = Customer{
		name:    "Muhamad Ilham Darmawan",
		address: "pd salak",
		age:     17,
	}

	fmt.Println(ilham2)
}
