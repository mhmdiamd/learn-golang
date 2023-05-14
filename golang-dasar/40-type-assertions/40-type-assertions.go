package main

import "fmt"

func random() interface{} {
	return true
}

func main() {
	var result = random()
	// Mengguanakn type assertions
	// var resultString = result.(string)
	// fmt.Println(resultString)

	// mengguanakn type assertions dengan switch case

	switch value := result.(type) {
	case string:
		fmt.Println("Value", value, "is string")

	case int:
		fmt.Println("Value", value, "is integer")

	case bool:
		fmt.Println("Value", value, "is boolean")

	}
}
