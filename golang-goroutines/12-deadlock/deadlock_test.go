package deadlock

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

type UserBalance struct {
	sync.Mutex
	Name    string
	Balance int
}

func (user *UserBalance) Lock() {
	user.Mutex.Lock()
}

func (user *UserBalance) Unlock() {
	user.Mutex.Unlock()
}

func (user *UserBalance) Change(amount int) {
	user.Balance = user.Balance + amount
}

// Mentrasfer Balance dari user1 ke user2
func Transfer(user1 *UserBalance, user2 *UserBalance, amount int) {
	user1.Lock()
	fmt.Println("User ", user1.Name, " Lock")
	user1.Change(-amount)

	user2.Lock()
	fmt.Println("User ", user2.Name, " Lock")
	user2.Change(amount)

	time.Sleep(2 * time.Second)

	user1.Unlock()
	user2.Unlock()
}

func TestDeadlockc(t *testing.T) {
	user1 := UserBalance{
		Name:    "Muhamad Ilham Darmawan",
		Balance: 1000000,
	}

	user2 := UserBalance{
		Name:    "Muhamad Farhan",
		Balance: 1000000,
	}

	go Transfer(&user2, &user1, 100000)
	go Transfer(&user1, &user2, 100000)

	time.Sleep(10 * time.Second)

	fmt.Println("User "+user1.Name+" Balance ", user1.Balance)
	fmt.Println("User "+user2.Name+" Balance ", user2.Balance)

}
