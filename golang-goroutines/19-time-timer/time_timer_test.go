package timetimer_test

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

// Timer adalah representasi satu kejadian
// Ketika waktu timer sudah expire, maka event akan dikirim ke dalam channel
// Untuk membuat Timer kita bisa menggunakan time.NewTimer(duration)

func TestTimeTimer(t *testing.T) {
	timer := time.NewTimer(5 * time.Second)
	fmt.Println(time.Now())

	time := timer.C
	fmt.Println(time)
}

// Kadang kita hanya butuh channel nya saja, tidak membutuhkan data Timer nya
// Untuk melakukan hal itu kita bisa menggunakan function time.After(duration)

func TestTimeAfter(t *testing.T) {
	channel := time.After(5 * time.Second)
	fmt.Println(time.Now())

	tick := <-channel
	fmt.Println(tick)
}

/*
	-	Kadang ada kebutuhan kita ingin menjalankan sebuah function dengan delay waktu tertentu
	-	Kita bisa memanfaatkan Timer dengan menggunakan function time.AfterFunc()
	-	Kita tidak perlu lagi menggunakan channel nya, cukup kirim kan function yang akan dipanggil ketika Timer mengirim kejadiannya

*/

func TestTimerAfterFunc(t *testing.T) {
	group := &sync.WaitGroup{}
	group.Add(1)

	time.AfterFunc(1*time.Second, func() {
		defer group.Done()

		fmt.Println("execute after 1 second")
	})

	group.Wait()
}
