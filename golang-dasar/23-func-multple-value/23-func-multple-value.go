package main

import "fmt"

func getFullName() (string, string, string) {
	return "Muhamad", "Ilham", "Darmawan"
}

func main() {
	firstName, middleName, lastName := getFullName()
	fmt.Println(firstName)
	fmt.Println(middleName)
	fmt.Println(lastName)
}
