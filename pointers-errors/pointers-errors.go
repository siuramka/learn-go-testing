package pointers_errors

import (
	"errors"
	"fmt"
)

type Bitcoin int

func (b Bitcoin) String() string {
	return fmt.Sprintf("%d BTC", b)
}

type Wallet struct {
	balance Bitcoin
}

func (w *Wallet) Deposit(m Bitcoin) {
	w.balance += m
}

const overdrawErrorMessage = "overdraw"

func (w *Wallet) Withdraw(m Bitcoin) error {
	if w.Balance() < m {
		return errors.New(overdrawErrorMessage)
	}
	w.balance -= m
	return nil
}

func (w *Wallet) Balance() Bitcoin {
	return w.balance
}
