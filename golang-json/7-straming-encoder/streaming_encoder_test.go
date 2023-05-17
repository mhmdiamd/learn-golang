package stramingencoder

import (
	"encoding/json"
	"fmt"
	"os"
	"testing"
)

// Selain decoder, package json juga mendukung membuat Encoder yang bisa digunakan untuk menulis langsung JSON nya ke io.Writer
// Dengan begitu, kita tidak perlu menyimpan JSON datanya terlebih dahulu ke dalam variable string atau []byte, kita bisa langsung tulis ke io.Writer

// Untuk membuat Encoder, kita bisa menggunakan function json.NewEncoder(writer)
// Dan untuk menulis data sebagai JSON langsung ke writer, kita bisa gunakan function Encode(interface{})

type Customer struct {
	FirstName, MiddleName, LastName string
}

func TestStreamingEncoder(t *testing.T) {
	writer, _ := os.Create("sample_output.json")
	encoder := json.NewEncoder(writer)

	customer := &Customer{
		FirstName:  "Muhamad",
		MiddleName: "Ilham",
		LastName:   "Darmawan",
	}

	_ = encoder.Encode(customer)

	fmt.Println(customer)
}
