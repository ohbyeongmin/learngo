package accounts

import (
	"errors"
	"fmt"
)

// Account struct
type account struct {
	owner string
	balance int
}

var errNoMoney = errors.New("ERROR : Can't withdraw")

// NewAccount creates account
func NewAccount(owner string) *account{
	account := account{owner: owner, balance: 0}
	return &account
}


// receiver
// Deposit x account on your account
func (a *account) Deposit(amount int) {
	a.balance += amount
}

// Balance of your account
func (a account) Balance() int {
	return a.balance
}

// Withdraw x account from your account
func (a *account) Withdraw(amount int) error {
	if a.balance < amount {
		return errNoMoney
	}
	a.balance -= amount
	return nil
}

// ChangeOwner x account from your account
func (a *account) ChangeOwner(newOwner string) {
	a.owner = newOwner
}


// Owner of the account
func (a account) Owner() string {
	return a.owner
}

// 내부적으로 호출 됨
// account print
func (a account) String() string{
	return fmt.Sprint(a.Owner(), "'s account.\nHas: ", a.Balance())
}