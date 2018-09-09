# cache

Generic cache use and cache manager for golang.

Drivers:

- file
- redis
- memory
- memCached

## Godoc

- [doc on gowalker](https://gowalker.org/github.com/gookit/cache)
- [godoc for gopkg](https://godoc.org/gopkg.in/gookit/cache.v1)
- [godoc for github](https://godoc.org/github.com/gookit/cache)

## Usage

```go
package main

import (
	"fmt"
	"github.com/gookit/cache"
	"github.com/gookit/cache/redis"
)

func main() {
	// register one(or some) cache driver
	cache.Register(cache.DvrFile, cache.NewFileCache(""))
	cache.Register(cache.DvrMemory, cache.NewMemoryCache())
	cache.Register(cache.DvrRedis, redis.Connect("127.0.0.1:6379", "", 0))
	
	// setting default driver name
	cache.SetDefName(cache.DvrRedis)

	// quick use.(it is default driver)
	//
	// set
	cache.Set("name", "cache value", cache.TwoMinutes)
	// get
	val := cache.Get("name")
	// del
	cache.Del("name")

	// get: "cache value"
	fmt.Print(val)
}
```

## Interface

```go
// CacheFace interface definition
type CacheFace interface {
	// basic op
	Has(key string) bool
	Get(key string) interface{}
	Set(key string, val interface{}, ttl time.Duration) error
	Del(key string) error
	// multi op
	GetMulti(keys []string) []interface{}
	SetMulti(mv map[string]interface{}, ttl time.Duration) error
	DelMulti(keys []string) error
	// clear
	Clear() error
}
```

## License

**MIT**
