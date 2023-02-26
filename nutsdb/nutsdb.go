// Package nutsdb use the https://github.com/xujiajun/nutsdb as cache driver
package nutsdb

import (
	"time"
)

// Name driver name
const Name = "nutsdb"

// NutsDB definition TODO
type NutsDB struct {
	// cache.BaseDriver
}

func (c *NutsDB) Has(key string) bool {
	panic("implement me")
}

func (c *NutsDB) Get(key string) any {
	panic("implement me")
}

func (c *NutsDB) Set(key string, val any, ttl time.Duration) (err error) {
	panic("implement me")
}

func (c *NutsDB) Del(key string) error {
	panic("implement me")
}

func (c *NutsDB) GetMulti(keys []string) map[string]any {
	panic("implement me")
}

func (c *NutsDB) SetMulti(values map[string]any, ttl time.Duration) (err error) {
	panic("implement me")
}

func (c *NutsDB) DelMulti(keys []string) error {
	panic("implement me")
}

func (c *NutsDB) Clear() error {
	panic("implement me")
}

func (c *NutsDB) Close() error {
	panic("implement me")
}
