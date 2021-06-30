// Package buntdb use the github.com/tidwall/buntdb as cache driver
package buntdb

import (
	"time"

	"github.com/gookit/cache"
	"github.com/tidwall/buntdb"
)

// Memory open a file that does not persist to disk.
const Memory = ":memory:"
// Name driver name
const Name = "buntDB"

// BuntDB definition.
type BuntDB struct {
	cache.BaseDriver
	// db file path. eg "path/to/my.db"
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
	if file == "" { // use memory
		file = Memory
	}

	db, err := buntdb.Open(file)
	if err != nil {
		panic(err)
	}

	return &BuntDB{
		db: db,
	}
}

// Db get
func (c *BuntDB) Db() *buntdb.DB {
	return c.db
}

// Has key
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
	err := c.db.View(func(tx *buntdb.Tx) error {
		str, err := tx.Get(key, false)
		if err != nil {
			return err
		}

		return c.UnmarshalTo([]byte(str), &val)
	})

	if err != nil {
		return nil
	}
	return val
}

// Set value by key
func (c *BuntDB) Set(key string, val interface{}, ttl time.Duration) (err error) {
	bts, err := c.MustMarshal(val)
	if err != nil {
		return err
	}

	return c.db.Update(func(tx *buntdb.Tx) (err error) {
		opt := &buntdb.SetOptions{}
		if ttl > 0 {
			opt.TTL = ttl
			opt.Expires = true
		}

		_, _, err = tx.Set(key, string(bts), opt)
		return err
	})
}

// Del value by key
func (c *BuntDB) Del(key string) error {
	return c.db.Update(func(tx *buntdb.Tx) error {
		_, err := tx.Delete(key)
		return err
	})
}

// GetMulti values by multi key
func (c *BuntDB) GetMulti(keys []string) map[string]interface{} {
	results := make(map[string]interface{}, len(keys))
	err := c.db.View(func(tx *buntdb.Tx) error {
		for _, key := range keys {
			str, err := tx.Get(key, false)
			if err != nil {
				return err
			}

			var val interface{}
			err = c.UnmarshalTo([]byte(str), &val)
			if err != nil {
				return err
			}

			results[key] = val
		}
		return nil
	})

	if err != nil {
		return nil
	}

	return results
}

// SetMulti values by multi key
func (c *BuntDB) SetMulti(values map[string]interface{}, ttl time.Duration) (err error) {
	return c.db.Update(func(tx *buntdb.Tx) (err error) {
		opt := &buntdb.SetOptions{}
		if ttl > 0 {
			opt.TTL = ttl
			opt.Expires = true
		}

		for key, val := range values {
			bts, err := c.MustMarshal(val)
			if err != nil {
				return err
			}

			_, _, err = tx.Set(key, string(bts), opt)
			if err != nil {
				return err
			}
		}

		return nil
	})
}

// DelMulti values by multi key
func (c *BuntDB) DelMulti(keys []string) error {
	return c.db.Update(func(tx *buntdb.Tx) (err error) {
		for _, k := range keys {
			if _, err = tx.Delete(k); err != nil {
				return err
			}
		}

		return nil
	})
}

// Clear all cache data
func (c *BuntDB) Clear() error {
	return c.db.Update(func(tx *buntdb.Tx) error {
		return tx.DeleteAll()
	})
}

// Close cache db
func (c *BuntDB) Close() error {
	return c.db.Close()
}
