// bank_test.go

package bank_test

import (
	"fmt"
	"sync"
	"testing"

	"bank"
)

func TestBank(t *testing.T) {
	// Deposit [1..1000] concurrently.
	var n sync.WaitGroup
	for i := 1; i <= 1000; i++ {
		n.Add(1)
		go func(amount int) {
			bank.Deposit(amount)
			n.Done()
		}(i)
	}
	n.Wait()
	accountBalance := bank.Balance()
	fmt.Print("Account balance is: $", accountBalance, "\n")
	if got, want := bank.Balance(), (1000+1)*1000/2; got != want {
		t.Errorf("Balance = %d, want %d", got, want)
	}
}
