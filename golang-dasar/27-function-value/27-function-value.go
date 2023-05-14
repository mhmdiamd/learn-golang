package main

import "fmt"

func goodByeSay(name string) string {
	return "Hello" + name
}

func main() {
	goodBye := goodByeSay

	fmt.Println(goodBye("Ilham"))
}
