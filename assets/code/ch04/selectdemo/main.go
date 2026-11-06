// Command selectdemo shows select with multiple channels, a timeout,
// a non-blocking default case, and a ticker loop with a deadline.
//
// Run from assets/code:  go run ./ch04/selectdemo
package main

import (
	"fmt"
	"time"
)

// fetch returns a channel that delivers name after d.
func fetch(name string, d time.Duration) <-chan string {
	ch := make(chan string, 1) // buffered so the goroutine can always finish (no leak)
	go func() {
		time.Sleep(d)
		ch <- name
	}()
	return ch
}

func main() {
	// 1. Race two sources: the first ready case wins.
	a := fetch("mirror-A", 50*time.Millisecond)
	b := fetch("mirror-B", 10*time.Millisecond)
	select {
	case r := <-a:
		fmt.Println("winner:", r)
	case r := <-b:
		fmt.Println("winner:", r)
	}

	// 2. Timeout: give up if nothing arrives in time.
	select {
	case r := <-fetch("slow-db", 200*time.Millisecond):
		fmt.Println("got:", r)
	case <-time.After(50 * time.Millisecond):
		fmt.Println("timeout: slow-db took too long")
	}

	// 3. Non-blocking send with default.
	queue := make(chan int, 1)
	for i := 1; i <= 2; i++ {
		select {
		case queue <- i:
			fmt.Println("queued", i)
		default:
			fmt.Println("queue full, dropped", i)
		}
	}

	// 4. Loop with a ticker and a deadline.
	tick := time.NewTicker(20 * time.Millisecond)
	defer tick.Stop()
	deadline := time.After(75 * time.Millisecond)
	for n := 1; ; n++ {
		select {
		case <-tick.C:
			fmt.Println("tick", n)
		case <-deadline:
			fmt.Println("deadline reached, stopping")
			return
		}
	}
}
