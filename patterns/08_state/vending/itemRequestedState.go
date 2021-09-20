package vending

import (
	"errors"
	"fmt"
	"log"
)

// Конкретное состояние - товар запрошен
type ItemRequestedState struct {
	VendingMachine *VendingMachine
}

func (i *ItemRequestedState) RequestItem() error {
	return errors.New("item already requested")
}

func (i *ItemRequestedState) AddItem(count int) error {
	return errors.New("item dispense in progress")
}

func (i *ItemRequestedState) InsertMoney(money int) error {
	if money < i.VendingMachine.ItemPrice {
		return fmt.Errorf("inserted money is less. please insert %d", i.VendingMachine.ItemPrice)
	}
	log.Println("money entered is ok")
	i.VendingMachine.SetState(i.VendingMachine.HasMoney)
	return nil
}
func (i *ItemRequestedState) DispenseItem() error {
	return errors.New("please insert money first")
}
