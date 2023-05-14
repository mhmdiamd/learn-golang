package syncmap

import (
	"fmt"
	"sync"
	"testing"
)

/*
- Map ini mirip Go-Lang map, namun yang membedakan, Map ini aman untuk menggunaan concurrent menggunakan goroutine
- Ada beberapa function yang bisa kita gunakan di Map :
		- Store(key, value) untuk menyimpan data ke Map
		- Load(key) untuk mengambil data dari Map menggunakan key
		- Delete(key) untuk menghapus data di Map menggunakan key
		-	Range(function(key, value)) digunakan untuk melakukan iterasi seluruh data di Map
*/

func AddData(data *sync.Map, value int, group *sync.WaitGroup) {
	defer group.Done()

	group.Add(1)
	data.Store(value, value)
}

func TestSyncMap(t *testing.T) {
	data := &sync.Map{}
	group := sync.WaitGroup{}

	for i := 0; i < 100; i++ {
		go AddData(data, i, &group)
	}

	group.Wait()

	data.Range(func(key, value interface{}) bool {
		fmt.Println(key, " : ", value)
		return true
	})

}
