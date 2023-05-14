package main

import (
	"container/ring"
	"fmt"
	"strconv"
)

/*
	Package container/ring adalah implementasi struktur data circular list
	Circular list adalah struktur data ring, dimana diakhir element akan kembali ke element awal (HEAD)
*/

func main() {
	var data *ring.Ring = ring.New(5)

	// Looping data ring
	for i := 0; i < data.Len(); i++ {
		data.Value = "Value " + strconv.FormatInt(int64(i), 10)
		data = data.Next()
	}

	data.Do(func(i interface{}) {
		fmt.Println(i)
	})

}
