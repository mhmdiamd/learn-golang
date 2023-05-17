package methodnotallowed

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/julienschmidt/httprouter"
	"github.com/stretchr/testify/assert"
)

// Namun pada Router, kita bisa menentukan HTTP Method yang ingin kita gunakan, lantas apa yang terjadi jika client tidak mengirim HTTP Method sesuai dengan yang kita tentukan?
// Maka akan terjadi error Method Not Allowed
// Secara default, jika terjadi error seperti ini, maka Router akan memanggil function http.Error
// Jika kita ingin mengubahnya, kita bisa gunakan router.MethodNotAllowed = http.Handler

func TestMethodNotAllowed(t *testing.T) {
	router := httprouter.New()

	router.MethodNotAllowed = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Not Allowed")
	})

	router.POST("/", func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		fmt.Fprint(w, "POST")
	})

	req := httptest.NewRequest("POST", "http://localhost:8080/", nil)
	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, req)

	response := recorder.Result()

	bodyBytes, _ := io.ReadAll(response.Body)

	assert.Equal(t, "Not Allowed", string(bodyBytes))
}
