package cache

// Общий интерфейс стратегии
type EvictionAlgo interface {
	Evict(c *Cache)
}
