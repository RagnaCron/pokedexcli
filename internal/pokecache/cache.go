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
	mux      sync.RWMutex
	interval time.Duration
}

func NewCache(interval time.Duration) *Cache {
	cache := &Cache{
		entries:  make(map[string]cacheEntry),
		mux:      sync.RWMutex{},
		interval: interval,
	}

	go cache.reapLoop()

	return cache
}

func (c *Cache) reapLoop() {
	ticker := time.NewTicker(c.interval)
	defer ticker.Stop()

	for range ticker.C {
		c.mux.Lock()
		for key, entry := range c.entries {
			if time.Since(entry.createdAt) > c.interval {
				delete(c.entries, key)
			}
		}
		c.mux.Unlock()
	}
}

func (c *Cache) Add(key string, val []byte) {
	c.mux.Lock()

	c.entries[key] = cacheEntry{
		createdAt: time.Now(),
		val:       val,
	}

	c.mux.Unlock()
}

func (c *Cache) Get(key string) ([]byte, bool) {
	c.mux.RLock()
	entry, ok := c.entries[key]
	c.mux.RUnlock()

	if !ok {
		return nil, false
	}

	return entry.val, true
}
