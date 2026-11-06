---
type: Reference
title: Go Cheatsheet
description: Every everyday Go construct on one page — syntax, idioms, and commands you will reach for a hundred times a week.
tags: [go, reference, cheatsheet, syntax]
status: stable
chapter: 6
order: 1
level: beginner
reading_time: 10 min
examples: [/assets/code/ch06/cheatsheet/main.go]
generated: { by: claude_code/claude-opus-5-5, at: 2026-11-06T00:00:00Z }
sources:
  - id: go-spec
    resource: https://go.dev/ref/spec
    title: The Go Programming Language Specification
  - id: effective-go
    resource: https://go.dev/doc/effective_go
    title: Effective Go
---

# Go Cheatsheet

*Pin it to the wall. You'll come back to this page more than any other.*

> **In 30 seconds:** Go is small enough to fit on one page. Below are the
> declarations, control flow, types, concurrency, and tooling commands that
> cover most of the Go you will ever write. Every snippet is valid Go.[^go-spec]

## Declarations

```go
var a int            // zero value: 0
var b, c = 1, "two"  // types inferred
d := 3.14            // short declaration (inside functions only)
const Pi = 3.14159   // untyped constant
const (
    KB = 1 << (10 * (iota + 1)) // 1024
    MB                          // 1048576
)
type Celsius float64           // new named type
type Alias = string            // alias, same type
```

| Type family | Examples | Zero value |
|---|---|---|
| Boolean | `bool` | `false` |
| Integers | `int` `int8..64` `uint` `uint8..64` `uintptr` | `0` |
| Floats / complex | `float32` `float64` `complex128` | `0` |
| Text | `string`, `byte` (=`uint8`), `rune` (=`int32`) | `""` |
| Reference-like | pointer, slice, map, chan, func, interface | `nil` |
| Aggregates | array `[4]int`, struct | each field zeroed |

## Control flow

```go
if v, err := strconv.Atoi(s); err != nil { /* handle */ } else { use(v) }

for i := 0; i < 10; i++ {}     // classic
for cond {}                    // while
for {}                         // forever
for i, v := range slice {}     // index + value
for k, v := range m {}         // map (random order!)
for i := range 10 {}           // integers (Go 1.22+)

switch x := f(); {             // tagless switch
case x > 100: fmt.Println("big")
case x > 10:  fmt.Println("medium")
default:      fmt.Println("small")
}                              // no fallthrough unless you write `fallthrough`
```

## Functions

```go
func divmod(a, b int) (q, r int) { return a / b, a % b } // multiple returns
func sum(nums ...int) int { /* variadic */ }
add := func(x, y int) int { return x + y }               // function literal
defer f.Close()                                          // runs at function exit, LIFO
```

## Collections

```go
s := []int{1, 2, 3}
s = append(s, 4)            // may reallocate — always reassign
t := s[1:3]                 // shares backing array!
u := slices.Clone(s)        // independent copy
m := map[string]int{"a": 1}
v, ok := m["b"]             // comma-ok: ok == false, v == 0
delete(m, "a")
slices.Sort(s); slices.Contains(s, 3); maps.Keys(m) // stdlib helpers
```

## Structs, methods, interfaces

```go
type Point struct{ X, Y int }
func (p Point) Len() float64 { return math.Hypot(float64(p.X), float64(p.Y)) }
func (p *Point) Scale(k int) { p.X *= k; p.Y *= k } // pointer receiver mutates

type Shape interface{ Area() float64 } // satisfied implicitly

switch v := anyValue.(type) {          // type switch
case string: fmt.Println("string", v)
case int:    fmt.Println("int", v)
}
```

## Generics

```go
func Max[T cmp.Ordered](a, b T) T { if a > b { return a }; return b }
type Stack[T any] struct{ items []T }
func (s *Stack[T]) Push(v T) { s.items = append(s.items, v) }
```

## Errors

```go
if err != nil { return fmt.Errorf("load config: %w", err) } // wrap
errors.Is(err, fs.ErrNotExist)                              // match sentinel
var pe *fs.PathError; errors.As(err, &pe)                   // match type
```

## Concurrency

```go
go work()                                // start a goroutine
ch := make(chan int, 10)                 // buffered channel
ch <- 1; v := <-ch; close(ch)
for v := range ch {}                     // until closed
select {
case v := <-ch:              use(v)
case <-time.After(time.Second): fmt.Println("timeout")
case <-ctx.Done():           return ctx.Err()
}
var wg sync.WaitGroup; wg.Add(1); go func() { defer wg.Done() }(); wg.Wait()
var mu sync.Mutex; mu.Lock(); defer mu.Unlock()
ctx, cancel := context.WithTimeout(ctx, 2*time.Second); defer cancel()
```

## Toolchain

| Command | What it does |
|---|---|
| `go mod init example.com/app` | start a module |
| `go run .` | compile + run |
| `go build -o app ./cmd/app` | produce a binary |
| `go test ./... -race -cover` | test everything, race detector, coverage |
| `go test -bench=. -benchmem` | benchmarks |
| `go fmt ./...` / `go vet ./...` | format / static checks |
| `go mod tidy` | sync `go.mod` with imports |
| `go doc fmt.Printf` | docs in the terminal |
| `GOOS=linux GOARCH=arm64 go build` | cross-compile |

## Try it

A program that exercises a little of everything above:

```sh
cd assets/code
go run ./ch06/cheatsheet
```

```text
divmod(17, 5) = 3 2
sorted: [1 2 3 4 5]
max: banana
squares: [1 4 9 16 25]
err wraps ErrNotFound: true
```

## Watch out

- `:=` declares **new** variables. Inside an `if` or `for`, it can shadow an outer `err`.
- `append` may or may not reuse the backing array. Always write `s = append(s, x)`.
- Map iteration order is random on purpose. Sort the keys if order matters.
- Lock a mutex as close as possible to the data it protects, and never copy a struct that contains one.

## Key takeaways

- Most Go fits on one page. Its strength is that there is little to learn and a lot you can build with it.
- Effective Go is the companion reading to this page.[^effective-go]
- When in doubt, `go doc <pkg>.<Symbol>` beats searching the web.

[^go-spec]: The Go Programming Language Specification
[^effective-go]: Effective Go

---
⬅️ [Idioms and Style](/05-production-go/idioms-and-style.md) · 🏠 [Index](/index.md) · ➡️ [Gotchas](/06-reference/gotchas.md)
