// Package nutsdb use the https://github.com/xujiajun/nutsdb as cache driver
package nutsdb

import "time"

// NutsDB definition
type NutsDB struct {
}

func (c *NutsDB) Has(key string) bool {
	panic("implement me")
}

func (c *NutsDB) Get(key string) interface{} {
	panic("implement me")
}

func (c *NutsDB) Set(key string, val interface{}, ttl time.Duration) (err error) {
	panic("implement me")
}

func (c *NutsDB) Del(key string) error {
	panic("implement me")
}

func (c *NutsDB) GetMulti(keys []string) map[string]interface{} {
	panic("implement me")
}

func (c *NutsDB) SetMulti(values map[string]interface{}, ttl time.Duration) (err error) {
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
