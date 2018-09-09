// Package redis is a simple redis cache implement.
// base on the package: github.com/garyburd/redigo
package redis

import (
	"errors"
	"fmt"
	"github.com/garyburd/redigo/redis"
	"log"
	"time"
)

// RedisCache definition.
// redigo doc link: https://godoc.org/github.com/gomodule/redigo/redis
type RedisCache struct {
	Debug bool
	// key prefix
	Prefix string
	Logger *log.Logger
	// info
	url   string
	pwd   string
	dbNum int
	// redis connection pool
	pool *redis.Pool
	// last error
	lastErr error
}

// New redis cache
func New(url, pwd string, dbNum int) *RedisCache {
	rc := &RedisCache{
		url: url, pwd: pwd, dbNum: dbNum,
	}

	return rc
}

// Connect create and connect to redis server
func Connect(url, pwd string, dbNum int) *RedisCache {
	return New(url, pwd, dbNum).Connect()
}

// Connect to redis server
func (c *RedisCache) Connect() *RedisCache {
	c.pool = newPool(c.url, c.pwd, c.dbNum)
	c.logf("connect to server %s db is %d", c.url, c.dbNum)

	return c
}

// Get value by key
func (c *RedisCache) Get(key string) interface{} {
	val, err := c.exec("Get", c.Key(key))
	if err != nil {
		c.lastErr = err
		return nil
	}

	return val
}

// Set value by key
func (c *RedisCache) Set(key string, val interface{}, ttl time.Duration) (err error) {
	// bs, _ := cache.Marshal(val)
	// _, err = c.exec("SetEx", c.Key(key), int64(ttl/time.Second), bs)
	_, err = c.exec("SetEx", c.Key(key), int64(ttl/time.Second), val)
	return
}

// Del value by key
func (c *RedisCache) Del(key string) (err error) {
	_, err = c.exec("Del", c.Key(key))
	if err != nil {
		c.logf("redis error: %s\n", err.Error())
		c.lastErr = err
	}

	return
}

// Has cache key
func (c *RedisCache) Has(key string) bool {
	// return 0 OR 1
	one, err := redis.Int(c.exec("Exists", c.Key(key)))
	if err != nil {
		c.logf("redis error: %s\n", err.Error())
		c.lastErr = err
	}

	return one == 1
}

// GetMulti values by keys
func (c *RedisCache) GetMulti(keys []string) map[string]interface{} {
	conn := c.pool.Get()
	defer conn.Close()

	var args []interface{}
	for _, key := range keys {
		args = append(args, c.Key(key))
	}

	list, err := redis.Values(conn.Do("MGet", args...))
	if err != nil {
		c.lastErr = err
		return nil
	}

	values := make(map[string]interface{}, len(keys))
	for i, val := range list {
		values[keys[i]] = val
	}

	return values
}

// SetMulti values
func (c *RedisCache) SetMulti(values map[string]interface{}, ttl time.Duration) (err error) {
	conn := c.pool.Get()
	defer conn.Close()

	// open multi
	conn.Send("Multi")
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
func (c *RedisCache) DelMulti(keys []string) (err error) {
	conn := c.pool.Get()
	defer conn.Close()

	var args []interface{}
	for _, key := range keys {
		args = append(args, c.Key(key))
	}

	_, err = conn.Do("Del", args...)
	return
}

// Clear all caches
func (c *RedisCache) Clear() error {
	conn := c.pool.Get()
	defer conn.Close()
	_, err := conn.Do("FlushDb")
	c.lastErr = err
	return err
}

// String get
func (c *RedisCache) String() string {
	pwd := "*"
	if c.Debug {
		pwd = c.pwd
	}

	return fmt.Sprintf("connection info. url: %s, pwd: %s, dbNum: %d", c.url, pwd, c.dbNum)
}

/*************************************************************
 * helper methods
 *************************************************************/

// Driver object get
func (c *RedisCache) Pool() *redis.Pool {
	return c.pool
}

// Key build
func (c *RedisCache) Key(key string) string {
	if c.Prefix != "" {
		return fmt.Sprintf("%s:%s", c.Prefix, key)
	}

	return key
}

// LastErr get
func (c *RedisCache) LastErr() error {
	return c.lastErr
}

// actually do the redis cmds, args[0] must be the key name.
func (c *RedisCache) exec(commandName string, args ...interface{}) (reply interface{}, err error) {
	if len(args) < 1 {
		return nil, errors.New("missing required arguments")
	}

	conn := c.pool.Get()
	defer conn.Close()

	if c.Debug {
		st := time.Now()
		reply, err = conn.Do(commandName, args...)
		c.logf(
			"operate redis cache. command: %s, key: %v, elapsed time: %.03f\n",
			commandName, args[0], time.Since(st).Seconds()*1000,
		)
		return
	}

	return conn.Do(commandName, args...)
}

func (c *RedisCache) logf(format string, v ...interface{}) {
	if c.Logger != nil {
		c.Logger.Printf(format, v...)
	}
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
					c.Close()
					return nil, err
				}
			}
			c.Do("SELECT", dbNum)
			return c, err
		},
		TestOnBorrow: func(c redis.Conn, t time.Time) error {
			_, err := c.Do("PING")
			return err
		},
	}
}
