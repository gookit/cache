// Package gcache use the github.com/bluele/gcache as cache driver
package gcache

import (
	"time"

	"github.com/bluele/gcache"
)

// Name driver name
const Name = "gcache"

// GCache driver definition
type GCache struct {
	// cache.BaseDriver
	db gcache.Cache
}

// New create an instance
func New(size int) *GCache {
	return NewWithType(size, gcache.TYPE_LRU)
}

// NewWithType create an instance with cache type
func NewWithType(size int, tp string) *GCache {
	return &GCache{
		db: gcache.New(size).EvictType(tp).Build(),
	}
}

// Close connection
func (g *GCache) Close() error {
	return nil
}

// Clear all caches
func (g *GCache) Clear() error {
	g.db.Purge()
	return nil
}

// Has cache key
func (g *GCache) Has(key string) bool {
	return g.Get(key) != nil
}

// Get cache by key
func (g *GCache) Get(key string) interface{} {
	val, _ := g.db.Get(key)
	return val
}

// Set cache by key
func (g *GCache) Set(key string, val interface{}, ttl time.Duration) (err error) {
	return g.db.SetWithExpire(key, val, ttl)
}

// Del cache by key
func (g *GCache) Del(key string) error {
	g.db.Remove(key)
	return nil
}

// GetMulti cache by keys
func (g *GCache) GetMulti(keys []string) map[string]interface{} {
	data := make(map[string]interface{}, len(keys))

	for _, key := range keys {
		val, err := g.db.Get(key)
		if err == nil {
			data[key] = val
		} // TODO log error
	}

	return data
}

// SetMulti cache by keys
func (g *GCache) SetMulti(values map[string]interface{}, ttl time.Duration) (err error) {
	for key, val := range values {
		err = g.db.SetWithExpire(key, val, ttl)
	}
	return
}

// DelMulti cache by keys
func (g *GCache) DelMulti(keys []string) error {
	for _, key := range keys {
		g.db.Remove(key)
	}
	return nil
}

// Db get the gcache.Cache
func (g *GCache) Db() gcache.Cache {
	return g.db
}
