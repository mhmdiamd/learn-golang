package jsonobject_test

import (
	"encoding/json"
	"fmt"
	"testing"
)

// JSON Object di Go-Lang direpresentasikan dengan tipe data Struct
// Dimana tiap attribute di JSON Object merupakan attribute di Struct

type Customer struct {
	FirstName, MiddleName, LastName string
}

func TestObjectJson(t *testing.T) {
	customer := Customer{
		FirstName:  "Muhamad",
		MiddleName: "Ilham",
		LastName:   "Darmawan",
	}

	body, err := json.Marshal(customer)

	if err != nil {
		panic(err)
	}

	fmt.Println(string(body))
}
