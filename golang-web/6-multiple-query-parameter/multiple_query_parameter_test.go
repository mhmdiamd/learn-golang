package multiplequeryparameter

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func SayHello(writer http.ResponseWriter, request *http.Request) {

	// Dalam parameter Request, terdapat attribute URL yang berisi data url.URL
	// Dari data URL ini, kita bisa mengambil data query parameter yang dikirim dari client dengan menggunakan method Query() yang akan mengembalikan map

	name := request.URL.Query().Get("name")
	age := request.URL.Query().Get("age")

	fmt.Fprint(writer, name)
	fmt.Fprint(writer, age)

}

func TestQueryParameterTest(t *testing.T) {

	// Untuk menambahkan query parameter, kita bisa gunakan tanda & lalu diikuti dengan query parameter berikutnya
	request := httptest.NewRequest("GET", "http://localhost:8080?name=ilham&age=19", nil)
	recorder := httptest.NewRecorder()

	SayHello(recorder, request)

	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)

	fmt.Println(string(body))
}
