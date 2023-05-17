package jsonmap

import (
	"encoding/json"
	"fmt"
	"testing"
)

// Saat menggunakan JSON, kadang mungkin kita menemukan kasus data JSON nya dynamic
// Artinya atribut nya tidak menentu, bisa bertambah, bisa berkurang, dan tidak tetap
// Pada kasus seperti itu, menggunakan Struct akan menyulitkan, karena pada Struct, kita harus menentukan semua atribut nya
// Untuk kasus seperti ini, kita bisa menggunakan tipe data map[string]interface{}
// Secara otomatis, atribut akan menjadi key di map, dan value menjadi value di map
// Namun karena value berupa interface{}, maka kita harus lakukan konversi secara manual jika ingin mengambil value nya
// Dan tipe data Map tidak mendukung JSON Tag lagi

func TestJsonMap(t *testing.T) {
	jsonRequest := `{"id" : "1", "name" : "apple/iphone", "price" : "200000"}`
	jsonBytes := []byte(jsonRequest)

	var result map[string]interface{}
	_ = json.Unmarshal(jsonBytes, &result)

	fmt.Println(result)
	fmt.Println(result["id"])
	fmt.Println(result["name"])
	fmt.Println(result["price"])
}
