package cache

import (
	"fmt"
	"time"
)

func NewCache(interval time.Duration) *Cache {
	c := &Cache{
		data: make(map[string]cacheEntry),
	}
	go c.reapLoop(interval)
	return c
}

func (c *Cache) Add(key string, val []byte) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.data[key] = cacheEntry{
		createdAt: time.Now(),
		val:       val,
	}
}
func (c *Cache) Get(key string) ([]byte, bool) {
	c.mu.Lock()
	defer c.mu.Unlock()
	cacheData, ok := c.data[key]
	if !ok {
		return []byte{}, false
	}
	return cacheData.val, true
}

func (c *Cache) reapLoop(interval time.Duration) {
	ticker := time.NewTicker(interval)
	for range ticker.C {
		c.mu.Lock()
		for key, entery := range c.data {
			if time.Since(entery.createdAt) > interval {
				delete(c.data, key)
				fmt.Println("Cache expeired ", key)
			}
		}
		c.mu.Unlock()
	}

}
