package main

import "fmt"

func sayHelloWithFilter(name string, filter func(string) string) {
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

	// Example 1
	sayHelloWithFilter("Anjing", spamFilter)

	// Example 2
	filter := spamFilter
	sayHelloWithFilter("Ilham", filter)
}
