package cache

import (
	"sync"
	"time"
)

// Item for memory cache
type Item struct {
	// Exp expire time
	Exp int64
	// Val cache value storage
	Val interface{}
}

// Expired check whether expired
func (item Item) Expired() bool {
	return item.Exp > 1 && item.Exp < time.Now().Unix()
}

// MemoryCache definition.
type MemoryCache struct {
	// locker
	lock sync.RWMutex
	// cache data in memory. or use sync.Map
	caches map[string]*Item
	// CacheSize TODO set max cache size
	CacheSize int
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
	defer c.lock.RUnlock()

	return c.get(key) != nil
}

// Get cache value by key
func (c *MemoryCache) Get(key string) interface{} {
	c.lock.RLock()
	defer c.lock.RUnlock()

	return c.get(key)
}

func (c *MemoryCache) get(key string) interface{} {
	if item, ok := c.caches[key]; ok {
		// check expire time. if has been expired, remove it.
		if item.Expired() {
			_ = c.del(key)
		}

		return item.Val
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

	data := make(map[string]interface{}, len(keys))
	for _, key := range keys {
		data[key] = c.get(key)
	}

	c.lock.RUnlock()
	return data
}

// SetMulti values by multi key
func (c *MemoryCache) SetMulti(values map[string]interface{}, ttl time.Duration) (err error) {
	c.lock.Lock()
	for key, val := range values {
		if err = c.set(key, val, ttl); err != nil {
			return
		}
	}

	c.lock.Unlock()
	return
}

// DelMulti values by multi key
func (c *MemoryCache) DelMulti(keys []string) error {
	c.lock.Lock()
	for _, key := range keys {
		_ = c.del(key)
	}

	c.lock.Unlock()
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
