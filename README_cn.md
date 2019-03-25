# cache

[![GoDoc](https://godoc.org/github.com/gookit/cache?status.svg)](https://godoc.org/github.com/gookit/cache)
[![Go Report Card](https://goreportcard.com/badge/github.com/gookit/cache)](https://goreportcard.com/report/github.com/gookit/cache)

> **[EN README](README.md)**

Go下通用的缓存使用库，通过包装各种常用的驱动，来提供统一的使用API。

支持的驱动:

- file 文件缓存(internal driver)
- memory 内存缓存(internal driver)
- redis powered by `github.com/gomodule/redigo`
- memCached powered by `github.com/bradfitz/gomemcache`
- buntdb powered by `github.com/tidwall/buntdb`
- boltdb powered by `github.com/etcd-io/bbolt`
- badger db https://github.com/dgraph-io/badger
- nutsdb https://github.com/xujiajun/nutsdb
- goleveldb https://github.com/syndtr/goleveldb

## GoDoc

- [doc on gowalker](https://gowalker.org/github.com/gookit/cache)
- [godoc for gopkg](https://godoc.org/gopkg.in/gookit/cache.v1)
- [godoc for github](https://godoc.org/github.com/gookit/cache)

## 接口方法

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

## 使用

```go
package main

import (
	"fmt"
	"github.com/gookit/cache"
	"github.com/gookit/cache/redis"
)

func main() {
	// 注册一个（或多个）缓存驱动
	cache.Register(cache.DvrFile, cache.NewFileCache(""))
	cache.Register(cache.DvrMemory, cache.NewMemoryCache())
	cache.Register(cache.DvrRedis, redis.Connect("127.0.0.1:6379", "", 0))
	
	// 设置默认驱动名称
	cache.SetDefName(cache.DvrRedis)

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
	client := cache.Use(cache.DvrFile)
	client.Set("key", "val", 0)
	val = client.Get("key")
	// Out: "val"
	fmt.Print(val)
}
```

## License

**[MIT](LICENSE)**
