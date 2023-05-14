package main

import "fmt"

// Parameter yang berada di posisi terakhir, memiliki kemampuan dijadikan sebuah varargs
// Varargs artinya datanya bisa menerima lebih dari satu input, atau anggap saja semacam Array.

func sum(numbers ...int) int {
	total := 0
	for _, number := range numbers {
		total += number
	}

	return total
}

func main() {
	total := sum(1, 2, 3, 4, 5)
	fmt.Println(total)
}
