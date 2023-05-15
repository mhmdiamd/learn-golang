package multiplevalueparameter

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"
)

func MultipleParameter(writer http.ResponseWriter, request *http.Request) {

	var query url.Values = request.URL.Query()
	var names []string = query["name"]
	fmt.Fprint(writer, strings.Join(names, ","))
}

func TestQueryParameterTest(t *testing.T) {

	// Untuk menambahkan query parameter, kita bisa gunakan tanda & lalu diikuti dengan query parameter berikutnya
	request := httptest.NewRequest("GET", "http://localhost:8080?name=ilham&name=darmawan", nil)
	recorder := httptest.NewRecorder()

	MultipleParameter(recorder, request)

	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)

	fmt.Println(string(body))
}
