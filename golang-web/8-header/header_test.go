package header

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

// Selain Query Parameter, dalam HTTP, ada juga yang bernama Header
// Header adalah informasi tambahan yang biasa dikirim dari client ke server atau sebaliknya
// Jadi dalam Header, tidak hanya ada pada HTTP Request, pada HTTP Response pun kita bisa menambahkan informasi header
// Saat kita menggunakan browser, biasanya secara otomatis header akan ditambahkan oleh browser, seperti informasi browser, jenis tipe content yang dikirim dan diterima oleh browser, dan lain-lain

func RequestHeader(writer http.ResponseWriter, request *http.Request) {
	contentType := request.Header.Get("content-type")
	fmt.Fprint(writer, contentType)
}

func TestRequestHeader(t *testing.T) {
	request := httptest.NewRequest("GET", "http://localhost:8080", nil)
	request.Header.Add("Content-Type", "application/json")
	recorder := httptest.NewRecorder()

	RequestHeader(recorder, request)

	// Get Response Result
	response := recorder.Result()

	// see the body if the value exist
	body, _ := io.ReadAll(response.Body)

	fmt.Println(string(body))
}
