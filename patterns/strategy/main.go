package main

import "github.com/vlasove/materials/tasks_2/patterns/strategy/cache"

func main() {
	lfu := &cache.Lfu{}
	myCache := cache.New(lfu, 0, 2)

	myCache.Add("a", "1")
	myCache.Add("b", "2")
	myCache.Add("c", "3")

	lru := &cache.Lru{}
	myCache.SetEvictionAlgo(lru)

	myCache.Add("d", "4")

	myCache.SetEvictionAlgo(&cache.Fifo{})
	myCache.Add("e", "5")
}
