// Package redis is a simple redis cache implement.
// base on the package: https://github.com/gomodule/redigo
package redis

import (
	"errors"
	"fmt"
	"time"

	"github.com/gomodule/redigo/redis"
	"github.com/gookit/cache"
)

const Name = "redigo"

// RedisCache fallback alias
type RedisCache = Redigo

// Redigo driver definition.
// redigo doc link: https://pkg.go.dev/github.com/gomodule/redigo/redis#pkg-examples
type Redigo struct {
	cache.BaseDriver
	// redis connection pool
	pool *redis.Pool
	// info
	url   string
	pwd   string
	dbNum int
}

// New redis cache
func New(url, pwd string, dbNum int) *Redigo {
	rc := &Redigo{
		url: url, pwd: pwd, dbNum: dbNum,
	}

	return rc
}

// Connect create and connect to redis server
func Connect(url, pwd string, dbNum int) *Redigo {
	return New(url, pwd, dbNum).Connect()
}

// Connect to redis server
func (c *Redigo) Connect() *Redigo {
	c.pool = newPool(c.url, c.pwd, c.dbNum)
	c.Logf("connect to server %s db is %d", c.url, c.dbNum)

	return c
}

/*************************************************************
 * methods implements of the gsr.SimpleCacher
 *************************************************************/

// Get value by key
func (c *Redigo) Get(key string) interface{} {
	bts, err := redis.Bytes(c.exec("Get", c.Key(key)))

	return c.Unmarshal(bts, err)
}

// GetAs get cache and unmarshal to ptr
func (c *Redigo) GetAs(key string, ptr interface{}) error {
	bts, err := redis.Bytes(c.exec("Get", c.Key(key)))
	if err != nil {
		return err
	}

	return c.UnmarshalTo(bts, ptr)
}

// Set value by key
func (c *Redigo) Set(key string, val interface{}, ttl time.Duration) (err error) {
	val, err = c.Marshal(val)
	if err != nil {
		return err
	}

	_, err = c.exec("SetEx", c.Key(key), int64(ttl/time.Second), val)
	return
}

// Del value by key
func (c *Redigo) Del(key string) (err error) {
	_, err = c.exec("Del", c.Key(key))
	return
}

// Has cache key
func (c *Redigo) Has(key string) bool {
	// return 0 OR 1
	one, err := redis.Int(c.exec("Exists", c.Key(key)))
	c.SetLastErr(err)

	return one == 1
}

// GetMulti values by keys
func (c *Redigo) GetMulti(keys []string) map[string]interface{} {
	conn := c.pool.Get()
	defer conn.Close()

	args := make([]interface{}, 0, len(keys))
	for _, key := range keys {
		args = append(args, c.Key(key))
	}

	list, err := redis.Values(c.exec("MGet", args...))
	if err != nil {
		c.SetLastErr(err)
		return nil
	}

	values := make(map[string]interface{}, len(keys))
	for i, val := range list {
		values[keys[i]] = val
	}

	return values
}

// SetMulti values
func (c *Redigo) SetMulti(values map[string]interface{}, ttl time.Duration) (err error) {
	conn := c.pool.Get()
	defer conn.Close()

	// open multi
	err = conn.Send("Multi")
	if err != nil {
		return err
	}

	ttlSec := int64(ttl / time.Second)
	for key, val := range values {
		// bs, _ := cache.Marshal(val)
		conn.Send("SetEx", c.Key(key), ttlSec, val)
	}

	// do exec
	_, err = redis.Ints(conn.Do("Exec"))
	return
}

// DelMulti values by keys
func (c *Redigo) DelMulti(keys []string) (err error) {
	args := make([]interface{}, 0, len(keys))
	for _, key := range keys {
		args = append(args, c.Key(key))
	}

	_, err = c.exec("Del", args...)
	return
}

// Close connection
func (c *Redigo) Close() error {
	return c.pool.Close()
}

// Clear all caches
func (c *Redigo) Clear() error {
	_, err := c.exec("FlushDb")
	return err
}

/*************************************************************
 * helper methods
 *************************************************************/

// Pool get
func (c *Redigo) Pool() *redis.Pool {
	return c.pool
}

// String get
func (c *Redigo) String() string {
	pwd := "*"
	if c.IsDebug() {
		pwd = c.pwd
	}

	return fmt.Sprintf("connection info. url: %s, pwd: %s, dbNum: %d", c.url, pwd, c.dbNum)
}

// actually do the redis cmds, args[0] must be the key name.
func (c *Redigo) exec(commandName string, args ...interface{}) (reply interface{}, err error) {
	if len(args) < 1 {
		return nil, errors.New("missing required arguments")
	}

	conn := c.pool.Get()
	defer conn.Close()

	if c.IsDebug() {
		st := time.Now()
		reply, err = conn.Do(commandName, args...)
		c.Logf(
			"operate redis cache. command: %s, key: %v, elapsed time: %.03f\n",
			commandName, args[0], time.Since(st).Seconds()*1000,
		)
		return
	}

	return conn.Do(commandName, args...)
}

// create new pool
func newPool(url, password string, dbNum int) *redis.Pool {
	return &redis.Pool{
		MaxIdle: 5,
		// timeout
		IdleTimeout: 240 * time.Second,
		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial("tcp", url)
			if err != nil {
				return nil, err
			}

			if password != "" {
				_, err := c.Do("AUTH", password)
				if err != nil {
					_ = c.Close()
					return nil, err
				}
			}
			_, _ = c.Do("SELECT", dbNum)
			return c, err
		},
		TestOnBorrow: func(c redis.Conn, t time.Time) error {
			_, err := c.Do("PING")
			return err
		},
	}
}
