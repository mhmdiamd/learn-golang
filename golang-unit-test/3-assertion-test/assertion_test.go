package assertion

// Assert berfungsi jika ada error testing maka program akan tetap di jalankan
// Require berfungsi jika ada error testing maka program akan diberhentikan dan kode dibawahnya tidak dijalankan

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestSayWorldAssertion(t *testing.T) {
	result := SayWorldAssertion("Ilham")

	// assert.Equal(t, "Hello Ilham", result)
	// fmt.Println("Dieksekusi")

	require.Equal(t, "Hello Ilham", result)
	fmt.Println("tidak dieksekusi")
}
