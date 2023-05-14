package main

import "fmt"

type Man struct {
	Name string
}

func (man *Man) Married() {
	man.Name = "Mr. " + man.Name
}

func main() {
	var ilham = Man{"ilham"}
	ilham.Married()

	fmt.Println(ilham)
}
