# cache

generic cache use and cache manager for golang.

Drivers:

- file
- redis
- memory
- memCached

## Usage


### Redis



## Interface

```go
// CacheFace interface definition
type CacheFace interface {
	// basic op
	Has(key string) bool
	Get(key string) interface{}
	Set(key string, val interface{}, ttl time.Duration) error
	Del(key string) error
	// multi op
	GetMulti(keys []string) []interface{}
	SetMulti(mv map[string]interface{}, ttl time.Duration) error
	DelMulti(keys []string) error
	// clear
	Clear() error
}
```

## License

**MIT**
