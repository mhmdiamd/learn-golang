package atomic_test

import (
	"fmt"
	"sync"
	"sync/atomic"
	"testing"
)

/*
	- Atomic merupakan package yang digunakan untuk menggunakan data primitive secara aman pada proses concurrent
	- Contohnya sebelumnya kita telah menggunakan Mutex untuk melakukan locking ketika ingin menaikkan angka di counter. 	Hal ini sebenarnya bisa digunakan menggunakan Atomic package
*/

// Contoh Menggunakan Atomic

func TestAtomic(t *testing.T) {

	var x int64 = 0
	var group = &sync.WaitGroup{}

	for i := 1; i <= 1000; i++ {
		group.Add(1)
		go func() {
			for j := 1; j <= 100; j++ {
				atomic.AddInt64(&x, 1)
			}
			group.Done()
		}()
	}

	group.Wait()
	fmt.Println("Total : ", x)
}
