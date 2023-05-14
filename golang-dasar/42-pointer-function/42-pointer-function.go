package main

import "fmt"

type Address struct {
	city, province, country string
}

// Menggunakan * untuk mengubah parameter menjadi pointer, yang man jika  ada perubahan pada parameter address di function tersebut makan akan mempengaruhi data asli yang menjadi parameternya
func sayHelloToPointer(address *Address) {
	address.city = "Tangerang Selatan"
}

func main() {
	address := Address{
		city:     "Bandung",
		province: "Jawa Barat",
		country:  "Indonesia",
	}

	// cara ke 1
	// sayHelloToPointer(&address)

	// Data city pada address akan berubah jadi tangerang selatan!

	// cara ke 2
	var addressPointer *Address = &address
	sayHelloToPointer(addressPointer)

	fmt.Println(address)

}
