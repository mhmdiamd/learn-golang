package queryparameter

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

// Query parameter adalah salah satu fitur yang biasa kita gunakan ketika membuat web
// Query parameter biasanya digunakan untuk mengirim data dari client ke server
// Query parameter ditempatkan pada URL
// Untuk menambahkan query parameter, kita bisa menggunakan ?nama=value pada URL nya

func SayHello(writer http.ResponseWriter, request *http.Request) {

	// Dalam parameter Request, terdapat attribute URL yang berisi data url.URL
	// Dari data URL ini, kita bisa mengambil data query parameter yang dikirim dari client dengan menggunakan method Query() yang akan mengembalikan map

	name := request.URL.Query().Get("name")

	if name == "" {
		fmt.Fprint(writer, "Hello World")
	} else {
		fmt.Fprint(writer, "Hello "+name)
	}
}

func TestQueryParameterTest(t *testing.T) {
	request := httptest.NewRequest("GET", "http://localhost:8080?name=ilham", nil)
	recorder := httptest.NewRecorder()

	SayHello(recorder, request)

	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)

	fmt.Println(string(body))
}
