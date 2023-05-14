package subtest

// Sub test adalah test2 kecil yang terdapat di dalam Main Test yang dijalankan di dalam parameter t.Run()

/*
	- Kita sudah tahu jika ingin menjalankan sebuah unit test function, kita bisa gunakan perintah :
	go test -run TestNamaFunction
	- Jika kita ingin menjalankan hanya salah satu sub test, kita bisa gunakan perintah :
	go test -run TestNamaFunction/NamaSubTest
	-	Atau untuk semua test semua sub test di semua function, kita bisa gunakan perintah :
	go test -run /NamaSubTest
*/

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSayHello(t *testing.T) {
	t.Run("Ilham Test", func(t *testing.T) {
		result := SayHello("Ilham")
		assert.Equal(t, "Hello Ilham", result)
	})

	t.Run("Not Ilham Test", func(t *testing.T) {
		result := SayHello("Not Ilham")
		assert.Equal(t, "Hello Ilham", result)
	})
}
