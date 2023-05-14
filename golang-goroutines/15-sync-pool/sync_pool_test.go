package syncpool

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

/*
Pool adalah implementasi design pattern bernama object pool pattern.
Sederhananya, design pattern Pool ini digunakan untuk menyimpan data, selanjutnya untuk menggunakan datanya, kita bisa mengambil dari Pool, dan setelah selesai menggunakan datanya, kita bisa menyimpan kembali ke Pool nya
Implementasi Pool di Go-Lang ini sudah aman dari problem race condition
*/

func TestPool(t *testing.T) {
	// var pool sync.Pool

	// Memuat default value data pool
	pool := sync.Pool{
		New: func() interface{} {
			return "Default Value"
		},
	}

	// Menambahkan data ke dalam pool
	pool.Put("Muhamad")
	pool.Put("Ilham")
	pool.Put("Darmawan")

	for i := 0; i < 10; i++ {
		go func() {
			data := pool.Get()
			fmt.Println(data)
			time.Sleep(1 * time.Second)
			pool.Put(data)
		}()
	}

	time.Sleep(11 * time.Second)
	fmt.Println("Selesait")

}
