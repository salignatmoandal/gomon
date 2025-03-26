package core

import (
	"sync"
	"time"
)

// This code implements a simple, thread-safe in-memory cache for storing metrics data with a configurable Time-To-Live (TTL). It consists of the following key components:
// 1. CacheEntry: A struct that stores the cached data and its expiration timestamp.
// 2. MetricsCache: A struct that maintains a map of cache entries, a TTL value, and a read/write mutex to ensure concurrent access is safe.
// 3. init: A function that initializes the global cache instance with the default TTL.
// 4. SetCacheTTL: A function to configure the TTL for the cache.
// 5. Get: A method to retrieve a cached value by key. It returns the data only if it exists and hasn't expired.
// 6. Set: A method to store a value in the cache, setting its expiration based on the current TTL.
// 7. Cleanup: A method that iterates through the cache and removes any entries that have expired.

// CacheEntry holds the cached data along with its expiration timestamp.
type CacheEntry struct {
	data      interface{} // The actual cached data.
	expiresAt time.Time   // The time when the cache entry expires.
}

// MetricsCache is a thread-safe in-memory cache for metrics data.
type MetricsCache struct {
	mu    sync.RWMutex          // Mutex to ensure safe concurrent access.
	ttl   time.Duration         // Time-to-live duration for each cache entry.
	cache map[string]CacheEntry // Map that holds cache entries indexed by a string key.
}

// Default cache configuration: TTL is set to 10 seconds.
var (
	defaultTTL = 10 * time.Second
	cache      *MetricsCache
)

// init initializes the global cache instance with the default TTL.
func init() {
	cache = &MetricsCache{
		ttl:   defaultTTL,
		cache: make(map[string]CacheEntry),
	}
}

// SetCacheTTL sets a new TTL for the cache.
// This function acquires a write lock to safely update the TTL.
func SetCacheTTL(duration time.Duration) {
	cache.mu.Lock()
	defer cache.mu.Unlock()
	cache.ttl = duration
}

// Get retrieves a value from the cache by its key.
// It returns the cached data and a boolean indicating if the data was found and valid.
func (c *MetricsCache) Get(key string) (interface{}, bool) {
	c.mu.RLock()
	defer c.mu.RUnlock()

	// Retrieve the cache entry by key.
	entry, exists := c.cache[key]
	if !exists {
		// Key not found in the cache.
		return nil, false
	}

	// Check if the cache entry has expired.
	if time.Now().After(entry.expiresAt) {
		// Entry expired; consider it as not found.
		return nil, false
	}

	// Return the cached data as it is still valid.
	return entry.data, true
}

// Set stores a value in the cache under the specified key.
// The value will expire after the duration specified by the TTL.
func (c *MetricsCache) Set(key string, value interface{}) {
	c.mu.Lock()
	defer c.mu.Unlock()

	// Create a new cache entry with the current time plus TTL as the expiration.
	c.cache[key] = CacheEntry{
		data:      value,
		expiresAt: time.Now().Add(c.ttl),
	}
}

// Cleanup removes all expired entries from the cache.
// It iterates over the cache and deletes entries that have passed their expiration time.
func (c *MetricsCache) Cleanup() {
	c.mu.Lock()
	defer c.mu.Unlock()

	now := time.Now()
	// Iterate through the cache entries.
	for key, entry := range c.cache {
		// If the current time is after the expiration time, delete the entry.
		if now.After(entry.expiresAt) {
			delete(c.cache, key)
		}
	}
}
