# cache

[![GoDoc](https://godoc.org/github.com/gookit/cache?status.svg)](https://godoc.org/github.com/gookit/cache)
[![Go Report Card](https://goreportcard.com/badge/github.com/gookit/cache)](https://goreportcard.com/report/github.com/gookit/cache)

> **[EN README](README.md)**

Golang 通用的缓存管理使用库。

通过包装各种常用的驱动，屏蔽掉底层各个驱动的不同使用方法，来提供统一的使用API。

> 所有缓存驱动程序都实现了 `cache.Cache` 接口。 因此，您可以添加任何自定义驱动程序。

**支持的驱动:**

- file 简单的文件缓存(当前包的内置实现)
- memory 简单的内存缓存(当前包的内置实现)
- `redis`  by `github.com/gomodule/redigo`
- `memCached` by `github.com/bradfitz/gomemcache`
- `buntdb` by `github.com/tidwall/buntdb`
- `boltdb`  by `github.com/etcd-io/bbolt`
- `badger db` by `github.com/dgraph-io/badger`
- `nutsdb` by `github.com/xujiajun/nutsdb`
- `goleveldb` by `github.com/syndtr/goleveldb`

## GoDoc

- [doc on gowalker](https://gowalker.org/github.com/gookit/cache)
- [godoc for gopkg](https://godoc.org/gopkg.in/gookit/cache.v1)
- [godoc for github](https://godoc.org/github.com/gookit/cache)

## 安装

```bash
go get github.com/gookit/cache
```

## 接口方法

所有缓存驱动程序都实现了cache.Cache接口。 因此，您可以添加任何自定义驱动程序。

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
	client := cache.Driver(cache.DvrFile)
	client.Set("key", "val", 0)
	val = client.Get("key")
	// Out: "val"
	fmt.Print(val)
}
```

## License

**[MIT](LICENSE)**
