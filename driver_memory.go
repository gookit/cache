package cache

import (
	"sync"
	"time"
)

// MemoryCache definition.
type MemoryCache struct {
	// locker
	lock sync.RWMutex
	// cache data in memory. or use sync.Map
	caches map[string]*Item
	// last error
	lastErr error
	// TODO set max cache number
	MaxCacheNum int
}

// Item for memory cache
type Item struct {
	// Exp expire time
	Exp int64
	// Val cache value storage
	Val interface{}
}

// NewMemoryCache create a memory cache instance
func NewMemoryCache() *MemoryCache {
	return &MemoryCache{
		caches: make(map[string]*Item),
	}
}

// Has cache key
func (c *MemoryCache) Has(key string) bool {
	c.lock.RLock()
	_, ok := c.caches[key]
	c.lock.RUnlock()
	return ok
}

// Get cache value by key
func (c *MemoryCache) Get(key string) interface{} {
	c.lock.RLock()
	defer c.lock.RUnlock()

	return c.get(key)
}

func (c *MemoryCache) get(key string) interface{} {
	if item, ok := c.caches[key]; ok {
		// check expire time
		if item.Exp == 0 || item.Exp > time.Now().Unix() {
			return item.Val
		}

		// has been expired. delete it.
		_ = c.Del(key)
	}

	return nil
}

// Set cache value by key
func (c *MemoryCache) Set(key string, val interface{}, ttl time.Duration) (err error) {
	c.lock.Lock()
	defer c.lock.Unlock()

	return c.set(key, val, ttl)
}

func (c *MemoryCache) set(key string, val interface{}, ttl time.Duration) (err error) {
	item := &Item{Val: val}
	if ttl > 0 {
		item.Exp = time.Now().Unix() + int64(ttl/time.Second)
	}

	c.caches[key] = item
	return
}

// Del cache by key
func (c *MemoryCache) Del(key string) error {
	c.lock.Lock()
	defer c.lock.Unlock()

	return c.del(key)
}

func (c *MemoryCache) del(key string) error {
	if _, ok := c.caches[key]; ok {
		delete(c.caches, key)
	}

	return nil
}

// GetMulti values by multi key
func (c *MemoryCache) GetMulti(keys []string) map[string]interface{} {
	c.lock.RLock()
	defer c.lock.RUnlock()

	values := make(map[string]interface{}, len(keys))
	for _, key := range keys {
		values[key] = c.get(key)
	}

	return values
}

// SetMulti values by multi key
func (c *MemoryCache) SetMulti(values map[string]interface{}, ttl time.Duration) (err error) {
	c.lock.Lock()
	defer c.lock.Unlock()

	for key, val := range values {
		if err = c.set(key, val, ttl); err != nil {
			return
		}
	}
	return
}

// DelMulti values by multi key
func (c *MemoryCache) DelMulti(keys []string) error {
	c.lock.Lock()
	defer c.lock.Unlock()

	for _, key := range keys {
		_ = c.del(key)
	}
	return nil
}

// Close cache
func (c *MemoryCache) Close() error {
	return nil
}

// Clear all caches
func (c *MemoryCache) Clear() error {
	c.caches = nil
	return nil
}

// Count cache item number
func (c *MemoryCache) Count() int {
	return len(c.caches)
}

// Restore DB from a file
func (c *MemoryCache) Restore(file string) error {
	return nil
}

// DumpDB to a file
func (c *MemoryCache) DumpDB(file string) error {
	return nil
}

// Iter iteration all caches
func (c *MemoryCache) Iter(file string) error {
	return nil
}

// LastErr get
func (c *MemoryCache) LastErr() error {
	return c.lastErr
}
