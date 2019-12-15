package cache

import "time"

// default supported cache driver name
const (
	DvrFile      = "file"
	DvrRedis     = "redis"
	DvrMemory    = "memory"
	DvrMemCached = "memCached"
	DvrBoltDB    = "boltDB"
	DvrBuntDB    = "buntDB"
)

/*************************************************************
 * Cache Manager
 *************************************************************/

// Manager definition
type Manager struct {
	// Debug bool
	// default driver name
	defName string
	// drivers map
	drivers map[string]Cache
}

// NewManager create a cache manager instance
func NewManager() *Manager {
	return &Manager{
		// defName: driverName,
		drivers: make(map[string]Cache),
	}
}

// SetDefName set default driver name. alias of DefaultUse()
// Deprecated
//  please use DefaultUse() instead it
func (m *Manager) SetDefName(driverName string) {
	m.defName = driverName
}

// DefaultUse set default driver name
func (m *Manager) DefaultUse(driverName string) {
	m.defName = driverName
}

// Register new driver object
func (m *Manager) Register(name string, driver Cache) *Manager {
	m.drivers[name] = driver
	return m
}

// Default returns the default driver instance
func (m *Manager) Default() Cache {
	return m.drivers[m.defName]
}

// Use driver object by name and set it as default driver.
func (m *Manager) Use(driverName string) Cache {
	m.DefaultUse(driverName)
	return m.Driver(driverName)
}

// Cache driver object by name. alias of Driver()
func (m *Manager) Cache(driverName string) Cache {
	return m.drivers[driverName]
}

// Driver get a driver instance by name
func (m *Manager) Driver(driverName string) Cache {
	return m.drivers[driverName]
}

// DefName get default driver name
func (m *Manager) DefName() string {
	return m.defName
}

/*************************************************************
 * Quick use by default cache driver
 *************************************************************/

// Has cache key
func (m *Manager) Has(key string) bool {
	return m.Default().Has(key)
}

// Get value by key
func (m *Manager) Get(key string) interface{} {
	return m.Default().Get(key)
}

// Set value by key
func (m *Manager) Set(key string, val interface{}, ttl time.Duration) error {
	return m.Default().Set(key, val, ttl)
}

// Del value by key
func (m *Manager) Del(key string) error {
	return m.Default().Del(key)
}

// GetMulti values by keys
func (m *Manager) GetMulti(keys []string) map[string]interface{} {
	return m.Default().GetMulti(keys)
}

// SetMulti values
func (m *Manager) SetMulti(mv map[string]interface{}, ttl time.Duration) error {
	return m.Default().SetMulti(mv, ttl)
}

// DelMulti values by keys
func (m *Manager) DelMulti(keys []string) error {
	return m.Default().DelMulti(keys)
}
