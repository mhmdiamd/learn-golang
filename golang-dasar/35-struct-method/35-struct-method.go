package main

import "fmt"

/*
Struct adalah tipe data seperti tipe data lainnya, dia bisa digunakan sebagai parameter untuk function
Namun jika kita ingin menambahkan method ke dalam structs, sehingga seakan-akan sebuah struct memiliki function
*/

type Customer struct {
	name, address string
	age           int
}

func (customer Customer) sayHello() {
	fmt.Println("Hello " + customer.name)
}

func main() {
	customerName := Customer{name: "Ilham"}
	customerName.sayHello()
}
