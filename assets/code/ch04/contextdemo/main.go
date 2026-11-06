// Command contextdemo shows cancellation, timeouts, causes and
// request-scoped values flowing through a tree of contexts.
//
// Run from assets/code:  go run ./ch04/contextdemo
package main

import (
	"context"
	"errors"
	"fmt"
	"sync"
	"time"
)

// An unexported key type prevents collisions with other packages' keys.
type ctxKey string

const requestIDKey ctxKey = "request-id"

// work simulates a job that respects cancellation.
func work(ctx context.Context, name string, d time.Duration) error {
	id, _ := ctx.Value(requestIDKey).(string)
	select {
	case <-time.After(d):
		fmt.Printf("[%s] %-8s finished\n", id, name)
		return nil
	case <-ctx.Done():
		fmt.Printf("[%s] %-8s stopped: %v\n", id, name, ctx.Err())
		return ctx.Err()
	}
}

func main() {
	root := context.WithValue(context.Background(), requestIDKey, "req-42")

	// 1. Timeout: the deadline cancels slow work automatically.
	ctx, cancel := context.WithTimeout(root, 50*time.Millisecond)
	defer cancel() // always call cancel to release resources
	fmt.Println("fast err:", work(ctx, "fast", 10*time.Millisecond))
	err := work(ctx, "slow", 500*time.Millisecond)
	fmt.Println("slow err is DeadlineExceeded:", errors.Is(err, context.DeadlineExceeded))

	// 2. Manual cancel: one call stops every child.
	parent, stop := context.WithCancel(root)
	var wg sync.WaitGroup
	for _, name := range []string{"worker-1", "worker-2", "worker-3"} {
		wg.Add(1)
		go func() {
			defer wg.Done()
			_ = work(parent, name, time.Second)
		}()
	}
	time.Sleep(20 * time.Millisecond)
	fmt.Println("-- user pressed cancel --")
	stop()
	wg.Wait()

	// 3. Cause: record WHY it was cancelled (Go 1.20+).
	cctx, cancelCause := context.WithCancelCause(root)
	cancelCause(errors.New("shutting down for deploy"))
	fmt.Println("err:  ", cctx.Err())
	fmt.Println("cause:", context.Cause(cctx))
}
