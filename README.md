# Cache

![GitHub go.mod Go version](https://img.shields.io/github/go-mod/go-version/gookit/cache?style=flat-square)
[![GoDoc](https://godoc.org/github.com/gookit/cache?status.svg)](https://pkg.go.dev/github.com/gookit/cache)
[![Go Report Card](https://goreportcard.com/badge/github.com/gookit/cache)](https://goreportcard.com/report/github.com/gookit/cache)
[![Actions Status](https://github.com/gookit/cache/workflows/Unit-Tests/badge.svg)](https://github.com/gookit/cache/actions)

> **[中文说明](README_cn.md)**

Generic cache use and cache manager for golang. 

Provide a unified usage API by packaging various commonly used drivers.

> All cache driver implemented the `cache.Cache` interface. So, You can add any custom driver.

**Packaged Drivers:**

- `goredis` https://github.com/go-redis/redis
- `redis` https://github.com/gomodule/redigo
- `memcached` https://github.com/bradfitz/gomemcache
- `buntdb` https://github.com/tidwall/buntdb
- `boltdb`  https://github.com/etcd-io/bbolt
- `badger` https://github.com/dgraph-io/badger
- `nutsdb` https://github.com/xujiajun/nutsdb
- `goleveldb` https://github.com/syndtr/goleveldb
- `gcache` https://github.com/bluele/gcache
- `gocache` https://github.com/patrickmn/go-cache
- `bigcache` https://github.com/allegro/bigcache

Internal:

- file internal driver [driver_file.go](driver_file.go)
- memory internal driver [driver_memory.go](driver_memory.go)

> Notice: The built-in implementation is relatively simple and is not recommended for production environments;
> the production environment recommends using the third-party drivers listed above.

## GoDoc

- [doc on gowalker](https://gowalker.org/github.com/gookit/cache)
- [godoc for gopkg](https://pkg.go.dev/gopkg.in/gookit/cache.v1)
- [godoc for github](https://pkg.go.dev/github.com/gookit/cache)

## Install

The package supports 3 last Go versions and requires a Go version with modules support.

```bash
go get github.com/gookit/cache
```

## Cache Interface

All cache driver implemented the `cache.Cache` interface. So, You can add any custom driver.

```go
type Cache interface {
	// basic operation
	Has(key string) bool
	Get(key string) any
	Set(key string, val any, ttl time.Duration) (err error)
	Del(key string) error
	// multi operation
	GetMulti(keys []string) map[string]any
	SetMulti(values map[string]any, ttl time.Duration) (err error)
	DelMulti(keys []string) error
	// clear and close
	Clear() error
	Close() error
}
```

## Usage

```go
package main

import (
	"fmt"

	"github.com/gookit/cache"
	"github.com/gookit/cache/gcache"
	"github.com/gookit/cache/gocache"
	"github.com/gookit/cache/goredis"
	"github.com/gookit/cache/redis"
)

func main() {
	// register one(or some) cache driver
	cache.Register(cache.DvrFile, cache.NewFileCache(""))
	// cache.Register(cache.DvrMemory, cache.NewMemoryCache())
	cache.Register(gcache.Name, gcache.New(1000))
	cache.Register(gocache.Name, gocache.NewGoCache(cache.OneDay, cache.FiveMinutes))
	cache.Register(redis.Name, redis.Connect("127.0.0.1:6379", "", 0))
	cache.Register(goredis.Name, goredis.Connect("127.0.0.1:6379", "", 0))

	// setting default driver name
	cache.DefaultUse(gocache.Name)

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

	// More ...
	// fc := cache.Driver(gcache.Name)
	// fc.Set("key", "value", 10)
	// fc.Get("key")
}
```

## With Options

```go
gords := goredis.Connect("127.0.0.1:6379", "", 0)
gords.WithOptions(cache.WithPrefix("cache_"), cache.WithEncode(true))

cache.Register(goredis.Name, gords)

// set
// real key is: "cache_name"
cache.Set("name", "cache value", cache.TwoMinutes)

// get: "cache value"
val := cache.Get("name")
```

## Gookit packages

- [gookit/ini](https://github.com/gookit/ini) Go config management, use INI files
- [gookit/rux](https://github.com/gookit/rux) Simple and fast request router for golang HTTP
- [gookit/gcli](https://github.com/gookit/gcli) build CLI application, tool library, running CLI commands
- [gookit/slog](https://github.com/gookit/slog) Lightweight, extensible, configurable logging library written in Go
- [gookit/event](https://github.com/gookit/event) Lightweight event manager and dispatcher implements by Go
- [gookit/cache](https://github.com/gookit/cache) Provide a unified usage API by packaging various commonly used drivers.
- [gookit/config](https://github.com/gookit/config) Go config management. support JSON, YAML, TOML, INI, HCL, ENV and Flags
- [gookit/color](https://github.com/gookit/color) A command-line color library with true color support, universal API methods and Windows support
- [gookit/filter](https://github.com/gookit/filter) Provide filtering, sanitizing, and conversion of golang data
- [gookit/validate](https://github.com/gookit/validate) Use for data validation and filtering. support Map, Struct, Form data
- [gookit/goutil](https://github.com/gookit/goutil) Some utils for the Go: string, array/slice, map, format, cli, env, filesystem, test and more
- More, please see https://github.com/gookit

## License

**MIT**
