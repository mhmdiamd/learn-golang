package router

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/julienschmidt/httprouter"
	"github.com/stretchr/testify/assert"
)

// Inti dari library HttpRouter adalah struct Router
// Router ini merupakan implementasi dari http.Handler, sehingga kita bisa dengan mudah menambahkan ke dalam http.Server
// Untuk membuat Router, kita bisa menggunakan function httprouter.New(), yang akan mengembalikan Router pointer

func TestRouter(t *testing.T) {
	router := httprouter.New()
	router.GET("/", func(writer http.ResponseWriter, req *http.Request, params httprouter.Params) {
		fmt.Fprint(writer, "Hello Get")
	})

	request := httptest.NewRequest("GET", "http://localhost:8080/", nil)
	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	response := recorder.Result()

	body, _ := io.ReadAll(response.Body)

	assert.Equal(t, "Hello Get", string(body))
}
