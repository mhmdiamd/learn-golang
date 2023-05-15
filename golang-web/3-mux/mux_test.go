package mux

import (
	"fmt"
	"net/http"
	"testing"
)

// Saat membuat web, kita biasanya ingin membuat banyak sekali endpoint URL
// HandlerFunc sayangnya tidak mendukung itu
// Alternative implementasi dari Handler adalah ServeMux
// ServeMux adalah implementasi Handler yang bisa mendukung multiple endpoint

func TestMux(t *testing.T) {

	mux := http.NewServeMux()

	mux.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprint(writer, "Hello World!")
	})

	mux.HandleFunc("/ilham", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprint(writer, "Hello Ilham!")
	})

	// Url Patter Example
	mux.HandleFunc("/ilham/darmawan", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprint(writer, "Hello Ilham Darmawan")
	})

	server := http.Server{
		Addr:    "localhost:8080",
		Handler: mux,
	}

	err := server.ListenAndServe()

	if err != nil {
		panic(err)
	}
}
