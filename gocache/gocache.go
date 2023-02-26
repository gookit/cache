// Package gocache is a memory cache driver implement.
// base on the package: github.com/patrickmn/go-cache
//
// Usage:
//
//	import "github.com/gookit/cache"
//
//	cache.Register(gocache.NewGoCache(0, cache.FiveMinutes))
//	// use
//	// cache.Set("key", "value")
package gocache

import (
	"time"

	goc "github.com/patrickmn/go-cache"
)

// Name driver name
const Name = "gocache"

// GoCache struct
type GoCache struct {
	db *goc.Cache
	// will handle expire on has,get
	expireManually bool
}

// New create instance
func New() *GoCache {
	return NewSimple()
}

// NewSimple create new simple instance
func NewSimple() *GoCache {
	return &GoCache{
		db: goc.New(goc.NoExpiration, goc.NoExpiration),
		// handle expire on has,get
		expireManually: true,
	}
}

// NewGoCache create instance with settings
func NewGoCache(defaultExpiration, cleanupInterval time.Duration) *GoCache {
	return &GoCache{
		db: goc.New(defaultExpiration, cleanupInterval),
	}
}

// Close connection
func (g *GoCache) Close() error {
	return nil
}

// Clear all caches
func (g *GoCache) Clear() error {
	g.db.Flush()
	return nil
}

// Has cache key
func (g *GoCache) Has(key string) bool {
	return g.Get(key) != nil
}

// Get cache by key
func (g *GoCache) Get(key string) any {
	if g.expireManually {
		g.db.DeleteExpired()
	}

	val, _ := g.db.Get(key)
	return val
}

// Set cache by key
func (g *GoCache) Set(key string, val any, ttl time.Duration) error {
	g.db.Set(key, val, ttl)
	return nil
}

// Del cache by key
func (g GoCache) Del(key string) error {
	g.db.Delete(key)
	return nil
}

// GetMulti cache by keys
func (g *GoCache) GetMulti(keys []string) map[string]any {
	data := make(map[string]any, len(keys))

	for _, key := range keys {
		val, ok := g.db.Get(key)
		if ok {
			data[key] = val
		}
	}

	return data
}

// SetMulti cache by keys
func (g GoCache) SetMulti(values map[string]any, ttl time.Duration) error {
	for key, val := range values {
		g.db.Set(key, val, ttl)
	}
	return nil
}

// DelMulti db by keys
func (g *GoCache) DelMulti(keys []string) error {
	for _, key := range keys {
		g.db.Delete(key)
	}
	return nil
}

// Db get the goc.Cache
func (g *GoCache) Db() *goc.Cache {
	return g.db
}
