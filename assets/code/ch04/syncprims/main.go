// Command syncprims demonstrates Mutex, RWMutex, Once, atomic counters
// and sync.Map. Run it with -race to confirm it is race-free.
//
// Run from assets/code:  go run ./ch04/syncprims
package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

// Counter protects its map with a Mutex. The zero value is ready to use.
type Counter struct {
	mu   sync.Mutex
	hits map[string]int
}

func (c *Counter) Inc(key string) {
	c.mu.Lock()
	defer c.mu.Unlock()
	if c.hits == nil {
		c.hits = make(map[string]int)
	}
	c.hits[key]++
}

// Config is read often and written rarely: a good RWMutex case.
type Config struct {
	mu   sync.RWMutex
	vals map[string]string
}

func (c *Config) Get(k string) string {
	c.mu.RLock() // many readers may hold the read lock at once
	defer c.mu.RUnlock()
	return c.vals[k]
}

func (c *Config) Set(k, v string) {
	c.mu.Lock() // a writer waits until all readers have left
	defer c.mu.Unlock()
	c.vals[k] = v
}

var (
	loadOnce sync.Once
	settings map[string]string
)

func loadSettings() map[string]string {
	loadOnce.Do(func() {
		fmt.Println("(loading settings: this prints exactly once)")
		settings = map[string]string{"theme": "dark"}
	})
	return settings
}

func main() {
	var wg sync.WaitGroup

	// Mutex: 100 goroutines bump one counter.
	var c Counter
	for range 100 {
		wg.Add(1)
		go func() { defer wg.Done(); c.Inc("home") }()
	}
	wg.Wait()
	fmt.Println("mutex counter: ", c.hits["home"])

	// atomic: a lock-free counter.
	var requests atomic.Int64
	for range 1000 {
		wg.Add(1)
		go func() { defer wg.Done(); requests.Add(1) }()
	}
	wg.Wait()
	fmt.Println("atomic counter:", requests.Load())

	// RWMutex: one writer, many readers.
	cfg := &Config{vals: map[string]string{"region": "us-east"}}
	wg.Add(1)
	go func() { defer wg.Done(); cfg.Set("region", "eu-west") }()
	for range 4 {
		wg.Add(1)
		go func() { defer wg.Done(); _ = cfg.Get("region") }()
	}
	wg.Wait()
	fmt.Println("config region: ", cfg.Get("region"))

	// Once: three goroutines race, the initializer runs once.
	for range 3 {
		wg.Add(1)
		go func() { defer wg.Done(); _ = loadSettings() }()
	}
	wg.Wait()
	fmt.Println("theme:         ", loadSettings()["theme"])

	// sync.Map: concurrent-safe map for write-once, read-many keys.
	var cache sync.Map
	for i := range 3 {
		wg.Add(1)
		go func() { defer wg.Done(); cache.Store(i, i*i) }()
	}
	wg.Wait()
	v, _ := cache.Load(2)
	actual, loaded := cache.LoadOrStore(2, 999)
	fmt.Println("sync.Map[2] =", v, "| LoadOrStore kept:", actual, "loaded:", loaded)
}
