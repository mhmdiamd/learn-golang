package manygoroutinetest

import (
	"fmt"
	"testing"
	"time"
)

func DisplayNumber(number int) {
	fmt.Println("Display ", number)
}

func TestDisplayNumber(t *testing.T) {
	for i := 1; i < 100000; i++ {
		go DisplayNumber(i)
	}

	time.Sleep(1 * time.Second)
}
