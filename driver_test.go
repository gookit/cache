package cache_test

import (
	"testing"

	"github.com/gookit/cache"
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