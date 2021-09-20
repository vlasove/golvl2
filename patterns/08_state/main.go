package main

import (
	"log"

	"github.com/vlasove/golvl2/patterns/08_state/vending"
)

func main() {
	vm := vending.New(1, 10)

	if err := vm.RequestItem(); err != nil {
		log.Fatal(err)
	}

	if err := vm.InsertMoney(10); err != nil {
		log.Fatal(err)
	}

	if err := vm.DispenseItem(); err != nil {
		log.Fatal(err)
	}

	if err := vm.AddItem(2); err != nil {
		log.Fatal(err)
	}

	if err := vm.RequestItem(); err != nil {
		log.Fatal(err)
	}

	if err := vm.InsertMoney(10); err != nil {
		log.Fatal(err)
	}

	if err := vm.DispenseItem(); err != nil {
		log.Fatal(err)
	}

}
