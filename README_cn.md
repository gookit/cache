# Cache

![GitHub go.mod Go version](https://img.shields.io/github/go-mod/go-version/gookit/cache?style=flat-square)
[![GoDoc](https://godoc.org/github.com/gookit/cache?status.svg)](https://pkg.go.dev/github.com/gookit/cache)
[![Go Report Card](https://goreportcard.com/badge/github.com/gookit/cache)](https://goreportcard.com/report/github.com/gookit/cache)
[![Actions Status](https://github.com/gookit/cache/workflows/Unit-Tests/badge.svg)](https://github.com/gookit/cache/actions)

> **[EN README](README.md)**

Golang 通用的缓存管理使用库。

通过包装各种常用的驱动，屏蔽掉底层各个驱动的不同使用方法，来提供统一的使用API。

> 所有缓存驱动程序都实现了 `cache.Cache` 接口。 因此，您可以添加任何自定义驱动程序。

**已经封装的驱动:**

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

**内置实现:**

- file 简单的文件缓存(_当前包的内置实现_)
- memory 简单的内存缓存(_当前包的内置实现_)

> 注意：内置实现比较简单，不推荐生产环境使用；生产环境推荐使用上面列出的三方驱动。

## GoDoc

- [doc on gowalker](https://gowalker.org/github.com/gookit/cache)
- [godoc for gopkg](https://pkg.go.dev/gopkg.in/gookit/cache.v1)
- [godoc for github](https://pkg.go.dev/github.com/gookit/cache)

## 安装

```bash
go get github.com/gookit/cache
```

## 接口方法

所有缓存驱动程序都实现了 `cache.Cache` 接口。 因此，您可以添加任何自定义驱动程序。

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
	// clear & close
	Clear() error
	Close() error
}
```

## 使用

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
	// 注册一个（或多个）缓存驱动
	cache.Register(cache.DvrFile, cache.NewFileCache(""))
	// cache.Register(cache.DvrMemory, cache.NewMemoryCache())
	cache.Register(gcache.Name, gcache.New(1000))
	cache.Register(gocache.Name, gocache.NewGoCache(cache.OneDay, cache.FiveMinutes))
	cache.Register(redis.Name, redis.Connect("127.0.0.1:6379", "", 0))
	cache.Register(goredis.Name, goredis.Connect("127.0.0.1:6379", "", 0))

	// 设置默认驱动名称
	cache.DefaultUse(goredis.Name)

	// 快速使用（默认驱动）
	//
	// set
	cache.Set("name", "cache value", cache.TwoMinutes)
	// get
	val := cache.Get("name")
	// del
	cache.Del("name")

	// Out: "cache value"
	fmt.Print(val)
	
	// 使用已注册的其他驱动
	client := cache.Driver(gcache.Name)
	client.Set("key", "val", cache.Seconds3)
	val = client.Get("key")
	// Out: "val"
	fmt.Print(val)
}
```

## 设置选项

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

- [gookit/rux](https://github.com/gookit/rux) Simple and fast request router for golang HTTP
- [gookit/gcli](https://github.com/gookit/gcli) build CLI application, tool library, running CLI commands
- [gookit/slog](https://github.com/gookit/slog) Lightweight, extensible, configurable logging library written in Go
- [gookit/event](https://github.com/gookit/event) Lightweight event manager and dispatcher implements by Go
- [gookit/cache](https://github.com/gookit/cache) Provide a unified usage API by packaging various commonly used drivers.
- [gookit/config](https://github.com/gookit/config) Go config management. support JSON, YAML, TOML, INI, HCL, ENV and Flags
- [gookit/color](https://github.com/gookit/color) A command-line color library with true color support, universal API methods and Windows support
- [gookit/filter](https://github.com/gookit/filter) Provide filtering, sanitizing, and conversion of golang data
- [gookit/validate](https://github.com/gookit/validate) Use for data validation and filtering. support Map, Struct, Form data
- [gookit/ini](https://github.com/gookit/ini) INI parse, simple go config management, use INI files
- [gookit/goutil](https://github.com/gookit/goutil) Some utils for the Go: string, array/slice, map, format, cli, env, filesystem, test and more
- More, please see https://github.com/gookit

## License

**[MIT](LICENSE)**
