package channel

/*
	- Untuk mengirim data, kita bisa gunakan kode : channel <- data
	- Sedangkan untuk menerima data, bisa gunakan kode : data <- channel
	- Jika selesai, jangan lupa untuk menutup channel menggunakan function close()
*/

import (
	"fmt"
	"testing"
	"time"
)

func TestChannel(t *testing.T) {
	channel := make(chan string)

	go func() {
		time.Sleep(2 * time.Second)
		channel <- "Muhamad Ilham Darmawan"
	}()

	data := <-channel
	fmt.Println(data)
	close(channel)

}
