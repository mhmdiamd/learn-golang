package main

import (
	"fmt"
	"strings"
)

/*
Package strings adalah package yang berisikan function-function untuk memanipulasi tipe data String
Ada banyak sekali function yang bisa kita gunakan
*/

func main() {
	fmt.Println(strings.Contains("Muhamad Ilham", "Ilham"))
	fmt.Println(strings.Split("Muhamad Ilham", " ")) // Mecah string jadi array
	fmt.Println(strings.ToLower("Muhamad Ilham"))
	fmt.Println(strings.ToUpper("Muhamad Ilham"))
	fmt.Println(strings.Trim("    Muhamad Ilham    ", " "))
	fmt.Println(strings.ReplaceAll("ilham ilham ilham", "ilham", "darmawan"))
}
