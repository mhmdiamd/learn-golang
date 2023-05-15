package header

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

// Selain Query Parameter, dalam HTTP, ada juga yang bernama Header
// Header adalah informasi tambahan yang biasa dikirim dari client ke server atau sebaliknya
// Jadi dalam Header, tidak hanya ada pada HTTP Request, pada HTTP Response pun kita bisa menambahkan informasi header
// Saat kita menggunakan browser, biasanya secara otomatis header akan ditambahkan oleh browser, seperti informasi browser, jenis tipe content yang dikirim dan diterima oleh browser, dan lain-lain

func ResponseHeader(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Add("X-Powered-By", "Muhamad Ilham Darmawan")
	fmt.Fprint(writer, "ok")
}

func TestResponseHeader(t *testing.T) {
	request := httptest.NewRequest("GET", "http://localhost:8080", nil)
	request.Header.Add("Content-Type", "application/json")
	recorder := httptest.NewRecorder()

	ResponseHeader(recorder, request)

	poweredBy := recorder.Header().Get("x-powered-by")

	fmt.Println(poweredBy)
}
