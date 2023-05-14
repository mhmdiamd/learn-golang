package synconce

import (
	"fmt"
	"sync"
	"testing"
)

/*
	-	Once adalah fitur di Go-Lang yang bisa kita gunakan untuk memastikan bahsa sebuah function di eksekusi hanya sekali
	-	Jadi berapa banyak pun goroutine yang mengakses, bisa dipastikan bahwa goroutine yang pertama yang bisa mengeksekusi function nya
	-	Goroutine yang lain akan di hiraukan, artinya function tidak akan dieksekusi lagi
*/

var counter = 0

func OnlyOnce() {
	counter++
}

func TestSyncOnce(t *testing.T) {
	once := sync.Once{}
	group := sync.WaitGroup{}

	for i := 1; i <= 100; i++ {
		go func() {
			defer group.Done()

			group.Add(1)
			once.Do(OnlyOnce)
		}()
	}

	group.Wait()
	fmt.Println("Counter : ", counter)

}
