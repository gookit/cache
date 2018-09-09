package cache

// default supported cache driver name
const (
	DvrFile      = "file"
	DvrRedis     = "redis"
	DvrMemory    = "memory"
	DvrMemCached = "memCached"
)

/*************************************************************
 * Manager
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

// SetDefName set default driver name
func (m *Manager) SetDefName(driverName string) {
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

// Use returns a driver instance
func (m *Manager) Use(driverName string) Cache {
	return m.drivers[driverName]
}

// Get driver object by name
func (m *Manager) Get(name string) Cache {
	return m.Use(name)
}

// Driver object get
func (m *Manager) Driver(name string) Cache {
	return m.Use(name)
}

// DefName get default driver name
func (m *Manager) DefName() string {
	return m.defName
}
