package decodejson

import (
	"encoding/json"
	"fmt"
	"testing"
)

// Untuk melakukan konversi dari JSON ke tipe data di Go-Lang (Decode), kita bisa menggunakan function json.Unmarshal(byte[], interface{})
// Dimana byte[] adalah data JSON nya, sedangkan interface{} adalah tempat menyimpan hasil konversi, biasa berupa pointer

type Customer struct {
	FirstName, MiddleName, LastName string
}

func TestDecodedJson(t *testing.T) {
	jsonRequest := `{"FirstName" : "Muhamad", "MiddleName" : "Ilham", "LastName" : "Darmawan"}`

	// melakukan konversi dari json string ke slice byte
	jsonBytes := []byte(jsonRequest)

	customer := &Customer{}
	json.Unmarshal(jsonBytes, customer)

	fmt.Println(customer)
}
