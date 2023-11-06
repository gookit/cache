package cache

import (
	"crypto/md5"
	"encoding/hex"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"time"
)

// FileCache definition.
type FileCache struct {
	BaseDriver
	// caches in memory
	MemoryCache
	// cache directory path
	cacheDir string
	// DisableMemCache disable cache in memory
	DisableMemCache bool
	// FilePrefix cache file prefix
	// FilePrefix string
	// security key for generate cache file name.
	securityKey string
}

// NewFileCache create a FileCache instance
func NewFileCache(dir string, pfxAndKey ...string) *FileCache {
	if dir == "" { // empty, use system tmp dir
		dir = os.TempDir()
	}

	c := &FileCache{
		cacheDir: dir,
		// init a memory cache.
		MemoryCache: MemoryCache{caches: make(map[string]*Item)},
	}

	if ln := len(pfxAndKey); ln > 0 {
		// c.prefix = pfxAndKey[0]
		c.opt.Prefix = pfxAndKey[0]

		if ln > 1 {
			c.securityKey = pfxAndKey[1]
		}
	}

	return c
}

// Has cache key. will check expire time
func (c *FileCache) Has(key string) bool {
	return c.get(key) != nil
}

// Get value by key
func (c *FileCache) Get(key string) any {
	return c.get(key)
}

func (c *FileCache) get(key string) any {
	c.lock.RLock()
	// read cache from memory
	if val := c.MemoryCache.get(key); val != nil {
		return val
	}
	c.lock.RUnlock()

	// read cache from file
	bs, err := ioutil.ReadFile(c.GetFilename(key))
	if err != nil {
		c.SetLastErr(err)
		return nil
	}

	item := &Item{}
	if err = c.UnmarshalTo(bs, item); err != nil {
		c.SetLastErr(err)
		return nil
	}

	// check expired
	if item.Expired() {
		c.SetLastErr(c.del(key))
		return nil
	}

	c.lock.Lock()
	c.caches[key] = item // save to memory.
	c.lock.Unlock()
	return item.Val
}

// Set value by key
func (c *FileCache) Set(key string, val any, ttl time.Duration) (err error) {
	c.lock.Lock()
	defer c.lock.Unlock()

	return c.set(key, val, ttl)
}

func (c *FileCache) set(key string, val any, ttl time.Duration) (err error) {
	err = c.MemoryCache.set(key, val, ttl)
	if err != nil {
		return
	}

	// cache item data to file
	bs, err := c.MustMarshal(c.caches[key])
	if err != nil {
		c.SetLastErr(err)
		return
	}

	file := c.GetFilename(key)
	dir := filepath.Dir(file)
	if err = os.MkdirAll(dir, 0755); err != nil {
		c.SetLastErr(err)
		return
	}

	f, err := os.Create(file)
	if err != nil {
		return err
	}
	defer f.Close()

	_, err = f.Write(bs)
	return
}

// Del value by key
func (c *FileCache) Del(key string) error {
	c.lock.Lock()
	defer c.lock.Unlock()

	return c.del(key)
}

func (c *FileCache) del(key string) error {
	if err := c.MemoryCache.del(key); err != nil {
		return err
	}

	file := c.GetFilename(key)
	if fileExists(file) {
		return os.Remove(file)
	}

	return nil
}

// GetMulti values by multi key
func (c *FileCache) GetMulti(keys []string) map[string]any {
	c.lock.RLock()
	defer c.lock.RUnlock()

	data := make(map[string]any, len(keys))
	for _, key := range keys {
		data[key] = c.get(key)
	}

	return data
}

// SetMulti values by multi key
func (c *FileCache) SetMulti(values map[string]any, ttl time.Duration) (err error) {
	c.lock.Lock()
	defer c.lock.Unlock()

	for key, val := range values {
		if err = c.set(key, val, ttl); err != nil {
			return
		}
	}
	return
}

// DelMulti values by multi key
func (c *FileCache) DelMulti(keys []string) error {
	c.lock.Lock()
	defer c.lock.Unlock()

	for _, key := range keys {
		_ = c.del(key)
	}
	return nil
}

// Close cache
func (c *FileCache) Close() error {
	return nil
}

// Clear caches and files
func (c *FileCache) Clear() error {
	c.lock.Lock()
	defer c.lock.Unlock()

	for key := range c.caches {
		if file := c.GetFilename(key); fileExists(file) {
			err := os.Remove(file)
			if err != nil {
				return err
			}
		}
	}

	c.caches = nil
	// clear cache files
	return os.RemoveAll(c.cacheDir)
}

// GetFilename cache file name build
func (c *FileCache) GetFilename(key string) string {
	h := md5.New()
	if c.securityKey != "" {
		h.Write([]byte(c.securityKey + key))
	} else {
		h.Write([]byte(key))
	}

	str := hex.EncodeToString(h.Sum(nil))

	// return fmt.Sprintf("%s/%s/%s.data", c.cacheDir, str[0:6], c.prefix+str)
	return strings.Join([]string{c.cacheDir, str[0:6], c.opt.Prefix + str + ".data"}, "/")
}

// fileExists reports whether the named file or directory exists.
func fileExists(name string) bool {
	if _, err := os.Stat(name); err != nil {
		if os.IsNotExist(err) {
			return false
		}
	}
	return true
}
