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
	Set(key string, v interface{}, ttl time.Duration) error
	Del(key string) error
	// multi op
	GetMulti(keys []string) []interface{}
	SetMulti(mv map[string]interface{}, ttl time.Duration) error
	DelMulti(keys []string) error
	// clear
	Clear() error
}

type MarshalFunc func(v interface{}) ([]byte, error)
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
var manager = &CacheManager{}

// Init default manager instance
func Init(name string, driver CacheFace) *CacheManager {
	manager.SetDefName(name)
	manager.Add(name, driver)
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
func Set(key string, v interface{}, ttl time.Duration) error {
	return manager.Default().Set(key, v, ttl)
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
