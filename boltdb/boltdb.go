// Package bolt use the github.com/etcd-io/bbolt as cache driver
package boltdb

import (
	"bytes"
	"encoding/gob"
	"time"

	"github.com/etcd-io/bbolt"
)

// BoltDB definition
type BoltDB struct {
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
	panic("implement me")
}

// Get value by key
func (c *BoltDB) Get(key string) interface{} {
	var val interface{}
	err := c.db.View(func(tx *bbolt.Tx) error {
		b := tx.Bucket([]byte(c.Bucket))
		bs := b.Get([]byte(key))

		buf := bytes.NewBuffer(bs)
		dec := gob.NewDecoder(buf)
		if err := dec.Decode(val); err != nil {
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
func (c *BoltDB) Set(key string, val interface{}, _ time.Duration) (err error) {
	buf := new(bytes.Buffer)
	enc := gob.NewEncoder(buf)
	if err = enc.Encode(val); err != nil {
		return
	}

	return c.db.Update(func(tx *bbolt.Tx) error {
		b := tx.Bucket([]byte(c.Bucket))
		err := b.Put([]byte(key), buf.Bytes())
		return err
	})
}

// Del value by key
func (c *BoltDB) Del(key string) error {
	panic("implement me")
}

// GetMulti values by multi key
func (c *BoltDB) GetMulti(keys []string) map[string]interface{} {
	panic("implement me")
}

// SetMulti values by multi key
func (c *BoltDB) SetMulti(values map[string]interface{}, ttl time.Duration) (err error) {
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
	return c.db.Close()
}
