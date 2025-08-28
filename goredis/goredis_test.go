package goredis_test

import (
	"fmt"
	"testing"

	"github.com/gookit/cache"
	"github.com/gookit/cache/goredis"
	"github.com/gookit/goutil/dump"
	"github.com/gookit/goutil/strutil"
	"github.com/gookit/goutil/testutil/assert"
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

	key := strutil.RandomCharsV2(12)
	dump.P("cache key: " + c.Key(key))

	assert.False(t, c.Has(key))

	err := c.Set(key, "value", cache.Forever)
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
	u1 := user {
		Age: 12,
		Name: "inhere",
	}

	key := strutil.RandomCharsV2(12)
	dump.P("cache key: " + c.Key(key))

	assert.False(t, c.Has(key))

	err := c.Set(key, u1, cache.Seconds3)
	assert.NoError(t, err)
	assert.True(t, c.Has(key))

	v := c.Get(key)
	assert.NotEmpty(t, v)

	dump.P(v)

	u2 := user{}
	err = c.GetAs(key, &u2)
	assert.NoError(t, err)
	dump.P(u2)
	assert.Equal(t, "inhere", u2.Name)

	err = c.Del(key)
	assert.NoError(t, err)
	assert.False(t, c.Has(key))
	assert.Empty(t, c.Get(key))
}
