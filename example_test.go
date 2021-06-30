package cache_test

import (
	"fmt"

	"github.com/gookit/cache"
	"github.com/gookit/cache/goredis"
	"github.com/gookit/cache/redis"
	"github.com/gookit/goutil/dump"
)

func Example() {
	// register some cache driver
	cache.Register(cache.DvrFile, cache.NewFileCache(""))
	cache.Register(cache.DvrMemory, cache.NewMemoryCache())
	cache.Register(redis.Name, redis.Connect("127.0.0.1:6379", "", 0))
	cache.Register(goredis.Name, goredis.Connect("127.0.0.1:6379", "", 0))

	// setting default driver name
	cache.DefaultUse(goredis.Name)

	// quick use.(it is default driver)
	//
	// set
	_ = cache.Set("name", "cache value", cache.TwoMinutes)
	// get
	val := cache.Get("name")
	// del
	_ = cache.Del("name")

	// get: "cache value"
	fmt.Print(val)

	// More ...
	// fc := cache.GetCache(DvrFile)
	// fc.Set("key", "value", 10)
	// fc.Get("key")
}

func ExampleMemoryCache() {
	c := cache.NewMemoryCache()
	key := "name"

	// set
	c.Set(key, "cache value", cache.TwoMinutes)
	fmt.Println(c.Has(key), c.Count())

	// get
	val := c.Get(key)
	fmt.Println(val)

	// del
	c.Del(key)
	fmt.Println(c.Has(key), c.Count())

	// Output:
	// true 1
	// cache value
	// false 0
}

func ExampleFileCache() {
	c := cache.NewFileCache("./testdata")
	key := "name"

	// set
	c.Set(key, "cache value", cache.TwoMinutes)
	fmt.Println(c.Has(key))

	// get
	val := c.Get(key)
	fmt.Println(val)

	// del
	c.Del(key)
	fmt.Println(c.Has(key))

	// Output:
	// true
	// cache value
	// false
}

func Example_withOptions() {
	gords := goredis.Connect("127.0.0.1:6379", "", 0)
	gords.WithOptions(cache.WithPrefix("cache_"), cache.WithEncode(true))

	// register
	cache.Register(goredis.Name, gords)

	// set
	// real key is: "cache_name"
	cache.Set("name", "cache value", cache.TwoMinutes)

	// get: "cache value"
	val := cache.Get("name")

	dump.P(val)
}