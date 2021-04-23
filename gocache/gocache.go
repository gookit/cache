package gocache

import (
	goc "github.com/patrickmn/go-cache"
	"time"
)

type GoCache struct {
	cache *goc.Cache
}

func NewGoCache(defaultExpiration, cleanupInterval time.Duration) *GoCache {
	c := goc.New(defaultExpiration, cleanupInterval)
	return &GoCache{
		cache: c,
	}
}

func (g GoCache) Close() error {
	panic("implement me")
}

func (g GoCache) Clear() error {
	panic("implement me")
}

func (g GoCache) Has(key string) bool {
	return g.Get(key) != nil
}

func (g GoCache) Get(key string) interface{} {
	val, _ := g.cache.Get(key)
	return val
}

func (g GoCache) Set(key string, val interface{}, ttl time.Duration) (err error) {
	g.cache.Set(key, val, ttl)
	return nil
}

func (g GoCache) Del(key string) error {
	g.cache.Delete(key)
	return nil
}

func (g GoCache) GetMulti(keys []string) map[string]interface{} {
	return nil
}

func (g GoCache) SetMulti(values map[string]interface{}, ttl time.Duration) (err error) {
	return nil
}

func (g GoCache) DelMulti(keys []string) error {
	return nil
}
