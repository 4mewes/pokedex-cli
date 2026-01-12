package pokecache

import (
	"fmt"
	"sync"
	"time"
)

type Cache struct {
	cacheEntries map[string]cacheEntry
	cacheMutex   sync.Mutex
}

type cacheEntry struct {
	createAt time.Time
	val      []byte
}

func NewCache(interval time.Duration) Cache {
	fmt.Println(interval)
	newCache := Cache{
		cacheEntries: make(map[string]cacheEntry),
	}
	go newCache.reapLoop(interval)
	return newCache
}

func (c *Cache) Add(key string, val []byte) {
	newEntry := cacheEntry{
		createAt: time.Now(),
		val:      val,
	}
	c.cacheMutex.Lock()
	c.cacheEntries[key] = newEntry
	c.cacheMutex.Unlock()
}

func (c *Cache) Get(key string) ([]byte, bool) {
	entry, ok := c.cacheEntries[key]
	if !ok {
		fmt.Printf("cache miss for key:%s\n", key)
		return nil, false
	}
	fmt.Printf("retrieved from cache key:%s, val:%s\n", key, entry.val)
	return entry.val, true
}

func (c *Cache) reapLoop(interval time.Duration) {
	ticker := time.NewTicker(interval)
	defer ticker.Stop()
	for range ticker.C {
		for k, v := range c.cacheEntries {
			if v.createAt.Before(time.Now().Add(-interval)) {
				c.cacheMutex.Lock()
				delete(c.cacheEntries, k)
				c.cacheMutex.Unlock()
			}
		}
	}
}
