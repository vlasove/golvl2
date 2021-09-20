package main

import (
	"log"

	"github.com/vlasove/materials/tasks_2/utils/unpack/helper"
)

func main() {
	msg := `a2b3c\46`
	res, err := helper.Unpack(msg)
	if err != nil {
		log.Fatal(err)
	}
	log.Println(res)
}
