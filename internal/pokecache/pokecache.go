package pokecache

import (
	"sync"
	"time"
)

type Cache struct {
	cacheData map[string]cacheEntry
	mu        *sync.Mutex
}

type cacheEntry struct {
	createdAt time.Time
	entryData []byte
}

// Creating a New Cache
func NewCache(interval time.Duration) Cache {
	c := Cache{
		cacheData: make(map[string]cacheEntry),
		mu:        &sync.Mutex{},
	}
	c.reapLoop(interval)
	return c
}

// Adding data to existing Cache
func (c *Cache) Add(key string, val []byte) {
	c.mu.Lock()
	data := c.cacheData[key]
	data.entryData = val
	data.createdAt = time.Now()
	c.cacheData[key] = data
	c.mu.Unlock()
}

// Getting data from existing Cache
func (c *Cache) Get(key string) ([]byte, bool) {
	c.mu.Lock()
	defer c.mu.Unlock()
	data, ok := c.cacheData[key]

	if ok {
		return data.entryData, true
	}
	return nil, false
}

// Cleans out old data from cache after a certain interval
func (c *Cache) reapLoop(interval time.Duration) {
	ticker := time.NewTicker(interval)

	go func() {
		for t := time.Now(); true; t = <-ticker.C {
			c.mu.Lock()
			for k, v := range c.cacheData {
				if t.Sub(v.createdAt) > interval {
					delete(c.cacheData, k)
				}
			}
			c.mu.Unlock()
		}
	}()
}
