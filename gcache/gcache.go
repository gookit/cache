// Package gcache use the github.com/bluele/gcache as cache driver
package gcache

import (
	"github.com/bluele/gcache"
	"github.com/gookit/cache"
)

// GCache driver definition
type GCache struct {
	cache.BaseDriver
	db gcache.Cache
}

// New create an instance
func New(size int) *GCache {
	return &GCache{
		db: gcache.New(size).LRU().Build(),
	}
}
