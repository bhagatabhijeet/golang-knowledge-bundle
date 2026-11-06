# Chapter 4 · Concurrency

Go's best-known feature. You'll start with one goroutine and finish with production-grade pipelines that shut down cleanly. Read the lessons in order: each one builds on the last.

* [Goroutines](/04-concurrency/goroutines.md) - Launch thousands of lightweight concurrent tasks with one keyword and learn how Go's scheduler juggles them across your CPU cores.
* [Channels](/04-concurrency/channels.md) - Typed pipes that let goroutines hand each other data safely, so they share memory by communicating instead of fighting over it.
* [Select](/04-concurrency/select.md) - The switch statement for channels that lets one goroutine wait on many events at once and react to whichever happens first.
* [Sync Primitives](/04-concurrency/sync-primitives.md) - Mutexes, read-write locks, Once, atomics and sync.Map are the precision tools for guarding shared state, and the race detector checks that you used them right.
* [Context](/04-concurrency/context.md) - The one parameter that carries deadlines, cancellation signals and request-scoped values through every layer of a Go program.
* [Concurrency Patterns](/04-concurrency/concurrency-patterns.md) - Pipelines, fan-out and fan-in, worker pools, semaphores and errgroup are the five patterns behind most concurrent Go code in production.

# Runnable Examples

* [goroutines](/assets/code/ch04/goroutines/main.go) - WaitGroup and concurrent "coffee orders"
* [channels](/assets/code/ch04/channels/main.go) - unbuffered, buffered, directional and closed channels
* [selectdemo](/assets/code/ch04/selectdemo/main.go) - racing sources, timeouts, default and tickers
* [syncprims](/assets/code/ch04/syncprims/main.go) - Mutex, RWMutex, Once, atomic and sync.Map
* [contextdemo](/assets/code/ch04/contextdemo/main.go) - timeouts, cancellation trees and causes
* [patterns](/assets/code/ch04/patterns/main.go) - pipeline, fan-out/fan-in, worker pool and semaphore
