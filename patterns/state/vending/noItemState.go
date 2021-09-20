package vending

import "errors"

// Конкретное состояние - нет товара
type NoItemState struct {
	VendingMachine *VendingMachine
}

func (i *NoItemState) RequestItem() error {
	return errors.New("item out of stock")
}

func (i *NoItemState) AddItem(count int) error {
	i.VendingMachine.IncrementItemCount(count)
	i.VendingMachine.SetState(i.VendingMachine.HasItem)
	return nil
}

func (i *NoItemState) InsertMoney(money int) error {
	return errors.New("item out of stock")
}

func (i *NoItemState) DispenseItem() error {
	return errors.New("item out of stock")
}
