package encodejson

import (
	"encoding/json"
	"fmt"
	"testing"
)

func logJson(data interface{}) {
	body, err := json.Marshal(data)

	if err != nil {
		panic(err)
	}

	fmt.Println(string(body))
}

func TestEncodeJson(t *testing.T) {
	logJson("ilham")
	logJson(1)
	logJson(true)
	logJson([]string{"Muhamad", "Ilham", "Darmawan"})
}
