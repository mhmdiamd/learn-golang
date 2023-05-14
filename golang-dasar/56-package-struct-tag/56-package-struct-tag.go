package main

import (
	"fmt"
	"reflect"
)

type Sample struct {
	Name string `required:"true" max:"10"`
}

func isValid(data interface{}) bool {
	t := reflect.TypeOf(data)
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)

		if field.Tag.Get("required") == "true" {
			return reflect.ValueOf(data).Field(i).Interface() != ""
		}
	}
	return true
}

func main() {
	sample := Sample{"Ilham"}
	sampleType := reflect.TypeOf(sample)
	structField := sampleType.Field(0)
	required := structField.Tag.Get("required")

	fmt.Println(required)
	fmt.Println(sampleType.NumField())

	sample.Name = ""
	fmt.Println(isValid(sample))
}
