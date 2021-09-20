package cache

import "log"

// Конкретная реализация стратегии - FIFO
type Fifo struct{}

func (f *Fifo) Evict(c *Cache) {
	log.Println("evicting by FIFO strtegy")
}
