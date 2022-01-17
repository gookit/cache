module github.com/gookit/cache

go 1.14

require (
	github.com/bluele/gcache v0.0.0-20190518031135-bc40bd653833
	github.com/bradfitz/gomemcache v0.0.0-20190913173617-a41fca850d0b
	github.com/dgraph-io/badger v1.6.2
	github.com/dgraph-io/ristretto v0.0.3 // indirect
	github.com/dgryski/go-farm v0.0.0-20200201041132-a6ae2369ad13 // indirect
	github.com/go-redis/redis/v8 v8.11.4
	github.com/golang/snappy v0.0.1 // indirect
	github.com/gomodule/redigo v1.8.8
	github.com/gookit/goutil v0.4.3
	github.com/gookit/gsr v0.0.4
	github.com/patrickmn/go-cache v2.1.0+incompatible
	github.com/pkg/errors v0.9.1 // indirect
	github.com/stretchr/testify v1.7.0
	github.com/syndtr/goleveldb v1.0.0
	github.com/tidwall/buntdb v1.2.9
	go.etcd.io/bbolt v1.3.6
)

exclude github.com/gomodule/redigo v2.0.0+incompatible
