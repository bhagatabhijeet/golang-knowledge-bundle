---
type: Reference
title: Gotchas
description: The traps that catch every Go programmer at least once, each with a reproduction, an explanation, and a fix.
tags: [go, reference, gotchas, pitfalls, debugging]
status: stable
chapter: 6
order: 2
level: intermediate
reading_time: 9 min
examples: [/assets/code/ch06/gotchas/main.go]
generated: { by: claude_code/claude-opus-5-5, at: 2026-11-06T00:00:00Z }
sources:
  - id: faq-nil-error
    resource: https://go.dev/doc/faq#nil_error
    title: "Go FAQ: Why is my nil error value not equal to nil?"
  - id: loopvar
    resource: https://go.dev/blog/loopvar-preview
    title: Fixing For Loops in Go 1.22
  - id: slices-intro
    resource: https://go.dev/blog/slices-intro
    title: "Go Slices: usage and internals"
---

# Gotchas

*Every one of these has cost someone a night of debugging. Spend five minutes here instead.*

> **In 30 seconds:** Go has very few surprises, so the ones it has are well
> known. They are typed-nil interfaces, shared slice backing arrays, nil maps,
> eager `defer` arguments, the old loop-variable bug, and comparing wrapped
> errors with `==`. Run the demo and watch each one happen.

## Try it first

```sh
cd assets/code
go run ./ch06/gotchas
```

```text
1. err != nil: true
2. a after append to b: [1 2 99 4]
3. recovered: assignment to entry in nil map
4. deferred x = 1
5. closures see: 0 1 2
6. == : false | errors.Is: true
```

Now let's unpack each line.

## 1. The nil interface that isn't nil

```go
func mayFail() error {
    var p *MyErr // nil pointer
    return p     // ❌ error interface holds (type=*MyErr, value=nil)
}
```

An interface value is a pair: **(dynamic type, dynamic value)**. It equals
`nil` only when *both* halves are empty. Here the type half is set, so
`err != nil` is `true`.[^faq-nil-error]

**Fix:** return a literal `nil` on the success path, and never return a
concrete pointer type through an `error` result.

## 2. Slices that secretly share memory

```go
a := []int{1, 2, 3, 4}
b := a[:2]          // len 2, cap 4: same backing array
b = append(b, 99)   // ❌ fits in cap, so it writes over a[2]
```

A slice is a small header (pointer, length, capacity) that looks at an array.
Slicing copies the header, not the data.[^slices-intro]

**Fix:** use `slices.Clone(a[:2])`, or a full slice expression `a[:2:2]`, which
caps capacity so the next `append` must reallocate.

## 3. Writing to a nil map panics

```go
var m map[string]int
m["boom"] = 1 // ❌ panic: assignment to entry in nil map
```

Reading from a nil map is fine and returns zero values. Writing to it panics.

**Fix:** `m := make(map[string]int)` or `m := map[string]int{}`.

## 4. `defer` evaluates arguments now, runs later

```go
x := 1
defer fmt.Println(x) // captures 1 right now
x = 2                // prints 1
```

**Fix:** if you want the value at exit time, defer a closure:
`defer func() { fmt.Println(x) }()`.

## 5. The loop-variable bug (fixed in Go 1.22)

Before Go 1.22, every iteration shared one variable, so closures and
goroutines often all saw the last value. Since Go 1.22 each iteration has its
own copy, *as long as your `go.mod` says `go 1.22` or later*.[^loopvar]

**Fix:** keep the `go` line in `go.mod` current. In old code you will see the
idiom `i := i`. It is now redundant.

## 6. Comparing wrapped errors with `==`

```go
wrapped := fmt.Errorf("ctx: %w", base)
wrapped == base          // ❌ false
errors.Is(wrapped, base) // ✅ true
```

**Fix:** always use `errors.Is` for sentinel errors and `errors.As` for error
types. They walk the whole wrap chain.

## More traps worth knowing

| Trap | Symptom | Fix |
|---|---|---|
| Copying a `sync.Mutex` | silent races, `go vet` "copylocks" warning | pass structs holding locks by pointer |
| Goroutine leak | memory grows, goroutine count climbs | every goroutine needs a way to exit: `ctx.Done()`, a closed channel |
| Unclosed `resp.Body` | connections pile up | `defer resp.Body.Close()` right after the error check |
| `time.After` in a hot loop | timers pile up | reuse a `time.Timer` or use `context.WithTimeout` |
| Range over a string | indexes jump by >1 | it iterates **runes** with byte offsets. That is by design |
| Shadowed `err` with `:=` | an error is silently dropped | declare `err` once; let staticcheck or the `shadow` analyzer flag it |
| Integer division | `7/2 == 3` | convert to float first: `float64(7)/2` |
| JSON with unexported fields | fields missing from output | capitalize field names and add `json:"name"` tags |

## Key takeaways

- An interface is `nil` only when both its type and its value are nil.
- Slicing shares memory. Use clone or a full slice expression when you need independence.
- `make` your maps, wrap your errors with `%w`, and compare them with `errors.Is`/`errors.As`.
- Run `go vet` and the race detector. They catch half of this table automatically.

[^faq-nil-error]: Go FAQ: Why is my nil error value not equal to nil?
[^loopvar]: Fixing For Loops in Go 1.22
[^slices-intro]: Go Slices: usage and internals

---
⬅️ [Go Cheatsheet](/06-reference/cheatsheet.md) · 🏠 [Index](/index.md) · ➡️ [Glossary](/06-reference/glossary.md)
