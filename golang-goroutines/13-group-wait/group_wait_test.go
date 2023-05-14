package groupwait

/*
	-WaitGroup adalah fitur yang bisa digunakan untuk menunggu sebuah proses selesai dilakukan
	-	Hal ini kadang diperlukan, misal kita ingin menjalankan beberapa proses menggunakan goroutine, tapi kita ingin semua proses selesai terlebih dahulu sebelum aplikasi kita selesai
	- Kasus seperti ini bisa menggunakan WaitGroup
	-	Untuk menandai bahwa ada proses goroutine, kita bisa menggunakan method Add(int), setelah proses goroutine selesai, kita bisa gunakan method Done()
	-	Untuk menunggu semua proses selesai, kita bisa menggunakan method Wait()

*/

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func RunAsynchronous(group *sync.WaitGroup, index int) {
	defer group.Done()

	group.Add(1)
	fmt.Println("Hello ", index)
	time.Sleep(1 * time.Second)
}

func TestGroupWait(t *testing.T) {
	group := &sync.WaitGroup{}

	for i := 1; i <= 100; i++ {
		go RunAsynchronous(group, i)
	}

	group.Wait()
	fmt.Println("Program is Done!")
}
