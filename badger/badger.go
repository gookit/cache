// Package badger use the https://github.com/dgraph-io/badger as cache driver
package badger

import (
	"time"

	"github.com/dgraph-io/badger"
)

// Name driver name
const Name = "badger"

// BadgerDB definition
type BadgerDB struct {
	db *badger.DB
}

func (c *BadgerDB) Has(key string) bool {
	panic("implement me")
}

func (c *BadgerDB) Get(key string) any {
	panic("implement me")
}

func (c *BadgerDB) Set(key string, val any, ttl time.Duration) (err error) {
	panic("implement me")
}

func (c *BadgerDB) Del(key string) error {
	panic("implement me")
}

func (c *BadgerDB) GetMulti(keys []string) map[string]any {
	panic("implement me")
}

func (c *BadgerDB) SetMulti(values map[string]any, ttl time.Duration) (err error) {
	panic("implement me")
}

func (c *BadgerDB) DelMulti(keys []string) error {
	panic("implement me")
}

func (c *BadgerDB) Clear() error {
	panic("implement me")
}

func (c *BadgerDB) Close() error {
	return c.db.Close()
}
