package gomaxprocs_test

import (
	"fmt"
	"runtime"
	"sync"
	"testing"
	"time"
)

/*
	-	Untuk mengetahui berapa jumlah Thread, kita bisa menggunakan GOMAXPROCS, yaitu sebuah function di package runtime yang bisa kita gunakan untuk mengubah jumlah thread atau mengambil jumlah thread
	-	Secara default, jumlah thread di Go-Lang itu sebanyak jumlah CPU di komputer kita.
	-	Kita juga bisa melihat berapa jumlah CPU kita dengan menggunakan function runtime.NumCpu()
*/

func TestGomaxprocs(t *testing.T) {

	group := &sync.WaitGroup{}

	for i := 0; i < 100; i++ {
		group.Add(1)
		go func() {
			time.Sleep(3 * time.Second)
			group.Done()
		}()
	}

	totalCpu := runtime.NumCPU()
	fmt.Println("Total Cpu : ", totalCpu)

	runtime.GOMAXPROCS(20)
	totalThread := runtime.GOMAXPROCS(-1)
	fmt.Println("Total Thread : ", totalThread)

	totalGoroutine := runtime.NumGoroutine()
	fmt.Println("Total Goroutine : ", totalGoroutine)

	group.Wait()
}
