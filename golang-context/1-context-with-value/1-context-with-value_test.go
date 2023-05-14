package contextwithvalue

import (
	"context"
	"fmt"
	"testing"
)

func TestContext(t *testing.T) {
	// Membuat Context

	context := context.Background()
	fmt.Println(context)
}

/*
	-	Saat kita menambah value ke context, secara otomatis akan tercipta child context baru, artinya original context nya tidak akan berubah sama sekali
	-	Untuk membuat menambahkan value ke context, kita bisa menggunakan function context.WithValue(parent, key, value)
*/

func TestContextWithValue(t *testing.T) {
	contextA := context.Background()

	contextB := context.WithValue(contextA, "a", "A")
	contextC := context.WithValue(contextB, "b", "B")

	fmt.Println(contextC.Value("b"))
}
