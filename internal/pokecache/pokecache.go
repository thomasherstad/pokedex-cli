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
	cache map[string]cacheEntry
	mu    *sync.Mutex
}

func NewCache(interval time.Duration) Cache {
	var cache Cache
	cache.cache = make(map[string]cacheEntry)
	cache.mu = &sync.Mutex{}

	go cache.ReapLoop(interval) //Må ikke denne være concurrent?

	return cache
}

func (c Cache) Add(key string, val []byte) {
	var entry cacheEntry
	entry.createdAt = time.Now()
	entry.val = val

	c.mu.Lock()
	c.cache[key] = entry
	c.mu.Unlock()
}

func (c Cache) Get(key string) ([]byte, bool) {
	c.mu.Lock()
	data, ok := c.cache[key]
	c.mu.Unlock()
	if !ok {
		return nil, false
	}
	return data.val, true
}

func (c Cache) Delete(key string) {
	//delete entry
	c.mu.Lock()
	defer c.mu.Unlock()
	delete(c.cache, key)
}

func (c Cache) ReapLoop(interval time.Duration) {
	ticker := time.NewTicker(interval)
	for range ticker.C {
		c.reap(time.Now(), interval)
	}
}

func (c Cache) reap(now time.Time, interval time.Duration) {
	for key, entry := range c.cache {
		elapsed := now.Sub(entry.createdAt)
		if elapsed > interval {
			c.Delete(key)
		}
	}

}
