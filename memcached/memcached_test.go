package memcached_test

import (
	"fmt"

	"github.com/gookit/cache/memcached"
)

func Example() {
	c := memcached.Connect("10.0.0.1:11211", "10.0.0.2:11211")

	// set
	err := c.Set("name", "cache value", 60)
	if err != nil {
		panic(err)
	}

	// get
	val := c.Get("name")
	// del
	err = c.Del("name")
	if err != nil {
		panic(err)
	}

	// get: "cache value"
	fmt.Print(val)
}
