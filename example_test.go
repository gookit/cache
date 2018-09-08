package cache

import (
	"fmt"
	"github.com/gookit/cache/redis"
)

func Example_Redis() {
	// init a driver
	SetDefName(DvrRedis)
	Register(DvrRedis, redis.Connect("127.0.0.1:6379", "", 0))

	// usage
	//
	// set
	Set("name", "cache value", TwoMinutes)
	// get
	val := Get("name")
	// del
	Del("name")

	// get: "cache value"
	fmt.Print(val)
}
