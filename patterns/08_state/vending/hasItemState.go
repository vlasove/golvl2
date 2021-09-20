package vending

import (
	"errors"
	"log"
)

// Конкретное состояние - товар имеется
type HasItemState struct {
	VendingMachine *VendingMachine
}

func (i *HasItemState) RequestItem() error {
	if i.VendingMachine.ItemCount == 0 {
		i.VendingMachine.SetState(i.VendingMachine.NoItem)
		return errors.New("no item present")
	}
	log.Println("item requestd")
	i.VendingMachine.SetState(i.VendingMachine.ItemRequested)
	return nil
}

func (i *HasItemState) AddItem(count int) error {
	log.Printf("%d items added\n", count)
	i.VendingMachine.IncrementItemCount(count)
	return nil
}

func (i *HasItemState) InsertMoney(money int) error {
	return errors.New("please select item first")
}
func (i *HasItemState) DispenseItem() error {
	return errors.New("please select item first")
}
