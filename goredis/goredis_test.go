package goredis_test

import (
	"fmt"
	"math/rand"
	"strconv"
	"testing"
	"time"

	"github.com/gookit/cache"
	"github.com/gookit/cache/goredis"
	"github.com/gookit/goutil/dump"
	"github.com/stretchr/testify/assert"
)

func Example() {
	// init driver
	c := goredis.Connect("127.0.0.1:6379", "", 0)

	// set
	_ = c.Set("name", "cache value", 60)

	// get
	val := c.Get("name")

	// del
	_ = c.Del("name")

	// get: "cache value"
	fmt.Print(val)
}

var c *goredis.GoRedis

func getC() *goredis.GoRedis {
	if c != nil {
		return c
	}

	c = goredis.New("127.0.0.1:6379", "", 0).Connect()
	c.WithOptions(cache.WithPrefix("gr"), cache.WithEncode(true))

	return c
}

func TestGoRedis_basic(t *testing.T) {
	c := getC()

	key := randomKey()
	t.Log("cache key", key)
	assert.False(t, c.Has(key))

	err := c.Set(key, "value", cache.Seconds3)
	assert.NoError(t, err)

	assert.True(t, c.Has(key))
	assert.Equal(t, "value", c.Get(key).(string))

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
	return time.Now().Format("20060102") + strconv.Itoa(rand.Intn(999))
}
