package fileserver_test

import (
	"embed"
	"io/fs"
	"net/http"
	"testing"
)

//go:embed resources
var resources embed.FS

func TestFileServerEmbed(t *testing.T) {

	// Jika kita ingin langsung mengakses file index.js tanpa menggunakan resources, kita bisa menggunakan function fs.Sub() untuk mendapatkan sub directory

	directory, _ := fs.Sub(resources, "resources")
	fileServer := http.FileServer(http.FS(directory))

	mux := http.NewServeMux()
	mux.Handle("/static/", http.StripPrefix("/static", fileServer))

	server := http.Server{
		Addr:    "localhost:8080",
		Handler: mux,
	}

	err := server.ListenAndServe()

	if err != nil {
		panic(err)
	}

}
