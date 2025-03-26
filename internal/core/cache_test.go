package core

import (
	"testing"
	"time"
)

func TestMetricsCache(t *testing.T) {
	cache := &MetricsCache{
		ttl:   2 * time.Second,
		cache: make(map[string]CacheEntry),
	}

	// Test Set et Get
	cache.Set("test", 123)
	if val, found := cache.Get("test"); !found || val.(int) != 123 {
		t.Error("Cache Set/Get failed")
	}

	// Test expiration
	time.Sleep(3 * time.Second)
	if _, found := cache.Get("test"); found {
		t.Error("Cache entry should have expired")
	}

	// Test Cleanup
	cache.Set("test1", 1)
	cache.Set("test2", 2)
	time.Sleep(3 * time.Second)
	cache.Cleanup()
	if len(cache.cache) > 0 {
		t.Error("Cache cleanup failed")
	}
}
