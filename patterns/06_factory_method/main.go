package main

import (
	"fmt"
	"log"

	"github.com/vlasove/materials/tasks_2/patterns/fabricmethod/gunfactory"
)

// Фабричный метод прочитан еще раз!
func main() {
	ak47, err := gunfactory.GetGun("ak47", 4)
	if err != nil {
		log.Fatal(err)
	}
	defer fmt.Println(ak47)

	musket, err := gunfactory.GetGun("musket", 900)
	if err != nil {
		log.Fatal(err)
	}
	defer fmt.Println(musket)
}
