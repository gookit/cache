package gocache

import (
	"time"

	"github.com/gookit/cache"
	goc "github.com/patrickmn/go-cache"
)

// GoCache struct
type GoCache struct {
	cache.BaseDriver
	cache *goc.Cache
	// expire handle
	expireManually bool
}

// New create instance
func New() *GoCache {
	c := goc.New(goc.NoExpiration, goc.NoExpiration)
	return &GoCache{
		cache:          c,
		expireManually: true,
	}
}

// NewGoCache create instance with settings
func NewGoCache(defaultExpiration, cleanupInterval time.Duration) *GoCache {
	c := goc.New(defaultExpiration, cleanupInterval)
	return &GoCache{
		cache: c,
	}
}

// Close connection
func (g *GoCache) Close() error {
	return nil
}

// Clear all caches
func (g *GoCache) Clear() error {
	g.cache.Flush()
	return nil
}

// Has cache key
func (g *GoCache) Has(key string) bool {
	return g.Get(key) != nil
}

// Get cache by key
func (g *GoCache) Get(key string) interface{} {
	if g.expireManually {
		g.cache.DeleteExpired()
	}
	val, _ := g.cache.Get(key)
	return val
}

// Set cache by key
func (g *GoCache) Set(key string, val interface{}, ttl time.Duration) (err error) {
	g.cache.Set(key, val, ttl)
	return nil
}

// Del cache by key
func (g GoCache) Del(key string) error {
	g.cache.Delete(key)
	return nil
}

func (g *GoCache) GetMulti(keys []string) map[string]interface{} {
	panic("implement me")
}

func (g GoCache) SetMulti(values map[string]interface{}, ttl time.Duration) (err error) {
	panic("implement me")
}

func (g *GoCache) DelMulti(keys []string) error {
	panic("implement me")
}
