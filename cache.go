// Package cache is a generic cache use and cache manager for golang.
// FileCache is a simple local file system cache implement.
// MemoryCache is a simple memory cache implement.
package cache

import (
	"encoding/json"
	"io"
	"time"
)

// Cache interface definition
type Cache interface {
	// close
	io.Closer
	// clear
	Clear() error
	// basic operation
	Has(key string) bool
	Get(key string) interface{}
	Set(key string, val interface{}, ttl time.Duration) (err error)
	Del(key string) error
	// multi operation
	GetMulti(keys []string) map[string]interface{}
	SetMulti(values map[string]interface{}, ttl time.Duration) (err error)
	DelMulti(keys []string) error
}

// some generic expire time define.
const (
	// Always exist
	Forever = 0
	// 1 second
	Seconds1 = time.Second
	// 2 second
	Seconds2 = 2*time.Second
	// 3 second
	Seconds3 = 3*time.Second
	// 5 second
	Seconds5 = 5*time.Second
	// 6 second
	Seconds6 = 6*time.Second
	// 7 second
	Seconds7 = 7*time.Second
	// 8 second
	Seconds8 = 8*time.Second
	// 9 second
	Seconds9 = 9*time.Second
	// 10 second
	Seconds10 = 10*time.Second
	// 15 second
	Seconds15 = 15*time.Second
	// 20 second
	Seconds20 = 20*time.Second
	// 30 second
	Seconds30 = 30*time.Second

	// 1 minutes
	OneMinutes = 60 * time.Second
	// 2 minutes
	TwoMinutes = 120 * time.Second
	// 3 minutes
	ThreeMinutes = 180 * time.Second
	// 5 minutes
	FiveMinutes = 300 * time.Second
	// 10 minutes
	TenMinutes = 600 * time.Second
	// 15 minutes
	FifteenMinutes = 900 * time.Second
	// half an hour
	HalfHour = 1800 * time.Second
	// 1 hour
	OneHour = 3600 * time.Second
	// 2 hours
	TwoHour = 7200 * time.Second
	// 3 hours
	ThreeHour = 10800 * time.Second
	// 12 hours(half of the day)
	HalfDay = 43200 * time.Second
	// 24 hours(1 day)
	OneDay = 86400 * time.Second
	// 2 day
	TwoDay = 172800 * time.Second
	// 3 day
	ThreeDay = 259200 * time.Second
	// 7 day(one week)
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
