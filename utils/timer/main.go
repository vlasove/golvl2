package main

import (
	"fmt"
	"log"

	"github.com/vlasove/materials/tasks_2/utils/timer/clock"
)

func main() {
	c, err := clock.New(clock.BaseHost)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(c.CurrentTime())
}
