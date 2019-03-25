// Package leveldb use the https://github.com/syndtr/goleveldb as cache driver
package leveldb

import "time"

// LevelDB definition
type LevelDB struct {
	db *leveldb.DB
}

func (c *LevelDB) Has(key string) bool {
	panic("implement me")
}

func (c *LevelDB) Get(key string) interface{} {
	panic("implement me")
}

func (c *LevelDB) Set(key string, val interface{}, ttl time.Duration) (err error) {
	panic("implement me")
}

func (c *LevelDB) Del(key string) error {
	panic("implement me")
}

func (c *LevelDB) GetMulti(keys []string) map[string]interface{} {
	panic("implement me")
}

func (c *LevelDB) SetMulti(values map[string]interface{}, ttl time.Duration) (err error) {
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


