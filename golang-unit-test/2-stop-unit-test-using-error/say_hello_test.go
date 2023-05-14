package stopunittestusingerror

/*
-> Selain Fail() dan FailNow(), ada juga Error() dan Fatal()
		- Error() function lebih seperti melakukan log (print) error, namun setelah melakukan log error, dia akan 	secara otomatis memanggil function Fail(), sehingga mengakibatkan unit test dianggap gagal
-> Namun karena hanya memanggil Fail(), artinya eksekusi unit test akan tetap berjalan sampai selesai
		-Fatal() mirip dengan Error(), hanya saja, setelah melakukan log error, dia akan memanggil FailNow(), sehingga mengakibatkan eksekusi unit test berhenti
*/

import (
	"fmt"
	"testing"
)

// Using Error
// func TestSayHello(t *testing.T) {
// 	result := SayHello("Darmawan")

// 	if result != "Hello Ilham" {
// 		t.Error("Result is not Hello Ilham")
// 	}

// 	fmt.Println("Dieksekusi")
// }

// Using Fatal
func TestSayHello(t *testing.T) {
	result := SayHello("Darmawan")

	if result != "Hello Ilham" {
		t.Fatal("Result is not Hello Ilham")
	}

	fmt.Println("Tidak dieksekusi")
}
