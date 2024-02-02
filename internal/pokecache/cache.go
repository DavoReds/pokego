package pokecache

import (
	"sync"
	"time"
)

type Cache struct {
	entries map[string]cacheEntry
	mu      sync.Mutex
}

type cacheEntry struct {
	createdAt time.Time
	val       []byte
}

// Creates a new Cache.
func NewCache(interval time.Duration) *Cache {
	cache := &Cache{
		entries: make(map[string]cacheEntry),
	}

	cache.reapLoop(interval)

	return cache
}

func (cache *Cache) reapLoop(interval time.Duration) {
	ticker := time.NewTicker(interval)

	go func() {
		for range ticker.C {
			cache.deleteStale(interval)
		}
	}()
}

// Adds an entry into the cache, overwriting it if it already exists.
func (cache *Cache) Add(key string, val []byte) {
	cache.mu.Lock()
	defer cache.mu.Unlock()

	cache.entries[key] = cacheEntry{
		createdAt: time.Now(),
		val:       val,
	}
}

// Returns an array of bytes corresponding to an entry in the cache and a
// boolean value representing whether or not the value exists in the cache.
func (cache *Cache) Get(key string) ([]byte, bool) {
	cache.mu.Lock()
	defer cache.mu.Unlock()

	if value, exists := cache.entries[key]; exists {
		return value.val, true
	} else {
		return []byte{}, false
	}
}

func (cache *Cache) deleteStale(duration time.Duration) {
	cache.mu.Lock()
	defer cache.mu.Unlock()

	for key, entry := range cache.entries {
		if time.Since(entry.createdAt) > duration {
			delete(cache.entries, key)
		}
	}
}
