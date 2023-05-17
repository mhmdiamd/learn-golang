package jsonarray_test

import (
	"encoding/json"
	"fmt"
	"testing"
)

// Di Go-Lang, JSON Array direpresentasikan dalam bentuk slice
// Konversi dari JSON atau ke JSON dilakukan secara otomatis oleh package json menggunakan tipe data slice

type Address struct {
	Street     string
	Country    string
	PostalCode string
}

type Customer struct {
	FirstName  string
	MiddleName string
	LastName   string
	Hobbies    []string
	Addresses  []Address
}

func TestJsonArray(t *testing.T) {
	customer := Customer{
		FirstName:  "Muhamad",
		MiddleName: "Ilham",
		LastName:   "Darmawan",
		Hobbies: []string{
			"Futsal", "Gaming", "Programming",
		},
		Addresses: []Address{
			{
				Street:     "Jl Pondok salak",
				Country:    "Indonesia",
				PostalCode: "15416",
			},
		},
	}

	body, err := json.Marshal(customer)

	if err != nil {
		panic(err)
	}

	fmt.Println(string(body))
}
