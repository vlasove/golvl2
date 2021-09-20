package cache

import "log"

// Конкретная реализация стратегии - LRU (Least Recent Use)
type Lru struct{}

func (l *Lru) Evict(c *Cache) {
	log.Println("evicting by LRU strtegy")
}
