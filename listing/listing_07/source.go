// Что выведет программа? Объяснить вывод программы.
package listing_07

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// Наполнение канала значениями с случайными таймаутами
func asChan(vs ...int) <-chan int {
	c := make(chan int)

	go func() {
		for _, v := range vs {
			c <- v
			time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
		}

		close(c)
	}()
	return c
}

// Слияние каналов
func merge(a, b <-chan int) <-chan int {
	c := make(chan int)
	go func() {
		for {
			// Проблема в том, что когда каналы закроются,
			// они будут возвращать дефолтные значения,
			// для проверки, закрыт ли канал - можно посмотреть на второе значение
			select {
			case v := <-a:
				c <- v

			case v := <-b:
				c <- v

			}
		}
	}()
	return c
}

// рабочий вариант merge
func myMerge(a, b <-chan int) <-chan int {
	var wg sync.WaitGroup
	out := make(chan int)

	output := func(c <-chan int) {
		for data := range c {
			out <- data
		}
		wg.Done()
	}
	wg.Add(2)
	go output(a)
	go output(b)

	go func() {
		wg.Wait()
		close(out)
	}()
	return out
}

func Start() {

	a := asChan(1, 3, 5, 7)
	b := asChan(2, 4, 6, 8)
	c := myMerge(a, b)
	for v := range c {
		fmt.Println(v)
	}
}
