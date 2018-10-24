// Package gcache use the github.com/bluele/gcache as cache driver
package gcache

import "github.com/bluele/gcache"

// GCache driver definition
type GCache struct {
	db gcache.Cache
}

func New(size int) *GCache {
	return &GCache{
		db: gcache.New(20).LRU().Build(),
	}
}
