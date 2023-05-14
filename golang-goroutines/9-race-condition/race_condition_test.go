package racecondition

import (
	"fmt"
	"testing"
	"time"
)

func TestRaceCondition(t *testing.T) {
	var x = 0

	for i := 1; i <= 1000; i++ {
		go func() {
			for j := 1; j <= 100; j++ {
				x = x + 1
			}
		}()
	}

	time.Sleep(5 * time.Second)
	// result counter seharusnya adalah 100000 tapi karna terjadi race condition maka ada eksekusi paralel yang nabrak yang mana menyebabkan salah satu ekseskusinya gagal
	fmt.Println("Counter : ", x)
}
