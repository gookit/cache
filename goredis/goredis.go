// Package goredis is a simple redis cache implement.
// base on the package: github.com/go-redis/redis
package goredis

import (
	"context"
	"fmt"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/gookit/cache"
)

// Name driver name
const Name = "goredis"

// CtxForExec Ctx for exec command
var CtxForExec = context.Background()

// GoRedis struct
type GoRedis struct {
	cache.BaseDriver
	// client
	rdb *redis.Client
	// config
	url   string
	pwd   string
	dbNum int
}

// Connect create and connect to redis server
func Connect(url, pwd string, dbNum int) *GoRedis {
	return New(url, pwd, dbNum).Connect()
}

// New redis cache
func New(url, pwd string, dbNum int) *GoRedis {
	rc := &GoRedis{
		url: url, pwd: pwd, dbNum: dbNum,
	}

	return rc
}

// String get
func (c *GoRedis) String() string {
	pwd := "*"
	if c.IsDebug() {
		pwd = c.pwd
	}

	return fmt.Sprintf("connection info. url: %s, pwd: %s, dbNum: %d", c.url, pwd, c.dbNum)
}

// Connect to redis server
func (c *GoRedis) Connect() *GoRedis {
	c.rdb = redis.NewClient(&redis.Options{
		Addr:     c.url,
		Password: c.pwd, // no password set
		DB:       c.dbNum,  // use default DB
	})
	c.Logf("connect to server %s db is %d", c.url, c.dbNum)

	return c
}

/*************************************************************
 * methods implements of the gsr.SimpleCacher
 *************************************************************/

// Close connection
func (c *GoRedis) Close() error {
	return c.rdb.Close()
}

// Clear all caches
func (c *GoRedis) Clear() error {
	return c.rdb.FlushDB(CtxForExec).Err()
}

// Has cache key
func (c *GoRedis) Has(key string) bool {
	n, err := c.rdb.Exists(CtxForExec, c.Key(key)).Result()
	if err != nil {
		c.SetLastErr(err)
		return false
	}

	return n == 1
}

// Get cache by key
func (c *GoRedis) Get(key string) interface{} {
	bts, err := c.rdb.Get(CtxForExec, c.Key(key)).Bytes()

	return c.Unmarshal(bts, err)
}

// GetAs get cache and unmarshal to ptr
func (c *GoRedis) GetAs(key string, ptr interface{}) error {
	bts, err := c.rdb.Get(CtxForExec, c.Key(key)).Bytes()
	if err != nil {
		return err
	}

	return c.UnmarshalTo(bts, ptr)
}

// Set cache by key
func (c *GoRedis) Set(key string, val interface{}, ttl time.Duration) (err error) {
	val, err = c.Marshal(val)
	if err != nil {
		return err
	}

	return c.rdb.SetEX(CtxForExec, c.Key(key), val, ttl).Err()
}

// Del cache by key
func (c *GoRedis) Del(key string) error {
	return c.rdb.Del(CtxForExec, c.Key(key)).Err()
}

// GetMulti cache by keys
func (c *GoRedis) GetMulti(keys []string) map[string]interface{} {
	panic("implement me")
}

// SetMulti cache by keys
func (c *GoRedis) SetMulti(values map[string]interface{}, ttl time.Duration) (err error) {
	panic("implement me")
}

// DelMulti cache by keys
func (c *GoRedis) DelMulti(keys []string) error {
	cks := make([]string,0, len(keys))
	for _, key := range keys {
		cks = append(cks, c.Key(key))
	}

	return c.rdb.Del(CtxForExec, cks...).Err()
}
