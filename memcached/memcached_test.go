package memcached

import "fmt"

func Example() {
	c := New("10.0.0.1:11211", "10.0.0.2:11211")

	// set
	c.Set("name", "cache value", 60)
	// get
	val := c.Get("name")
	// del
	c.Del("name")

	// get: "cache value"
	fmt.Print(val)
}
