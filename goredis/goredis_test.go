package goredis_test

import (
	"fmt"

	"github.com/gookit/cache/goredis"
)

func Example() {
	// init driver
	c := goredis.Connect("127.0.0.1:6379", "", 0)

	// set
	_ = c.Set("name", "cache value", 60)

	// get
	val := c.Get("name")

	// del
	_ = c.Del("name")

	// get: "cache value"
	fmt.Print(val)
}
