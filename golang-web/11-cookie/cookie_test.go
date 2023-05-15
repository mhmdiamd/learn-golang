package cookie

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

// Cookie merupakan data yang dibuat di server dan sengaja agar disimpan di web browser
// Untuk membuat cookie di server, kita bisa menggunakan function http.SetCookie()

func SetCookie(writer http.ResponseWriter, request *http.Request) {
	// assign the cookie first
	cookie := new(http.Cookie)
	// assign the cookie name
	cookie.Name = "X-iam-Name"
	// fill cookie valuie
	cookie.Value = request.URL.Query().Get("name")
	cookie.Path = "/"

	// Set the cookie
	http.SetCookie(writer, cookie)
	fmt.Fprint(writer, "Success create cookie")
}

func GetCookie(writer http.ResponseWriter, request *http.Request) {
	cookie, err := request.Cookie("X-iam-Name")

	if err != nil {
		fmt.Fprint(writer, "No Cookie")
	} else {
		fmt.Fprintf(writer, "Hi %s", cookie)
	}
}

func TestCookie(t *testing.T) {
	mux := http.NewServeMux()

	mux.HandleFunc("/set-cookie", SetCookie)
	mux.HandleFunc("/get-cookie", GetCookie)

	server := http.Server{
		Addr:    "localhost:8080",
		Handler: mux,
	}

	err := server.ListenAndServe()

	if err != nil {
		panic(err)
	}
}

func TestCreateCookie(t *testing.T) {
	// Create Request
	request := httptest.NewRequest("GET", "http://localhost:8080/set-cookie?name=ilham", nil)
	// Create Recorder
	recorder := httptest.NewRecorder()

	SetCookie(recorder, request)

	cookies := recorder.Result().Cookies()

	for _, cookie := range cookies {
		fmt.Printf("%s : %s\n", cookie.Name, cookie.Value)

	}
}

func TestGetCookie(t *testing.T) {

	// Create Request
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080/set-cookie?name=ilham", nil)
	cookie := new(http.Cookie)
	cookie.Name = "X-iam-Name"
	cookie.Value = "Ilham"
	request.AddCookie(cookie)

	recorder := httptest.NewRecorder()

	GetCookie(recorder, request)

	response := recorder.Result()

	body, _ := io.ReadAll(response.Body)
	fmt.Println(string(body))
}
