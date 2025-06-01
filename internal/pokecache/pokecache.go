package pokecache

import (
	"sync"
	"time"
)

type Cache struct {
	cache    map[string]cacheEntry
	mux      sync.Mutex
	cacheDur time.Duration
}

type cacheEntry struct {
	createdAt time.Time
	val       []byte
}

func NewCache(interval time.Duration) *Cache {
	c := new(Cache)
	c.cache = make(map[string]cacheEntry)
	c.cacheDur = interval
	go c.reapLoop()
	return c
}

func (c *Cache) reapLoop() {
	ticker := time.NewTicker(c.cacheDur)
	for range ticker.C {
		c.mux.Lock()
		for key, entry := range c.cache {
			if time.Now().Sub(entry.createdAt) > c.cacheDur {
				delete(c.cache, key)
			}
		}
		c.mux.Unlock()
	}
}

func (c *Cache) Add(key string, val []byte) {
	c.mux.Lock()
	entry := cacheEntry{time.Now(), val}
	c.cache[key] = entry
	c.mux.Unlock()
}

func (c *Cache) Get(key string) ([]byte, bool) {
	c.mux.Lock()
	entry, exists := c.cache[key]
	c.mux.Unlock()
	return entry.val, exists
}
