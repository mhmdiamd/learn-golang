package synccond_test

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

var locker = &sync.Mutex{}
var cond = sync.NewCond(locker)
var group = &sync.WaitGroup{}

func WaitCondition(value int) {
	defer group.Done()
	// Melakukan Lock sama seperti biasanya, hanya saja jiga menggunakan cond maka program yang berada di dalam lock tidak akan langsung di jalankan
	cond.L.Lock()
	// Melakukan waiting atau menunggu perintah untuk mengeksekusi kode yang ada di bawahnya
	cond.Wait()

	fmt.Println("Value : ", value)

	cond.L.Unlock()
}

func TestSyncCond(t *testing.T) {

	for i := 1; i < 10; i++ {
		group.Add(1)
		go WaitCondition(i)
	}

	go func() {
		for i := 1; i < 10; i++ {
			time.Sleep(1 * time.Second)
			cond.Signal() // Memberikan sinyal bahwa goroutines yang ada cond.Wait() sudah bisa di jalankan
		}
	}()

	group.Wait()
}
