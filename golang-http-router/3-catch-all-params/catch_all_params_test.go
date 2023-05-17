package catchallparams

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/julienschmidt/httprouter"
	"github.com/stretchr/testify/assert"
)

func TestCatchAllParams(t *testing.T) {
	router := httprouter.New()
	router.GET("/images/*image", func(writer http.ResponseWriter, req *http.Request, params httprouter.Params) {
		text := "Image : " + params.ByName("image")
		fmt.Fprint(writer, text)
	})

	request := httptest.NewRequest("GET", "http://localhost:8080/images/small/profile.png", nil)
	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)
	response := recorder.Result()

	bodyBytes, _ := io.ReadAll(response.Body)
	fmt.Println(string(bodyBytes))
	assert.Equal(t, "Image : /small/profile.png", string(bodyBytes))
}
