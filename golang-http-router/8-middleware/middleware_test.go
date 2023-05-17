package middleware

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/julienschmidt/httprouter"
	"github.com/stretchr/testify/assert"
)

type logMiddleware struct {
	http.Handler
}

func (middleware *logMiddleware) ServeHttp(w http.ResponseWriter, req *http.Request) {
	fmt.Println("Recive Request")
	middleware.Handler.ServeHTTP(w, req)
}

func TestMiddleware(t *testing.T) {
	router := httprouter.New()
	router.GET("/", func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		fmt.Fprint(w, "Middleware")
	})

	middleware := logMiddleware{
		Handler: router,
	}

	req := httptest.NewRequest("GET", "http://localhost:8080/", nil)
	recorder := httptest.NewRecorder()

	middleware.ServeHttp(recorder, req)

	response := recorder.Result()

	bodyBytes, _ := io.ReadAll(response.Body)

	assert.Equal(t, "Middleware", string(bodyBytes))
}
