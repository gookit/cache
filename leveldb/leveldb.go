// Package leveldb use the https://github.com/syndtr/goleveldb as cache driver
package leveldb

import (
	"time"

	"github.com/syndtr/goleveldb/leveldb"
)

// Name driver name
const Name = "leveldb"

// LevelDB definition
type LevelDB struct {
	// cache.BaseDriver
	db *leveldb.DB
}

func (c *LevelDB) Has(key string) bool {
	panic("implement me")
}

func (c *LevelDB) Get(key string) any {
	panic("implement me")
}

func (c *LevelDB) Set(key string, val any, ttl time.Duration) (err error) {
	panic("implement me")
}

func (c *LevelDB) Del(key string) error {
	panic("implement me")
}

func (c *LevelDB) GetMulti(keys []string) map[string]any {
	panic("implement me")
}

func (c *LevelDB) SetMulti(values map[string]any, ttl time.Duration) (err error) {
	panic("implement me")
}

func (c *LevelDB) DelMulti(keys []string) error {
	panic("implement me")
}

func (c *LevelDB) Clear() error {
	panic("implement me")
}

func (c *LevelDB) Close() error {
	return c.db.Close()
}
