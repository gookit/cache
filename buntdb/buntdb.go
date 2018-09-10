// Package buntdb use the github.com/tidwall/buntdb as cache driver
package buntdb

import (
	"github.com/tidwall/buntdb"
	"time"
)

// Memory open a file that does not persist to disk.
const Memory = ":memory:"

// BuntDB definition.
type BuntDB struct {
	// db file path. 
	file string
	// db instance
	db *buntdb.DB
}

// NewMemory new a memory db
func NewMemory() *BuntDB {
	return New(Memory)
}

// New a BuntDB instance
func New(file string) *BuntDB {
	db, err := buntdb.Open(file)
	if err != nil {
		panic(err)
	}

	return &BuntDB{
		db: db,
	}
}

func (c *BuntDB) Has(key string) bool {
	has := false
	err := c.db.View(func(tx *buntdb.Tx) error {
		val, err := tx.Get(key, false)
		has = val != ""
		return err
	})

	if err != nil {
		has = false
	}

	return has
}

// Get value by key
func (c *BuntDB) Get(key string) interface{} {
	var val interface{}
	err := c.db.View(func(tx *buntdb.Tx) (err error) {
		val, err = tx.Get(key, false)
		return err
	})

	if err != nil {
		return nil
	}

	return val
}

// Get value by key
func (c *BuntDB) Set(key string, val interface{}, ttl time.Duration) (err error) {
	panic("implement me")
}

func (c *BuntDB) Del(key string) error {
	panic("implement me")
}

func (c *BuntDB) GetMulti(keys []string) map[string]interface{} {
	panic("implement me")
}

func (c *BuntDB) SetMulti(values map[string]interface{}, ttl time.Duration) (err error) {
	panic("implement me")
}

func (c *BuntDB) DelMulti(keys []string) error {
	panic("implement me")
}

// Clear cache data
func (c *BuntDB) Clear() error {
	return c.db.Close()
}

// Close cache db
func (c *BuntDB) Close() error {
	return c.db.Close()
}
