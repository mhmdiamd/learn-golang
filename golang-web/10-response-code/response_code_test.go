package responsecode

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

// Jika kita ingin mengubahnya, kita bisa menggunakan function ResponseWriter.WriteHeader(int)
// Semua data status code juga sudah disediakan di Go-Lang, jadi kita ingin, kita bisa gunakan variable yang sudah disediakan

func ResponseCode(writer http.ResponseWriter, request *http.Request) {
	name := request.URL.Query().Get("name")

	if name == "" {
		writer.WriteHeader(404)
		fmt.Fprint(writer, "query not found")
	} else {
		writer.WriteHeader(200)
		fmt.Fprintf(writer, "Hi %s", name)
	}
}

func TestResponseCode(t *testing.T) {
	// Create request First
	request := httptest.NewRequest("GET", "http://localhost:8080", nil)

	// Create Recorder
	recorder := httptest.NewRecorder()

	ResponseCode(recorder, request)

	// Get Response
	response := recorder.Result()

	body, _ := io.ReadAll(response.Body)
	fmt.Println(string(body))
	fmt.Println(response.StatusCode) // return will be 404
	fmt.Println(response.Status)     // Return will be 404 name not found
	fmt.Println(string(body))
}
