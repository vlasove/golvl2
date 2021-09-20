package framework

import (
	"errors"
	"log"
)

var (
	errInvalidBalance = errors.New("balance is not sufficient")
)

// wallet - кошелек (контроль пополнения/списания)
type Wallet struct {
	balance int
}

func NewWallet() *Wallet {
	return &Wallet{
		balance: 0,
	}
}

func (w *Wallet) CreditBalance(amount int) {
	w.balance += amount
	log.Println("wallet balance added successfully")
}

func (w *Wallet) DebitBalance(amount int) error {
	if w.balance < amount {
		return errInvalidBalance
	}
	log.Println("wallet balance is sufficient")
	w.balance = w.balance - amount
	return nil
}
