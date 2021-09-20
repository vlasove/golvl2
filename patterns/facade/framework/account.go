package framework

import (
	"errors"
	"log"
)

var (
	errIncorrectName = errors.New("account name is incorrect")
)

// account - базовый аккаунт
type Account struct {
	name string
}

func NewAccount(accountName string) *Account {
	return &Account{
		name: accountName,
	}
}
func (a *Account) CheckAccount(accountName string) error {
	if a.name != accountName {
		return errIncorrectName
	}
	log.Println("account verified successfully")
	return nil
}
