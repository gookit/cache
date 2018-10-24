# cache

[![GoDoc](https://godoc.org/github.com/gookit/cache?status.svg)](https://godoc.org/github.com/gookit/cache)
[![Go Report Card](https://goreportcard.com/badge/github.com/gookit/cache)](https://goreportcard.com/report/github.com/gookit/cache)

> **[中文说明](README_cn.md)**

Generic cache use and cache manager for golang.

Supported Drivers:

- file internal driver
- memory internal driver
- redis powered by `github.com/gomodule/redigo`
- memCached powered by `github.com/bradfitz/gomemcache`
- buntdb powered by `github.com/tidwall/buntdb`
- boltdb powered by `github.com/etcd-io/bbolt`

## GoDoc

- [doc on gowalker](https://gowalker.org/github.com/gookit/cache)
- [godoc for gopkg](https://godoc.org/gopkg.in/gookit/cache.v1)
- [godoc for github](https://godoc.org/github.com/gookit/cache)

## Interface

```go
// Cache interface definition
type Cache interface {
	// basic op
	Has(key string) bool
	Get(key string) interface{}
	Set(key string, val interface{}, ttl time.Duration) (err error)
	Del(key string) error
	// multi op
	GetMulti(keys []string) map[string]interface{}
	SetMulti(values map[string]interface{}, ttl time.Duration) (err error)
	DelMulti(keys []string) error
	// clear
	Clear() error
}
```

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

## License

**MIT**
