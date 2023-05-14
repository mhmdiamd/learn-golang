package main

import (
	"fmt"
	"regexp"
)

/*
	- Package regexp adalah utilitas di Go-Lang untuk melakukan pencarian regular expression
	- Regular expression di Go-Lang menggunakan library C yang dibuat Google bernama RE2
*/

func main() {

	// Membuat regex yang mana menetukan string yang depanya harus e dan akhirnya harus o dan tengah2 bolah hurif a-z
	var regex *regexp.Regexp = regexp.MustCompile("e([a-z])o")

	fmt.Println(regex.MatchString("eko"))
	fmt.Println(regex.MatchString("eki"))
	fmt.Println(regex.MatchString("edo"))

	search := regex.FindAllString("eko edi edo eto e1o elo owo", 2) // 2 adalah maksimal data yg di return
	fmt.Println(search)
}
