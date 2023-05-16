package servefile_test

import (
	"embed"
	"fmt"
	"net/http"
	"testing"
)

// Parameter function http.ServeFile hanya berisi string file name, sehingga tidak bisa menggunakan Go-Lang Embed
// Namun bukan berarti kita tidak bisa menggunakan Go-Lang embed, karena jika untuk melakukan load file, kita hanya butuh menggunakan package fmt dan ResponseWriter saja

// go:embed resources/index.html
var resourceOk embed.FS

//go:embed resources/notfound.html
var resrouceNotFound embed.FS

func ServeFileEmbed(writer http.ResponseWriter, request *http.Request) {
	if request.URL.Query().Get("name") != "" {
		fmt.Fprint(writer, resourceOk)
	} else {
		fmt.Fprint(writer, resrouceNotFound)
	}
}

func TestServeFileEmbed(t *testing.T) {
	server := http.Server{
		Addr:    "localhost:8080",
		Handler: http.HandlerFunc(ServeFile),
	}

	err := server.ListenAndServe()

	if err != nil {
		panic(err)
	}
}
