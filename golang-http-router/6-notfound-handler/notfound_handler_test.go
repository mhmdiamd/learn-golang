package notfoundhandler

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/julienschmidt/httprouter"
	"github.com/stretchr/testify/assert"
)

// Secara default, jika tidak ada route tidak ditemukan, Router akan melanjutkan request ke http.NotFound, namun kita bisa mengubah nya
// Caranya dengan mengubah router.NotFound = http.Handler

func TestNotfoundHandler(t *testing.T) {
	router := httprouter.New()
	router.NotFound = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Gak Ketemu")
	})

	req := httptest.NewRequest("GET", "http://localhost:8080/notfound", nil)
	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, req)

	response := recorder.Result()

	bodyBytes, _ := io.ReadAll(response.Body)

	fmt.Println(string(bodyBytes))
	assert.Equal(t, "Gak Ketemu", string(bodyBytes))
}
