package pokecache

import (
	"sync"
	"time"
)

type Cache struct {
	cache    map[string]cacheEntry
	mutex    sync.Mutex
	interval time.Duration
}

type cacheEntry struct {
	createdAt time.Time
	val       []byte
}

func NewCache(inter time.Duration) *Cache {
	newC := &Cache{
		cache:    make(map[string]cacheEntry),
		mutex:    sync.Mutex{},
		interval: inter,
	}

	//goroutine
	go newC.reapLoop()

	return newC
}

func (c *Cache) Add(key string, val []byte) {
	c.mutex.Lock()
	c.cache[key] = cacheEntry{
		createdAt: time.Now(),
		val:       val,
	}
	c.mutex.Unlock()
}

func (c *Cache) Get(key string) ([]byte, bool) {
	c.mutex.Lock()

	result, ok := c.cache[key]
	//c.cache[key]
	c.mutex.Unlock()
	if !ok {
		return nil, false
	}

	return result.val, true
}

func (c *Cache) reapLoop() {
	//call by NewCache
	//Each time an interval (the time.Duration passed to NewCache)
	//passes it should remove any entries that are older than the interval.
	//This makes sure that the cache doesn't grow too large over time.

	ticker := time.NewTicker(c.interval)
	for range ticker.C {
		c.mutex.Lock()
		for k, v := range c.cache {
			elapsed := time.Now().Sub(v.createdAt)
			if elapsed > c.interval {
				delete(c.cache, k)
			}
		}
		c.mutex.Unlock()
	}

}
