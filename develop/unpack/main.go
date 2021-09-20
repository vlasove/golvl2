package main

import (
	"log"

	"github.com/vlasove/golvl2/develop/unpack/helper"
)

func main() {
	msg := `a2b3c\46`
	res, err := helper.Unpack(msg)
	if err != nil {
		log.Fatal(err)
	}
	log.Println(res)
}
