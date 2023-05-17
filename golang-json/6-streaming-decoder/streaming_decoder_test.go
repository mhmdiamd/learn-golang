package streamingdecoder

import (
	"encoding/json"
	"fmt"
	"os"
	"testing"
)

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

func TestStreamingDecoder(t *testing.T) {
	reader, _ := os.Open("open.json")
	decoder := json.NewDecoder(reader)

	customer := &Customer{}

	_ = decoder.Decode(customer)

	fmt.Println(customer)
}
