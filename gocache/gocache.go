package gocache

import (
	"time"

	goc "github.com/patrickmn/go-cache"
)

type GoCache struct {
	cache          *goc.Cache
	expireManually bool
}

func New() *GoCache {
	c := goc.New(goc.NoExpiration, goc.NoExpiration)
	return &GoCache{
		cache:          c,
		expireManually: true,
	}
}

func NewGoCache(defaultExpiration, cleanupInterval time.Duration) *GoCache {
	c := goc.New(defaultExpiration, cleanupInterval)
	return &GoCache{
		cache: c,
	}
}

func (g *GoCache) Close() error {
	return nil
}

func (g *GoCache) Clear() error {
	g.cache.Flush()
	return nil
}

func (g *GoCache) Has(key string) bool {
	return g.Get(key) != nil
}

func (g *GoCache) Get(key string) interface{} {
	if g.expireManually {
		g.cache.DeleteExpired()
	}
	val, _ := g.cache.Get(key)
	return val
}

func (g *GoCache) Set(key string, val interface{}, ttl time.Duration) (err error) {
	g.cache.Set(key, val, ttl)
	return nil
}

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
