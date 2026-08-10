package account

import (
	"fmt"
	"time"
)

type Account struct {
	account_holder string
	amount         float64
}

func (a *Account) Init_account(holder string) {
	a.account_holder = holder
	now := get_time()
	fmt.Printf("Account opened for %s at %s\n", holder, now)
}

func (a *Account) Deposit(deposit_amount float64) {
	a.amount = a.amount + deposit_amount
}

func (a *Account) Withdraw(withdraw_amount float64) {
	a.amount = a.amount - withdraw_amount
}

func (a Account) GetAmount() float64 {
	return a.amount
}

func (a Account) GetHolder() string {
	return a.account_holder
}

func (a Account) Display() {
	fmt.Printf("Holder : %s\nAmount : %.2f\n", a.account_holder, a.amount)
}

// HELPER
func get_time() string {
	now := time.Now()
	return fmt.Sprintf("%s %d, %d", now.Month(), now.Day(), now.Year())
}
