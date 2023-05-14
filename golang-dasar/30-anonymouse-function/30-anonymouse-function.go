package main

import "fmt"

type BlackList func(string) bool

func registerUser(name string, blackList BlackList) {
	if blackList(name) {
		fmt.Println("you are blocked...")
	} else {
		fmt.Println("Hello " + name)
	}
}

func blakcListWords(name string) bool {
	return name == "Anjing"
}

func main() {
	registerUser("Ilham", blakcListWords)
	registerUser("Anjing", func(name string) bool {
		return name == "Anjing"
	})
}
