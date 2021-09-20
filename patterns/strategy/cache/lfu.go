package cache

import "log"

type Lfu struct {
}

func (l *Lfu) Evict(c *Cache) {
	log.Println("evicting by LFU strtegy")
}
