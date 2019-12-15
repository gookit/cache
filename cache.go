// Package cache is a generic cache use and cache manager for golang.
// FileCache is a simple local file system cache implement.
// MemoryCache is a simple memory cache implement.
package cache

import (
	"encoding/json"
	"time"
)

// Cache interface definition
type Cache interface {
	// basic operation
	Has(key string) bool
	Get(key string) interface{}
	Set(key string, val interface{}, ttl time.Duration) (err error)
	Del(key string) error
	// multi operation
	GetMulti(keys []string) map[string]interface{}
	SetMulti(values map[string]interface{}, ttl time.Duration) (err error)
	DelMulti(keys []string) error
	// clear and close
	Clear() error
	Close() error
}

// some generic expire time define.
const (
	// 永远存在
	Forever = 0
	// 1 分钟
	OneMinutes = 60 * time.Second
	// 2 分钟
	TwoMinutes = 120 * time.Second
	// 3 分钟
	ThreeMinutes = 180 * time.Second
	// 5 分钟
	FiveMinutes = 300 * time.Second
	// 10 分钟
	TenMinutes = 600 * time.Second
	// 半小时
	HalfHour = 1800 * time.Second
	// 1 小时
	OneHour = 3600 * time.Second
	// 2 小时
	TwoHour = 7200 * time.Second
	// 3 小时
	ThreeHour = 10800 * time.Second
	// 12 小时(半天)
	HalfDay = 43200 * time.Second
	// 24 小时(1 天)
	OneDay = 86400 * time.Second
	// 2 天
	TwoDay = 172800 * time.Second
	// 3 天
	ThreeDay = 259200 * time.Second
	// 7 天(一周)
	OneWeek = 604800 * time.Second
)

// MarshalFunc define
type MarshalFunc func(v interface{}) ([]byte, error)

// UnmarshalFunc define
type UnmarshalFunc func(data []byte, v interface{}) error

// data (Un)marshal func
var (
	Marshal   MarshalFunc   = json.Marshal
	Unmarshal UnmarshalFunc = json.Unmarshal
)

/*************************************************************
 * config default cache manager
 *************************************************************/

// default cache driver manager instance
var defMgr = NewManager()

// Register driver to manager instance
func Register(name string, driver Cache) *Manager {
	defMgr.DefaultUse(name)
	defMgr.Register(name, driver)
	return defMgr
}

// SetDefName set default driver name.
// Deprecated
//  please use DefaultUse() instead it
func SetDefName(driverName string) {
	defMgr.DefaultUse(driverName)
}

// DefaultUse set default driver name
func DefaultUse(driverName string) {
	defMgr.DefaultUse(driverName)
}

// Use driver object by name and set it as default driver.
func Use(driverName string) Cache {
	return defMgr.Use(driverName)
}

// GetCache returns a driver instance by name. alias of Driver()
func GetCache(driverName string) Cache {
	return defMgr.Cache(driverName)
}

// Driver get a driver instance by name
func Driver(driverName string) Cache {
	return defMgr.Driver(driverName)
}

// DefManager get default cache manager instance
func DefManager() *Manager {
	return defMgr
}

// Default get default cache driver instance
func Default() Cache {
	return defMgr.Default()
}

/*************************************************************
 * quick use by default cache driver
 *************************************************************/

// Has cache key
func Has(key string) bool {
	return defMgr.Default().Has(key)
}

// Get value by key
func Get(key string) interface{} {
	return defMgr.Default().Get(key)
}

// Set value by key
func Set(key string, val interface{}, ttl time.Duration) error {
	return defMgr.Default().Set(key, val, ttl)
}

// Del value by key
func Del(key string) error {
	return defMgr.Default().Del(key)
}

// GetMulti values by keys
func GetMulti(keys []string) map[string]interface{} {
	return defMgr.Default().GetMulti(keys)
}

// SetMulti values
func SetMulti(mv map[string]interface{}, ttl time.Duration) error {
	return defMgr.Default().SetMulti(mv, ttl)
}

// DelMulti values by keys
func DelMulti(keys []string) error {
	return defMgr.Default().DelMulti(keys)
}

// Clear all caches
func Clear() error {
	return defMgr.Default().Clear()
}
