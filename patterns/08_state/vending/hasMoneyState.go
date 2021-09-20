package vending

import (
	"errors"
	"log"
)

// Конкретное состояние - деньги за товар внесены
type HasMoneyState struct {
	VendingMachine *VendingMachine
}

func (i *HasMoneyState) RequestItem() error {
	return errors.New("item dispense in progress")
}

func (i *HasMoneyState) AddItem(count int) error {
	return errors.New("item dispense in progress")
}

func (i *HasMoneyState) InsertMoney(money int) error {
	return errors.New("item out of stock")
}
func (i *HasMoneyState) DispenseItem() error {
	log.Println("dispensing item")
	i.VendingMachine.ItemCount = i.VendingMachine.ItemCount - 1
	if i.VendingMachine.ItemCount == 0 {
		i.VendingMachine.SetState(i.VendingMachine.NoItem)
	} else {
		i.VendingMachine.SetState(i.VendingMachine.HasItem)
	}
	return nil
}
