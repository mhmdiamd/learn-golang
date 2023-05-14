package buffredchannel

import (
	"fmt"
	"testing"
)

/*
	- Kita bebas memasukkan berapa jumlah kapasitas antrian di dalam buffer
	- Jika kita set misal 5, artinya kita bisa menerima 5 data di buffer.
	- Jika kita mengirim data ke 6, maka kita diminta untuk menunggu sampai buffer ada yang kosong
	- Ini cocok sekali ketika memang goroutine yang menerima data lebih lambat dari yang mengirim data
*/

func TestBuffredChannel(t *testing.T) {

	channel := make(chan string, 3)

	channel <- "Muhamad ilham Darmawan"
	channel <- "Muhamad ilham Darmawan2"

	fmt.Println(cap(channel)) // melihat Panjang Buffer
	fmt.Println(len(channel)) // Melihat banyak data buffer yang ada di dalam cahnnel

	close(channel)
}
