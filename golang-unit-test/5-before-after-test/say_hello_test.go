package beforeaftertest

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestMain(m *testing.M) {
	fmt.Println("Sebelum unitest")

	m.Run() // Rrunn semua Testing

	fmt.Println("End semua testing")
}

func TestSayHello(t *testing.T) {
	result := SayHello("Ilham")

	require.Equal(t, "Hello Ilham", result)
}
