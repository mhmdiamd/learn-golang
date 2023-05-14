package main

import "fmt"

type Address struct {
	city, province, nation string
}

func main() {
	address1 := Address{"Tangsel", "Banten", "Indonesia"}
	address2 := address1

	address2.city = "Tangsel2"
	fmt.Println(address1) // City masih tangsel
	fmt.Println(address2) // City masih tangsel2

	// Pass by refrence mengguanakan tanda & "and"
	// Dengan mengguanakan tanda "&" untuk pass by refrence maka perubahan yang terjadi untuk data cloningnya akan memempengaruhi data utama
	address3 := &address1
	address3.city = "Bandung"
	fmt.Println(address1) // City Menjadi Bandung
	fmt.Println(address3) // City Menjadi Bandung
}
