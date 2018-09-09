package cache

// default supported cache driver name
const (
	DvrFile      = "file"
	DvrRedis     = "redis"
	DvrMemory    = "memory"
	DvrMemCached = "memCached"
)

/*************************************************************
 * CacheManager
 *************************************************************/

// CacheManager definition
type CacheManager struct {
	// Debug bool
	// default driver name
	defName string
	// drivers map
	drivers map[string]CacheFace
}

// NewManager create a cache manager instance
func NewManager() *CacheManager {
	return &CacheManager{
		// defName: driverName,
		drivers: make(map[string]CacheFace),
	}
}

// SetDefName set default driver name
func (m *CacheManager) SetDefName(driverName string) {
	m.defName = driverName
}

// Register new driver object
func (m *CacheManager) Register(name string, driver CacheFace) *CacheManager {
	m.drivers[name] = driver
	return m
}

// Default returns the default driver instance
func (m *CacheManager) Default() CacheFace {
	return m.drivers[m.defName]
}

// Use returns a driver instance
func (m *CacheManager) Use(driverName string) CacheFace {
	return m.drivers[driverName]
}

// Get driver object by name
func (m *CacheManager) Get(name string) CacheFace {
	return m.Use(name)
}

// Driver object get
func (m *CacheManager) Driver(name string) CacheFace {
	return m.Use(name)
}

// DefName get default driver name
func (m *CacheManager) DefName() string {
	return m.defName
}
