package redis

import (
	"errors"
	"fmt"
	"github.com/garyburd/redigo/redis"
	"github.com/gookit/cache"
	"log"
	"time"
)

// RedisCache definition.
// redigo doc link: https://godoc.org/github.com/gomodule/redigo/redis
type RedisCache struct {
	Debug bool
	// info
	url   string
	pwd   string
	dbNum int
	// key prefix
	prefix string
	// redis connection pool
	pool *redis.Pool
	// last error
	lastErr error
	logger  *log.Logger
}

// New redis cache
func New(url, pwd string, dbNum int) *RedisCache {
	rc := &RedisCache{
		url: url, pwd: pwd, dbNum: dbNum,
	}

	rc.pool = newPool(url, pwd, dbNum)
	return rc
}

// Driver object get
func (c *RedisCache) Pool() *redis.Pool {
	return c.pool
}

// Has cache key
func (c *RedisCache) Has(key string) bool {
	// return 0 OR 1
	one, err := redis.Int(c.exec("Exists", key))
	if err != nil {
		c.logger.Println("redis error: " + err.Error())
		c.lastErr = err
	}

	return one == 1
}

// Get value by key
func (c *RedisCache) Get(key string) interface{} {
	val, err := c.exec("Get", key)
	if err != nil {
		c.lastErr = err
		return nil
	}

	return val
}

// MapValue get cache value and map to a struct
func (c *RedisCache) MapValue(key string, ptr interface{}) error {
	val, err := c.exec("Get", key)
	if err == nil {
		// val must convert to byte
		return cache.Unmarshal(val.([]byte), ptr)
	}

	c.lastErr = err
	return err
}

// Set value by key
func (c *RedisCache) Set(key string, v interface{}, ttl time.Duration) (err error) {
	bs, _ := cache.Marshal(v)
	_, err = c.exec("SetEx", key, int64(ttl/time.Second), bs)
	return
}

// Del value by key
func (c *RedisCache) Del(key string) (err error) {
	_, err = c.exec("Del", key)
	if err != nil {
		c.logger.Println("redis error: " + err.Error())
		c.lastErr = err
	}

	return
}

// GetMulti values by keys
func (c *RedisCache) GetMulti(keys []string) []interface{} {
	conn := c.pool.Get()
	defer conn.Close()

	var args []interface{}
	for _, key := range keys {
		args = append(args, key)
	}

	values, err := redis.Values(conn.Do("MGet", args...))
	if err != nil {
		c.lastErr = err
		return nil
	}

	return values
}

// SetMulti values
func (c *RedisCache) SetMulti(values map[string]interface{}, ttl time.Duration) (err error) {
	conn := c.pool.Get()
	defer conn.Close()

	ttlSec := int64(ttl / time.Second)
	// open multi
	conn.Send("Multi")

	for key, val := range values {
		bs, _ := cache.Marshal(val)
		conn.Send("SetEx", key, ttlSec, bs)
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
		args = append(args, key)
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

// LastErr get
func (c *RedisCache) LastErr() error {
	return c.lastErr
}

// String get
func (c *RedisCache) String() string {
	pwd := "*"
	if c.Debug {
		pwd = c.pwd
	}

	return fmt.Sprintf("connection info. url: %s, pwd: %s, dbNum: %d", c.url, pwd, c.dbNum)
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
		c.logger.Printf(
			"operate redis cache. command: %s, key: %v, elapsed time: %.03f\n",
			commandName, args[0], time.Since(st).Seconds()*1000,
		)
		return
	}

	return conn.Do(commandName, args...)
}
