package goredis_test

import (
	"fmt"
	"math/rand"
	"strconv"
	"testing"
	"time"

	"github.com/gookit/cache"
	"github.com/gookit/cache/goredis"
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

func TestGoRedis_Get(t *testing.T) {
	c := goredis.New("127.0.0.1:6379", "", 0).Connect()
	c.Prefix = "gr_"

	key := randomKey()
	assert.False(t, c.Has(key))
	err := c.Set(key, "value", cache.Seconds3)
	assert.NoError(t, err)
	assert.True(t, c.Has(key))
	assert.Equal(t, "value", c.Get(key).(string))
}

func randomKey() string {
	return time.Now().Format("20060102") + strconv.Itoa(rand.Intn(999))
}
