package contextwithdeadline_test

import (
	"context"
	"fmt"
	"runtime"
	"testing"
	"time"
)

/*
	Selain menggunakan timeout untuk melakukan cancel secara otomatis, kita juga bisa menggunakan deadline
	Pengaturan deadline sedikit berbeda dengan timeout, jika timeout kita beri waktu dari sekarang, kalo deadline ditentukan kapan waktu timeout nya, misal jam 12 siang hari ini
	Untuk membuat context dengan cancel signal secara otomatis menggunakan deadline, kita bisa menggunakan function context.WithDeadline(parent, time)
*/

func CreateCounter(ctx context.Context) chan int {
	destination := make(chan int)

	go func() {
		defer close(destination)
		counter := 1
		for {
			select {
			case <-ctx.Done():
				return
			default:
				destination <- counter
				counter++
				time.Sleep(1 * time.Second)
			}
		}
	}()

	return destination
}

func TestContextWithDeadline(t *testing.T) {
	fmt.Println("Jumlah Goroutine : ", runtime.NumGoroutine())

	parent := context.Background()
	child, cancel := context.WithDeadline(parent, time.Now().Add(5*time.Second))
	defer cancel()

	destination := CreateCounter(child)
	for value := range destination {
		fmt.Println("Counter : ", value)
	}

	fmt.Println("Total Goroutine", runtime.NumGoroutine())
}
