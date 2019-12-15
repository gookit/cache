package cache

import (
	"fmt"

	"github.com/gookit/cache/redis"
)

func Example() {
	// register some cache driver
	Register(DvrFile, NewFileCache(""))
	Register(DvrMemory, NewMemoryCache())
	Register(DvrRedis, redis.Connect("127.0.0.1:6379", "", 0))

	// setting default driver name
	SetDefName(DvrRedis)

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
}

func ExampleMemoryCache() {
	c := NewMemoryCache()
	key := "name"

	// set
	c.Set(key, "cache value", TwoMinutes)
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
	c := NewFileCache("./testdata")
	key := "name"

	// set
	c.Set(key, "cache value", TwoMinutes)
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
