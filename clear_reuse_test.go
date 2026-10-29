package cache_test

import (
	"os"
	"sync"
	"testing"
	"time"

	"github.com/gookit/cache"
)

func TestBuiltInCachesReusableAfterClear(t *testing.T) {
	for name, makeCache := range map[string]func(*testing.T) (cache.Cache, string){
		"memory": func(_ *testing.T) (cache.Cache, string) { return cache.NewMemoryCache(), "" },
		"file": func(t *testing.T) (cache.Cache, string) {
			dir := t.TempDir()
			return cache.NewFileCache(dir), dir
		},
	} {
		t.Run(name, func(t *testing.T) {
			for _, multi := range []bool{false, true} {
				method := "Set"
				if multi {
					method = "SetMulti"
				}
				t.Run(method, func(t *testing.T) {
					defer func() {
						if p := recover(); p != nil {
							t.Fatalf("cache reuse panicked: %v", p)
						}
					}()
					c, dir := makeCache(t)
					if err := c.Set("old", "removed", time.Minute); err != nil {
						t.Fatal(err)
					}
					for i := 0; i < 2; i++ {
						if err := c.Clear(); err != nil {
							t.Fatal(err)
						}
					}
					if c.Has("old") || c.Get("old") != nil {
						t.Fatal("Clear retained old entry")
					}
					if file, ok := c.(*cache.FileCache); ok {
						if _, err := os.Stat(file.GetFilename("old")); !os.IsNotExist(err) {
							t.Fatalf("old cache file survives: %v", err)
						}
					}
					if multi {
						if err := c.SetMulti(map[string]any{"new": "retained", "second": "value"}, time.Minute); err != nil {
							t.Fatal(err)
						}
						if c.Get("second") != "value" {
							t.Fatal("SetMulti lost second entry")
						}
					} else if err := c.Set("new", "retained", time.Minute); err != nil {
						t.Fatal(err)
					}
					if !c.Has("new") || c.Get("new") != "retained" {
						t.Fatal("cache cannot be reused after Clear")
					}
					if file, ok := c.(*cache.FileCache); ok {
						if _, err := os.Stat(file.GetFilename("new")); err != nil {
							t.Fatal(err)
						}
						reloaded := cache.NewFileCache(dir)
						if reloaded.Get("new") != "retained" {
							t.Fatal("new cache file cannot be reloaded")
						}
						if multi && reloaded.Get("second") != "value" {
							t.Fatal("second cache file cannot be reloaded")
						}
					}
				})
			}
		})
	}
}

func TestMemoryClearConcurrentReads(t *testing.T) {
	c := cache.NewMemoryCache()
	if err := c.Set("key", "value", 0); err != nil {
		t.Fatal(err)
	}
	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		defer wg.Done()
		for i := 0; i < 1000; i++ {
			_ = c.Clear()
		}
	}()
	go func() {
		defer wg.Done()
		for i := 0; i < 1000; i++ {
			_ = c.Get("key")
			_ = c.Has("key")
		}
	}()
	wg.Wait()
}
