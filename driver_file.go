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

// Has cache key.
// TODO decode value, and check expire time
func (c *FileCache) Has(key string) bool {
	if c.MemoryCache.Has(key) {
		return true
	}

	path := c.GetFilename(key)
	return fileExists(path)
}

// Get value by key
func (c *FileCache) Get(key string) interface{} {
	c.lock.RLock()
	defer c.lock.RUnlock()

	return c.get(key)
}

func (c *FileCache) get(key string) interface{} {
	// read cache from memory
	if val := c.MemoryCache.get(key); val != nil {
		return val
	}

	// read cache from file
	bs, err := ioutil.ReadFile(c.GetFilename(key))
	if err != nil {
		c.SetLastErr(err)
		return nil
	}

	item := &Item{}
	if err = c.MustUnmarshal(bs, item); err != nil {
		c.SetLastErr(err)
		return nil
	}

	// check expire time
	if item.Exp == 0 || item.Exp > time.Now().Unix() {
		c.caches[key] = item // save to memory.
		return item.Val
	}

	// has been expired. delete it.
	c.SetLastErr(c.del(key))
	return nil
}

// Set value by key
func (c *FileCache) Set(key string, val interface{}, ttl time.Duration) (err error) {
	c.lock.Lock()
	defer c.lock.Unlock()

	return c.set(key, val, ttl)
}

func (c *FileCache) set(key string, val interface{}, ttl time.Duration) (err error) {
	if err = c.MemoryCache.set(key, val, ttl); err != nil {
		c.SetLastErr(err)
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

	if _, err = f.Write(bs); err != nil {
		return err
	}

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
func (c *FileCache) GetMulti(keys []string) map[string]interface{} {
	c.lock.RLock()
	defer c.lock.RUnlock()

	values := make(map[string]interface{}, len(keys))
	for _, key := range keys {
		values[key] = c.get(key)
	}

	return values
}

// SetMulti values by multi key
func (c *FileCache) SetMulti(values map[string]interface{}, ttl time.Duration) (err error) {
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
	return strings.Join([]string{
		c.cacheDir,
		str[0:6],
		c.opt.Prefix+str + ".data",
	}, "/")
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
