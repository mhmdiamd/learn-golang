package timeticker_test

import (
	"fmt"
	"testing"
	"time"
)

/*
	- Ticker adalah representasi kejadian yang berulang
	- Ketika waktu ticker sudah expire, maka event akan dikirim ke dalam channel
	- Untuk membuat ticker, kita bisa menggunakan time.NewTicker(duration)
	- Untuk menghentikan ticker, kita bisa menggunakan Ticker.Stop()
*/

func TestTimeTicker(t *testing.T) {
	ticker := time.NewTicker(1 * time.Second)

	for time := range ticker.C {
		fmt.Println(time.Second())
	}
}

/*
	- Kadang kita tidak butuh data Ticker nya, kita hanya butuh channel nya saja
	- Jika demikian, kita bisa menggunakan function timer.Tick(duration), function ini tidak akan mengembalikan Ticker, hanya mengembalikan channel timer nya saja

*/

func TestTimeTickerTick(t *testing.T) {
	var timer = time.Tick(2 * time.Second)

	for tick := range timer {
		fmt.Println(tick)
	}
}
