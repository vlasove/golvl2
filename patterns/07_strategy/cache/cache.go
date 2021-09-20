package cache

// Контекст (объект, к которому применяются стратегии)
type Cache struct {
	Storage      map[string]string
	EvictionAlgo EvictionAlgo
	Capacity     int
	MaxCapacity  int
}

func New(e EvictionAlgo, capacity int, maxCapacity int) *Cache {
	return &Cache{
		Storage:      make(map[string]string),
		EvictionAlgo: e,
		Capacity:     capacity,
		MaxCapacity:  maxCapacity,
	}
}

func (c *Cache) SetEvictionAlgo(e EvictionAlgo) {
	c.EvictionAlgo = e
}

func (c *Cache) Add(key, val string) {
	if c.Capacity == c.MaxCapacity {
		c.Evict()
	}
	c.Capacity++
	c.Storage[key] = val
}

func (c *Cache) Get(key string) string {
	result := c.Storage[key]
	delete(c.Storage, key)
	return result
}

func (c *Cache) Evict() {
	c.EvictionAlgo.Evict(c)
	c.Capacity--
}
