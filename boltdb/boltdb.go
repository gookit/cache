// Package bolt use the go.etcd.io/bbolt(github.com/etcd-io/bbolt) as cache driver
package boltdb

import (
	"time"

	"github.com/gookit/cache"
	"go.etcd.io/bbolt"
)

// Name driver name
const Name = "boltDB"

// BoltDB definition
type BoltDB struct {
	cache.BaseDriver
	// db file path. eg "path/to/my.db"
	file string
	// db instance
	db *bbolt.DB
	// Bucket name
	Bucket string
}

// New a BuntDB instance
func New(file string) *BoltDB {
	db, err := bbolt.Open(file, 0666, nil)
	if err != nil {
		panic(err)
	}

	return &BoltDB{db: db, Bucket: "myBucket"}
}

// Has value check by key
func (c *BoltDB) Has(key string) bool {
	return c.Get(key) != nil
}

// Get value by key
func (c *BoltDB) Get(key string) any {
	var val any
	err := c.db.View(func(tx *bbolt.Tx) error {
		b := tx.Bucket([]byte(c.Bucket))
		bs := b.Get([]byte(key))

		if err := c.UnmarshalTo(bs, &val); err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		return nil
	}
	return val
}

// Set value by key
func (c *BoltDB) Set(key string, val any, _ time.Duration) (err error) {
	bts, err := c.MustMarshal(val)
	if err != nil {
		return
	}

	return c.db.Update(func(tx *bbolt.Tx) error {
		b := tx.Bucket([]byte(c.Bucket))
		err := b.Put([]byte(key), bts)
		return err
	})
}

// Del value by key
func (c *BoltDB) Del(key string) error {
	panic("implement me")
}

// GetMulti values by multi key
func (c *BoltDB) GetMulti(keys []string) map[string]any {
	panic("implement me")
}

// SetMulti values by multi key
func (c *BoltDB) SetMulti(values map[string]any, ttl time.Duration) (err error) {
	panic("implement me")
}

// DelMulti values by multi key
func (c *BoltDB) DelMulti(keys []string) error {
	panic("implement me")
}

// Clear all data
func (c *BoltDB) Clear() error {
	panic("implement me")
}

// Close db
func (c *BoltDB) Close() error {
	err := c.db.Sync()
	if err != nil {
		return err
	}

	// do close
	return c.db.Close()
}
