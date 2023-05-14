package channelparameter

import (
	"fmt"
	"testing"
	"time"
)

/*
	- Dalam kenyataan pembuatan aplikasi, seringnya kita akan mengirim channel ke function lain via parameter
	- Sebelumnya kita tahu bahkan di Go-Lang by default, parameter adalah pass by value, artinya value akan diduplikasi lalu dikirim ke function parameter, sehingga jika kita ingin mengirim data asli, kita biasa gunakan pointer (agar pass by reference).
	- Berbeda dengan Channel, kita tidak perlu melakukan hal tersebut

*/

func GiveMeResponse(channel chan string) {
	time.Sleep(2 * time.Second)
	channel <- "Muhamad Ilham Darmawan"
}

func TestGiveMeResponse(t *testing.T) {
	channel := make(chan string)

	go GiveMeResponse(channel)

	data := <-channel
	fmt.Println(data)
	close(channel)
}
