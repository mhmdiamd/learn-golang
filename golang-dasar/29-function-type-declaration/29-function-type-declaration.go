package main

import "fmt"

// Type Declarations juga bisa digunakan untuk membuat alias function, sehingga akan mempermudah kita menggunakan function sebagai parameter

type Filter func(string) string

func sayHelloWithFilter(name string, filter Filter) {
	fmt.Println("Hello " + filter(name))
}

func spamFilter(name string) string {
	if name == "Anjing" {
		return "...."
	} else {
		return name
	}
}

func main() {

	sayHelloWithFilter("Anjing", spamFilter)

	var filter = spamFilter
	sayHelloWithFilter("Ilham", filter)
}
