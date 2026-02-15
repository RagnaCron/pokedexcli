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
	Entries  map[string]cacheEntry
	Mux      sync.RWMutex
	Interval time.Duration
}

func NewCache(interval time.Duration) *Cache {
	cache := &Cache{
		Entries:  make(map[string]cacheEntry),
		Mux:      sync.RWMutex{},
		Interval: interval,
	}

	go cache.reapLoop()

	return cache
}

func (c *Cache) reapLoop() {
	ticker := time.NewTicker(c.Interval)
	defer ticker.Stop()

	for range ticker.C {
		c.Mux.Lock()
		for key, entry := range c.Entries {
			if time.Since(entry.createdAt) > c.Interval {
				delete(c.Entries, key)
			}
		}
		c.Mux.Unlock()
	}
}

func (c *Cache) Add(key string, val []byte) {
	c.Mux.Lock()

	c.Entries[key] = cacheEntry{
		createdAt: time.Now(),
		val:       val,
	}

	c.Mux.Unlock()
}

func (c *Cache) Get(key string) ([]byte, bool) {
	c.Mux.RLock()
	entry, ok := c.Entries[key]
	c.Mux.RUnlock()

	if !ok {
		return nil, false
	}

	return entry.val, true
}
