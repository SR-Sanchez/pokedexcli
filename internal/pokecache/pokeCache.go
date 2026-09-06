package pokecache

import (
	"sync"
	"time"
)

// cacheEntry represents a single stored item in the cache.
type cacheEntry struct {
	createdAt time.Time // when this entry was added, used to check staleness
	val       []byte    // the raw cached data (e.g. an HTTP response body)
}

// Cache holds a thread-safe map of cached entries.
type Cache struct {
	cache    map[string]cacheEntry
	mu       sync.Mutex    // protects `cache` from concurrent read/write access
	interval time.Duration // how old an entry can get before it's reaped
}

// Add inserts or overwrites a cache entry for the given key.
func (c *Cache) Add(key string, val []byte) {
	c.mu.Lock()         // only one goroutine may touch the map at a time
	defer c.mu.Unlock() // safe here because Add returns quickly (no infinite loop)
	c.cache[key] = cacheEntry{
		createdAt: time.Now(),
		val:       val,
	}
}

// Get retrieves a cache entry for the given key, if it exists.
func (c *Cache) Get(key string) ([]byte, bool) {
	c.mu.Lock()
	defer c.mu.Unlock()
	entry, ok := c.cache[key]
	if !ok {
		return nil, false
	}
	return entry.val, true
}

// reapLoop runs forever in the background, periodically removing
// entries older than c.interval. Started as a goroutine so it never
// blocks the caller.
func (c *Cache) reapLoop() {
	ticker := time.NewTicker(c.interval) // fires once per c.interval
	for range ticker.C {                 // blocks until each tick arrives
		c.mu.Lock()
		for key, entry := range c.cache {
			if time.Since(entry.createdAt) > c.interval {
				delete(c.cache, key)
			}
		}
		c.mu.Unlock() // explicit unlock each tick (NOT defer - this loop never returns)
	}
}

// NewCache creates a ready-to-use cache and starts its background
// reaping goroutine immediately.
func NewCache(interval time.Duration) *Cache {
	c := map[string]cacheEntry{}
	ca := &Cache{
		cache:    c,
		interval: interval,
	}
	go ca.reapLoop() // runs concurrently; NewCache returns without waiting
	return ca
}