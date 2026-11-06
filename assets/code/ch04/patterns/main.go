// Command patterns builds a pipeline with fan-out / fan-in, a bounded
// worker pool, and a semaphore, using only the standard library.
//
// Run from assets/code:  go run ./ch04/patterns
package main

import (
	"context"
	"fmt"
	"slices"
	"sync"
	"time"
)

// generate is stage 1: it emits numbers, then closes its output.
func generate(ctx context.Context, nums ...int) <-chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		for _, n := range nums {
			select {
			case out <- n:
			case <-ctx.Done():
				return
			}
		}
	}()
	return out
}

// square is stage 2. Several copies read the same input (fan-out).
func square(ctx context.Context, in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		for n := range in {
			time.Sleep(5 * time.Millisecond) // simulated work
			select {
			case out <- n * n:
			case <-ctx.Done():
				return
			}
		}
	}()
	return out
}

// merge is stage 3: it combines many channels into one (fan-in).
func merge(ctx context.Context, chans ...<-chan int) <-chan int {
	out := make(chan int)
	var wg sync.WaitGroup
	for _, c := range chans {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for v := range c {
				select {
				case out <- v:
				case <-ctx.Done():
					return
				}
			}
		}()
	}
	go func() { wg.Wait(); close(out) }() // close only after every sender is done
	return out
}

// workerPool runs jobs on a fixed number of workers.
func workerPool(jobs []string, workers int) []string {
	jobCh := make(chan string)
	resCh := make(chan string)
	var wg sync.WaitGroup
	for w := 1; w <= workers; w++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for j := range jobCh {
				resCh <- fmt.Sprintf("%s(w%d)", j, w)
			}
		}()
	}
	go func() {
		defer close(jobCh)
		for _, j := range jobs {
			jobCh <- j
		}
	}()
	go func() { wg.Wait(); close(resCh) }()

	var out []string
	for r := range resCh {
		out = append(out, r)
	}
	return out
}

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// Pipeline: generate -> square x3 (fan-out) -> merge (fan-in).
	src := generate(ctx, 1, 2, 3, 4, 5, 6, 7, 8, 9)
	results := merge(ctx, square(ctx, src), square(ctx, src), square(ctx, src))
	var squares []int
	for v := range results {
		squares = append(squares, v)
	}
	slices.Sort(squares) // arrival order varies, so sort before printing
	fmt.Println("pipeline squares:", squares)

	// Worker pool.
	done := workerPool([]string{"resize", "crop", "upload", "notify"}, 2)
	fmt.Println("worker pool finished", len(done), "jobs")

	// Semaphore: a buffered channel caps concurrency at 2.
	sem := make(chan struct{}, 2)
	var (
		wg           sync.WaitGroup
		mu           sync.Mutex
		active, peak int
	)
	for range 6 {
		wg.Add(1)
		go func() {
			defer wg.Done()
			sem <- struct{}{}        // acquire a slot
			defer func() { <-sem }() // release it
			mu.Lock()
			active++
			peak = max(peak, active)
			mu.Unlock()
			time.Sleep(10 * time.Millisecond)
			mu.Lock()
			active--
			mu.Unlock()
		}()
	}
	wg.Wait()
	fmt.Println("semaphore peak concurrency:", peak, "(limit 2)")
}
