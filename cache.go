// Package cache is a generic cache use and cache manager for golang.
// FileCache is a simple local file system cache implement.
// MemoryCache is a simple memory cache implement.
package cache

import (
	"encoding/json"
	"time"
)

// CacheFace interface definition
type CacheFace interface {
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

// some generic expire time define.
const (
	// 永远存在
	FOREVER = 0
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
var manager = NewManager()

// Register driver to manager instance
func Register(name string, driver CacheFace) *CacheManager {
	manager.SetDefName(name)
	manager.Register(name, driver)
	return manager
}

// SetDefName set default driver name
func SetDefName(driverName string) {
	manager.SetDefName(driverName)
}

// Use returns a driver instance
func Use(driverName string) CacheFace {
	return manager.drivers[driverName]
}

// Manager get default cache manager instance
func Manager() *CacheManager {
	return manager
}

// Default get default cache driver instance
func Default() CacheFace {
	return manager.Default()
}

/*************************************************************
 * quick use
 *************************************************************/

// Has cache key
func Has(key string) bool {
	return manager.Default().Has(key)
}

// Get value by key
func Get(key string) interface{} {
	return manager.Default().Get(key)
}

// Set value by key
func Set(key string, val interface{}, ttl time.Duration) error {
	return manager.Default().Set(key, val, ttl)
}

// Del value by key
func Del(key string) error {
	return manager.Default().Del(key)
}

// GetMulti values by keys
func GetMulti(keys []string) []interface{} {
	return manager.Default().GetMulti(keys)
}

// SetMulti values
func SetMulti(mv map[string]interface{}, ttl time.Duration) error {
	return manager.Default().SetMulti(mv, ttl)
}

// DelMulti values by keys
func DelMulti(keys []string) error {
	return manager.Default().DelMulti(keys)
}

// Clear all caches
func Clear() error {
	return manager.Default().Clear()
}
