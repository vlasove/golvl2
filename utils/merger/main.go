package main

import (
	"fmt"
	"time"
)

func or(channs ...<-chan interface{}) <-chan interface{} {
	switch len(channs) {
	case 0:
		return nil
	case 1:
		return channs[0]
	}
	orDone := make(chan interface{})
	go func() {
		defer close(orDone)

		switch len(channs) {
		case 2:
			select {
			case <-channs[0]:
			case <-channs[1]:
			}
		default:
			select {
			case <-channs[0]:
			case <-channs[1]:
			case <-channs[2]:
			case <-or(append(channs[3:], orDone)...):
			}
		}
	}()
	return orDone
}

func main() {
	sig := func(after time.Duration) <-chan interface{} {
		c := make(chan interface{})
		go func() {
			defer close(c)
			time.Sleep(after)
		}()
		return c
	}

	start := time.Now()
	<-or(
		sig(2*time.Hour),
		sig(5*time.Minute),
		sig(1*time.Second),
		sig(1*time.Hour),
		sig(1*time.Minute),
	)

	fmt.Printf("done after %v\n", time.Since(start))
}
