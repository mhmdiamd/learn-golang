package fileserver_test

import (
	"net/http"
	"testing"
)

// Dengan ini, kita bisa membuat Handler di Go-Lang web yang digunakan sebagai static file server
// Dengan menggunakan FileServer, kita tidak perlu manual me-load file lagi
// FileServer adalah Handler, jadi bisa kita tambahka ke dalam http.Server atau http.ServeMux

func TestFileServer(t *testing.T) {
	directory := http.Dir("./resources")
	fileServer := http.FileServer(directory)

	mux := http.NewServeMux()

	// Jika kita coba jalankan, saat kita membuka misal /static/index.js, maka akan dapat error 404 Not Found
	// mux.Handle("/static/", fileServer)

	mux.Handle("/static/", http.StripPrefix("/static", fileServer))

	server := http.Server{
		Addr:    "localhost:8080",
		Handler: mux,
	}

	err := server.ListenAndServe()

	// Hal ini dikarenakan FileServer akan membaca url, lalu mencari file berdasarkan url nya, jadi jika kita membuat /static/index.js, maka FileServer akan mencari ke file /resources/static/index.js
	// Hal ini menyebabkan 404 Not Found karena memang file nya tidak bisa ditemukan
	// Oleh karena itu, kita bisa menggunakan function http.StripPrefix() untuk menghapus prefix di url

	if err != nil {
		panic(err)
	}
}
