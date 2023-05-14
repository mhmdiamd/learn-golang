package main

import "fmt"

// Recursive function adalah function yang memanggil function dirinya sendiri
// Contoh kasus yang lebih mudah diselesaikan menggunakan recursive adalah Factorial

func factorialForLoop(value int) int {
	result := 1

	for i := value; i > 0; i-- {
		result *= i
	}

	return result
}

func factorialRecursive(value int) int {
	if value == 1 {
		return 1
	} else {
		return value * factorialRecursive(value-1)
	}
}

func main() {

	fmt.Println(factorialForLoop(5))
	fmt.Println(factorialRecursive(5))
}
