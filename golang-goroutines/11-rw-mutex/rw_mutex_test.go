package rwmutex

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

/*
	- Kadang ada kasus dimana kita ingin melakukan locking tidak hanya pada proses mengubah data, tapi juga membaca data
	- Kita sebenarnya bisa menggunakan Mutex saja, namun masalahnya nanti akan rebutan antara proses membaca dan mengubah
	- Di Go-Lang telah disediakan struct RWMutex (Read Write Mutex) untuk menangani hal ini, dimana Mutex jenis ini memiliki dua lock, lock untuk Read dan lock untuk Write

*/

type BankingAccount struct {
	RWMutext sync.RWMutex
	Balance  int
}

func (account *BankingAccount) AddBalance(amount int) {
	account.RWMutext.Lock()
	account.Balance = account.Balance + amount
	account.RWMutext.Unlock()
}

func (account *BankingAccount) GetBalance() int {
	account.RWMutext.RLock()
	balance := account.Balance
	account.RWMutext.RUnlock()
	return balance
}

func TestRWMutex(t *testing.T) {

	account := BankingAccount{}

	for i := 1; i <= 1000; i++ {
		go func() {
			for j := 1; j <= 100; j++ {
				account.AddBalance(1)
				fmt.Println(account.GetBalance())
			}
		}()
	}

	time.Sleep(5 * time.Second)
	fmt.Println(account.Balance)
}
