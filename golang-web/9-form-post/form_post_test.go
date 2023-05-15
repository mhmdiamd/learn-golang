package formpost

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

// Semua data form post yang dikirim dari client, secara otomatis akan disimpan dalam attribute Request.PostForm
// Namun sebelum kita bisa mengambil data di attribute PostForm, kita wajib memanggil method Request.ParseForm() terlebih dahulu, method ini digunakan untuk melakukan parsing data body apakah bisa di parsing menjadi form data atau tidak, jika tidak bisa di parsing, maka akan menyebabkan error

func FormPost(writer http.ResponseWriter, request *http.Request) {

	// cek error in formPost
	err := request.ParseForm()

	if err != nil {
		panic(err)
	}

	firstName := request.PostForm.Get("firstName")
	lastName := request.PostForm.Get("lastName")

	fmt.Fprintf(writer, "%s %s", firstName, lastName)
}

func TestPostForm(t *testing.T) {
	// Create Request Body First
	requestBody := strings.NewReader("firstName=Muhamad&lastName=Darmawan")

	// Send Request
	request := httptest.NewRequest("POST", "http://localhost/", requestBody)
	// Set up the header if you want to send as form data
	request.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	// Create Recorder
	recorder := httptest.NewRecorder()
	FormPost(recorder, request)

	// Get response Request
	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)
	fmt.Println(string(body))
}
