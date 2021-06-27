package redis_test

import (
	"fmt"
	"math/rand"
	"strconv"
	"testing"
	"time"

	"github.com/gookit/cache"
	"github.com/gookit/cache/redis"
	"github.com/gookit/goutil/dump"
	"github.com/stretchr/testify/assert"
)

func Example() {
	// init driver
	c := redis.Connect("127.0.0.1:6379", "", 0)

	// set
	_ = c.Set("name", "cache value", 60)

	// get
	val := c.Get("name")

	// del
	_ = c.Del("name")

	// get: "cache value"
	fmt.Print(val)
}

var c *redis.Redigo

func getC() *redis.Redigo {
	if c != nil {
		return c
	}

	c = redis.New("127.0.0.1:6379", "", 0).Connect()
	c.WithOptions(cache.WithPrefix("rdg"), cache.WithEncode(true))
	return c
}

func TestRedigo_basic(t *testing.T) {
	c := getC()

	key := randomKey()
	t.Log("cache key", c.Key(key))
	assert.False(t, c.Has(key))

	err := c.Set(key, "value", cache.Seconds3)
	assert.NoError(t, err)

	assert.True(t, c.Has(key))
	assert.Equal(t, "value", c.Get(key))

	err = c.Del(key)
	assert.NoError(t, err)
	assert.False(t, c.Has(key))
}

type user struct {
	Age int
	Name string
}

func TestRedigo_object(t *testing.T) {
	c := getC()
	b1 := user {
		Age: 12,
		Name: "inhere",
	}

	key := randomKey()
	t.Log("cache key", c.Key(key))
	assert.False(t, c.Has(key))

	err := c.Set(key, b1, cache.Seconds3)
	assert.NoError(t, err)
	assert.True(t, c.Has(key))

	v := c.Get(key)
	assert.NotEmpty(t, v)

	// dump.P(v.(string))
	dump.P(v)

	err = c.Del(key)
	assert.NoError(t, err)
	assert.False(t, c.Has(key))

	assert.Empty(t, c.Get(key))
}

func randomKey() string {
	return "k" + time.Now().Format("20060102") + strconv.Itoa(rand.Intn(999))
}
