package panichandler

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/julienschmidt/httprouter"
	"github.com/stretchr/testify/assert"
)

// di Router, sudah disediakan untuk menangani panic, caranya dengan menggunakan attribute PanicHandler : func(http.ResponseWriter, *http.Request, interface{})

func TestPanicHandler(t *testing.T) {
	router := httprouter.New()
	router.PanicHandler = func(w http.ResponseWriter, r *http.Request, i interface{}) {
		fmt.Fprint(w, "Panic : ", i)
	}

	router.GET("/", func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		panic("Ups")
	})

	req := httptest.NewRequest("GET", "http://localhost:8080/", nil)
	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, req)

	response := recorder.Result()

	bodyBytes, _ := io.ReadAll(response.Body)

	fmt.Println(string(bodyBytes))
	assert.Equal(t, "Panic : Ups", string(bodyBytes))

}
