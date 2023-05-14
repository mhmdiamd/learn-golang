package main

import (
	"container/list"
	"fmt"
)

/*
	Package container/list adalah implementasi struktur data double linked list di Go-Lang
*/

func main() {
	data := list.New()
	data.PushBack("Muhamad")
	data.PushBack("Ilham")
	data.PushBack("Darmawan")

	data.PushFront("Muhamad Front")

	for element := data.Front(); element != nil; element = element.Next() {
		fmt.Println(*element)
	}

}
