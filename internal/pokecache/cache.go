package pokecache

import (
	"sync"
	"time"
)

// Cache
type Cache struct {
	mu        sync.Mutex
	interval  time.Duration
	cacheData map[string]cacheEntry
}

type cacheEntry struct {
	createdAt time.Time
	val       []byte
}

// NewCache pointer it is necessary because of the mutex also not bad because of the dynamic mem size
func NewCache(inter time.Duration) *Cache {
	c := Cache{
		interval:  inter,
		cacheData: map[string]cacheEntry{},
	}

	go c.reapLoop()

	return &c
}

func (c *Cache) Add(key string, val []byte) {
	c.mu.Lock()
	defer c.mu.Unlock()

	c.cacheData[key] = cacheEntry{
		createdAt: time.Now(),
		val:       val,
	}
}

func (c *Cache) Get(key string) ([]byte, bool) {
	c.mu.Lock()
	defer c.mu.Unlock()

	entry, found := c.cacheData[key]

	if !found {
		return nil, false
	}

	return entry.val, true
}

func (c *Cache) reapLoop() {
	for {
		c.mu.Lock()
		for key, value := range c.cacheData {
			if time.Since(value.createdAt) >= c.interval {
				delete(c.cacheData, key)
			}
		}
		c.mu.Unlock()
		time.Sleep(c.interval)
	}

}
