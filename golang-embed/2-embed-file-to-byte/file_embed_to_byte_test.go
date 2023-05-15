package embedfiletobyte

import (
	_ "embed"
	"io/fs"
	"io/ioutil"
	"testing"
)

//go:ember download.jpeg
var photoByte []byte

func TestFileToByte(t *testing.T) {

	err := ioutil.WriteFile("logo_next.png", photoByte, fs.ModePerm)

	if err != nil {
		panic(err)
	}

}
