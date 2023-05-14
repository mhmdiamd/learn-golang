package contextwithtimeout_test

import (
	"context"
	"fmt"
	"runtime"
	"testing"
	"time"
)

/*
	Selain menambahkan value ke context, dan juga sinyal cancel, kita juga bisa menambahkan sinyal cancel ke context secara otomatis dengan menggunakan pengaturan timeout
	Dengan menggunakan pengaturan timeout, kita tidak perlu melakukan eksekusi cancel secara manual, cancel akan otomatis di eksekusi jika waktu timeout sudah terlewati
	Penggunaan context dengan timeout sangat cocok ketika misal kita melakukan query ke database atau http api, namun ingin menentukan batas maksimal timeout nya
	Untuk membuat context dengan cancel signal secara otomatis menggunakan timeout, kita bisa menggunakan function context.WithTimeout(parent, duration)
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

func TestContextWithTimeout(t *testing.T) {
	fmt.Println("Jumlah Goroutine : ", runtime.NumGoroutine())

	parent := context.Background()
	child, cancel := context.WithTimeout(parent, 5*time.Second)
	defer cancel()

	destination := CreateCounter(child)
	for value := range destination {
		fmt.Println("Counter : ", value)
	}

	fmt.Println("Total Goroutine", runtime.NumGoroutine())
}
