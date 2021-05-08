package gocache

import (
	"fmt"
	"time"
)

func ExampleGoCache() {
	c := New()
	key := "name"

	// set
	c.Set(key, "cache value", 2*time.Second)
	fmt.Println(c.Has(key))

	// get
	val := c.Get(key)
	fmt.Println(val)

	time.Sleep(2 * time.Second)

	// get expired
	val2 := c.Get(key)
	fmt.Println(val2)

	// del
	c.Del(key)
	fmt.Println(c.Has(key))

	// Output:
	// true
	// cache value
	// <nil>
	// false
}
