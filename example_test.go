package cache

import (
	"fmt"
	"time"

	"github.com/gookit/cache/redis"
)

func Example() {
	// register some cache driver
	Register(DvrFile, NewFileCache(""))
	Register(DvrMemory, NewMemoryCache())
	Register(DvrRedis, redis.Connect("127.0.0.1:6379", "", 0))

	// setting default driver name
	DefaultUse(DvrRedis)

	// quick use.(it is default driver)
	//
	// set
	_ = Set("name", "cache value", TwoMinutes)
	// get
	val := Get("name")
	// del
	_ = Del("name")

	// get: "cache value"
	fmt.Print(val)

	// More ...
	// fc := GetCache(DvrFile)
	// fc.Set("key", "value", 10)
	// fc.Get("key")
}

func ExampleMemoryCache() {
	c := NewMemoryCache()
	key := "name"

	// set
	c.Set(key, "cache value", TwoSeconds)
	fmt.Println(c.Has(key), c.Count())

	// get
	val := c.Get(key)
	fmt.Println(val)

	time.Sleep(TwoSeconds)

	// get expired
	val2 := c.Get(key)
	fmt.Println(val2)

	// del
	c.Del(key)
	fmt.Println(c.Has(key), c.Count())

	// Output:
	// true 1
	// cache value
	// <nil>
	// false 0
}

func ExampleFileCache() {
	c := NewFileCache("./testdata")
	key := "name"

	// set
	c.Set(key, "cache value", TwoSeconds)
	fmt.Println(c.Has(key))

	// get
	val := c.Get(key)
	fmt.Println(val)

	time.Sleep(TwoSeconds)

	// get expired
	val2 := c.Get(key)
	fmt.Println(val2)

	// del
	c.Del(key)
	fmt.Println(c.Has(key))

	// Output:
	// true
	// cache value
	// <nil>
	// false
}
