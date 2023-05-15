package patchmatcher

import (
	"embed"
	"fmt"
	"testing"
)

// Selain manual satu per satu, kita bisa mengguakan patch matcher untuk membaca multiple file yang kita inginkan
// Ini sangat cocok ketika misal kita punya pola jenis file yang kita inginkan untuk kita baca
// Caranya, kita perlu menggunakan path matcher seperti pada package function path.Match

// go:embed files/*.txt
var path embed.FS

func TestPatchMatcher(t *testing.T) {
	// Get Directory files
	dir, _ := path.ReadDir("files")

	for _, entry := range dir {
		if !entry.IsDir() {
			fmt.Println(entry.Name())
			content, _ := path.ReadFile("files/" + entry.Name())
			fmt.Println("Content : " + string(content))
		}
	}

}
