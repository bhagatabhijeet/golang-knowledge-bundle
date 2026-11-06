// Command goroutines shows how to launch goroutines and wait for them
// with sync.WaitGroup.
//
// Run from assets/code:  go run ./ch04/goroutines
package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

func brew(name string, d time.Duration, results chan<- string) {
	time.Sleep(d) // pretend to do real work
	results <- fmt.Sprintf("%-10s ready after %v", name, d)
}

func main() {
	fmt.Println("logical CPUs Go may use (GOMAXPROCS):", runtime.GOMAXPROCS(0))

	orders := map[string]time.Duration{
		"espresso":  30 * time.Millisecond,
		"latte":     60 * time.Millisecond,
		"cold brew": 90 * time.Millisecond,
	}

	var wg sync.WaitGroup
	results := make(chan string, len(orders)) // buffered: workers never block

	start := time.Now()
	for name, d := range orders {
		wg.Add(1)
		go func() { // Go 1.22+: name and d are fresh variables per iteration
			defer wg.Done()
			brew(name, d, results)
		}()
	}

	wg.Wait()      // block until every goroutine has called Done
	close(results) // safe: there are no more senders

	for r := range results {
		fmt.Println(r)
	}
	// All three ran at the same time, so the total is about the slowest order, not the sum.
	elapsed := time.Since(start).Round(10 * time.Millisecond)
	fmt.Printf("all done in ~%v (sequential would take 180ms)\n", elapsed)
}
