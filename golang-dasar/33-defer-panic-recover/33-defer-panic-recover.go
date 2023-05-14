package main

import "fmt"

// Defer function adalah function yang bisa kita jadwalkan untuk dieksekusi setelah sebuah function selesai di eksekusi

/*
-	Panic function adalah function yang bisa kita gunakan untuk menghentikan program
- Panic function biasanya dipanggil ketika terjadi error pada saat program kita berjalan
- Saat panic function dipanggil, program akan terhenti, namun defer function tetap akan dieksekusi

- Panic sama seperti throw exception pada javascript yang mana jika dia panic maka kode dibawahnya tidak di esksekusi
*/

/*
- Recover adalah function yang bisa kita gunakan untuk menangkap data panic
- Dengan recover proses panic akan terhenti, sehingga program akan tetap berjalan
*/

func endApp() {

	// Menangkap panic message dengan recover dan menjalankan ulang aplikasinya
	message := recover()
	fmt.Println("Error message : ", message)
	fmt.Println("Aplikasi Selesai!")
}

func runApp(isPanic bool) {
	defer endApp()

	if isPanic == true {
		panic("Aplikasi Error")
	}
}

func main() {
	runApp(true)
}
