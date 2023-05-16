package templatedata

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
	"text/template"
)

type Page struct {
	Title string
	Name  string
}

// kita perlu memberi tahu Field atau Key mana yang akan kita gunakan untuk mengisi data dinamis di template
// Kita bisa menyebutkan dengan cara seperti ini {{.NamaField}}

func TemplateDataStruct(writer http.ResponseWriter, req *http.Request) {
	// Create template first
	t := template.Must(template.ParseFiles("./template/index.gohtml"))

	// Execute file and send content to template engine menggunakan data map
	// t.ExecuteTemplate(writer, "index.gohtml", map[string]interface{}{
	// 	"Title": "Ilham Title",
	// 	"Name":  "Muahamad Ilham Darmawan",
	// })

	t.ExecuteTemplate(writer, "index.gohtml", Page{
		Title: "Ilham Title",
		Name:  "Muahamad Ilham Darmawan",
	})
}

func TestTemplateDataStruct(t *testing.T) {
	// Create request
	req := httptest.NewRequest("GET", "http://localhost/", nil)
	recorder := httptest.NewRecorder()

	TemplateDataStruct(recorder, req)

	response := recorder.Result()

	body, _ := io.ReadAll(response.Body)

	fmt.Println(string(body))
}

func TestTempateDataServer(t *testing.T) {

	server := http.Server{
		Addr: "localhost:8080",
	}

	err := server.ListenAndServe()

	if err != nil {
		panic(err)
	}
}
