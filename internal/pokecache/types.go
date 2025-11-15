package pokecache

import (
	"sync"
	"time"
)

type cacheEntry struct {
	createdAt time.Time
	val       []byte
}

type Cache struct {
	entries  map[string]cacheEntry
	mux      *sync.Mutex
	interval time.Duration
}

func NewCache(interval time.Duration) *Cache {
	cache := &Cache{
		interval: interval,
		mux:      &sync.Mutex{},
		entries:  map[string]cacheEntry{},
	}

	go cache.reapLoop()

	return cache
}

func (c *Cache) Add(key string, val []byte) error {
	c.mux.Lock()
	defer c.mux.Unlock()
	c.entries[key] = cacheEntry{
		createdAt: time.Now().UTC(),
		val:       val,
	}

	return nil
}

func (c *Cache) Get(key string) ([]byte, bool) {
	c.mux.Lock()
	defer c.mux.Unlock()

	value, ok := c.entries[key]

	return value.val, ok
}

func (c *Cache) reapLoop() {
	ticker := time.NewTicker(c.interval)
	defer ticker.Stop()

	for range ticker.C {
		c.mux.Lock()
		for key, val := range c.entries {
			now := time.Now().UTC()
			if val.createdAt.Before(now.Add(-c.interval)) {
				delete(c.entries, key)
			}
		}
		c.mux.Unlock()
	}
}
