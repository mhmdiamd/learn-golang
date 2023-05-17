package params

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/julienschmidt/httprouter"
	"github.com/stretchr/testify/assert"
)

// Kadang kita butuh membuat URL yang tidak fix, alias bisa berubah-ubah, misal /products/1, /products/2, dan seterusnya
// ServeMux tidak mendukung hal tersebut, namun Router mendukung hal tersebut
// Parameter yang dinamis yang terdapat di URL, secara otomatis dikumpulkan di Params

func TestParams(t *testing.T) {
	router := httprouter.New()
	router.GET("/products/:id", func(writer http.ResponseWriter, req *http.Request, params httprouter.Params) {
		text := "Products " + params.ByName("id")
		fmt.Fprint(writer, text)
	})

	request := httptest.NewRequest("GET", "http://localhost:8080/products/1", nil)
	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)
	response := recorder.Result()

	bodyBytes, _ := io.ReadAll(response.Body)
	fmt.Println(string(bodyBytes))
	assert.Equal(t, "Products 1", string(bodyBytes))
}
