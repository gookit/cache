package cache_test

import (
	"math/rand"
	"strconv"
	"testing"
	"time"

	"github.com/gookit/cache"
	"github.com/gookit/goutil/dump"
	"github.com/stretchr/testify/assert"
)

func TestNewMemoryCache(t *testing.T) {
	is := assert.New(t)
	c := cache.NewMemoryCache()

	key := "key"
	is.False(c.Has(key))

	err := c.Set(key, "value", cache.Seconds3)
	is.NoError(err)
	is.True(c.Has(key))

	val := c.Get(key)
	is.Equal("value", val)

	// del
	err = c.Del(key)
	is.NoError(err)
	is.False(c.Has(key))
}

type user struct {
	Age int
	Name string
}

func TestMemoryCache_object(t *testing.T) {
	is := assert.New(t)

	b1 := user {
		Age: 1,
		Name: "inhere",
	}

	c := cache.NewMemoryCache()

	key := randomKey()
	t.Log("cache key:", key)
	is.False(c.Has(key))

	err := c.Set(key, b1, cache.Seconds3)
	is.NoError(err)
	is.True(c.Has(key))

	b2 := c.Get(key).(user)
	dump.P(b2)
	is.Equal("inhere", b2.Name)
}

func TestMemoryCache_expired(t *testing.T) {
	is := assert.New(t)
	c := cache.NewMemoryCache()

	key := "key"
	is.False(c.Has(key))

	err := c.Set(key, "value", cache.Seconds1)
	is.NoError(err)
	is.Equal("value", c.Get(key))

	time.Sleep(cache.Seconds2)

	is.Nil(c.Get(key))
}

func TestNewFileCache(t *testing.T) {
	is := assert.New(t)
	c := cache.NewFileCache("./testdata")

	key := "key"
	is.False(c.Has(key))

	// set
	err := c.Set(key, "cache value", cache.OneMinutes)
	is.NoError(err)
	is.True(c.Has(key))

	err = c.Set("key2", "cache value2", cache.TwoMinutes)
	is.NoError(err)

	// get
	val := c.Get(key)
	is.Equal("cache value", val)

	// del
	err = c.Del(key)
	is.NoError(err)
	is.False(c.Has(key))
}

func TestFileCache_object(t *testing.T) {
	is := assert.New(t)
	c := cache.NewFileCache("./testdata")
	c.WithOptions(cache.WithEncode(true))

	b1 := user {
		Age: 12,
		Name: "inhere",
	}

	key := randomKey()
	dump.P("cache key:", c.Key(key))

	err := c.Set(key, b1, cache.Seconds3)
	is.NoError(err)
	is.True(c.Has(key))

	val := c.Get(key)
	dump.P("cache get:", val)

	// val2 := c.GetAs()
	// dump.P("cache get:", val)
}

func randomKey() string {
	rand.Seed(time.Now().UnixNano())
	return "k" + time.Now().Format("20060102") + strconv.Itoa(rand.Intn(999))
}
