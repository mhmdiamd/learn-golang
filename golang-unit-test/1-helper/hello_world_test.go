package helper

import "testing"

/*
- Nama function untuk unit test harus diawali dengan nama Test
- Misal jika kita ingin mengetest function HelloWorld, maka kita akan membuat function unit test dengan nama TestHelloWorld
*/

func TestHelloWorld(t *testing.T) {
	result := HelloWorld("Ilham")

	if result != "Hello Ilham" {
		panic("Result is not Hello Ilham")
	}
}
