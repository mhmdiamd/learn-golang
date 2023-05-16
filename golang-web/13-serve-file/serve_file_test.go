package servefile_test

import (
	"net/http"
	"testing"
)

// Kadang ada kasus misal kita hanya ingin menggunakan static file sesuai dengan yang kita inginkan
// Hal ini bisa dilakukan menggunakan function http.ServeFile()
// Dengan menggunakan function ini, kita bisa menentukan file mana yang ingin kita tulis ke http response

func ServeFile(writer http.ResponseWriter, request *http.Request) {
	if request.URL.Query().Get("name") != "" {
		http.ServeFile(writer, request, "./resources/index.html")
	} else {
		http.ServeFile(writer, request, "./resources/notfound.html")
	}
}

func TestServeFile(t *testing.T) {
	server := http.Server{
		Addr:    "localhost:8080",
		Handler: http.HandlerFunc(ServeFile),
	}

	err := server.ListenAndServe()

	if err != nil {
		panic(err)
	}
}
