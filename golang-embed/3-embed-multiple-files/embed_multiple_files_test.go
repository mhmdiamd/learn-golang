package embedmultiplefiles

import (
	"embed"
	"fmt"
	"testing"
)

// Kadang ada kebutuhan kita ingin melakukan embed beberapa file sekaligus
// Hal ini juga bisa dilakukan menggunakan embed package
// Kita bisa menambahkan komentar //go:embed lebih dari satu baris
// Selain itu variable nya bisa kita gunakan tipe data embed.FS

//go:embed files/a.txt
//go:embed files/b.txt
var files embed.FS

func TestEmbedMultipleFiles(t *testing.T) {
	a, _ := files.ReadFile("files/a.txt")
	fmt.Println(string(a))

	b, _ := files.ReadFile("files/b.txt")
	fmt.Println(string(b))
}
