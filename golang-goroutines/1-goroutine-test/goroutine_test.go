package goroutinetest

import (
	"fmt"
	"testing"
	"time"
)

func RunHelloWorld() {
	fmt.Println("Hello World")
}

func TestCreateGoroutine(t *testing.T) {
	go RunHelloWorld()
	fmt.Println("Ups...")

	// Time sleep berfungsi untuk bengong dlu terminalnya
	time.Sleep(1 * time.Second)
}
