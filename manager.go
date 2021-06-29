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
		drivers: make(map[string]Cache, 8),
	}
}

// Register new cache driver
func (m *Manager) Register(name string, driver Cache) *Manager {
	// always use latest as default driver.
	m.defName = name
	// save driver instance
	m.drivers[name] = driver
	return m
}

// Unregister an cache driver
func (m *Manager) Unregister(name string) {
	delete(m.drivers, name)

	// reset default driver name.
	if m.defName == name {
		m.defName = ""
	}
}

// SetDefName set default driver name. alias of DefaultUse()
// Deprecated
//  please use DefaultUse() instead it
func (m *Manager) SetDefName(driverName string) {
	m.DefaultUse(driverName)
}

// DefaultUse set default driver name
func (m *Manager) DefaultUse(driverName string) {
	if _, ok := m.drivers[driverName]; !ok {
		panic("cache driver: " + driverName + " is not registered")
	}

	m.defName = driverName
}

// Default returns the default driver instance
func (m *Manager) Default() Cache {
	if c, ok := m.drivers[m.defName]; ok {
		return c
	}

	panic("cache driver: " + m.defName + " is not registered")
}

// Use driver object by name and set it as default driver.
func (m *Manager) Use(driverName string) Cache {
	m.DefaultUse(driverName)
	return m.Driver(driverName)
}

// Cache get driver by name. alias of Driver()
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

// Close all drivers
func (m *Manager) Close() (err error) {
	for _, cache := range m.drivers {
		err = cache.Close()
	}
	return err
}

// UnregisterAll cache drivers
func (m *Manager) UnregisterAll(fn ...func(cache Cache)) {
	m.defName = ""

	for name, driver := range m.drivers {
		if len(fn) > 0 {
			fn[0](driver)
		}

		delete(m.drivers, name)
	}
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
