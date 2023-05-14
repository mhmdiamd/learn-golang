package main

import "fmt"

func getFullName() (firstName string, middleName string, lastName string) {
	firstName = "Muhamad"
	middleName = "Ilham"
	lastName = "Darmawan"

	return firstName, middleName, lastName
}

func main() {
	a, b, c := getFullName()

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
}
