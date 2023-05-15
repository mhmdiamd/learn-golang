package main

import (
	"embed"
	"fmt"
	"io/fs"
	"io/ioutil"
)

// go:embed version.txt
var version string

// go:embed download.jpeg
var logo []byte

var path embed.FS

func main() {
	fmt.Println(version)

	ioutil.WriteFile("logo_next.png", logo, fs.ModePerm)
}
