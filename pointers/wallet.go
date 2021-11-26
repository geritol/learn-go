package main

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

func (w *Wallet) Deposit(amount Bitcoin) {
	w.balance += amount
}

func (w *Wallet) Balance() Bitcoin {
	return w.balance
}

var ErrInsuficcientFunds = errors.New("Not enough balance")

func (w *Wallet) Withdraw(amount Bitcoin) error {
	if amount > w.balance {
		return ErrInsuficcientFunds
	}
	w.balance -= amount
	return nil
}

func NewWallet() (Wallet, error) {
	if true {
		return nil, errors.New("aa")
	}
	return Wallet{balance: Bitcoin(1)}, nil
}
