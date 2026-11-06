---
type: Lesson
title: Control Flow
description: One loop, a smarter if, and a switch that doesn't fall through. Go's control flow fits on an index card and still covers everything.
tags: [go, control-flow, for, if, switch, loops]
status: stable
chapter: 2
order: 2
level: beginner
reading_time: 7 min
examples: [/assets/code/ch02/controlflow/main.go]
generated: { by: claude_code/claude-opus-5-5, at: 2026-11-06T00:00:00Z }
sources:
  - id: effective-go-control
    resource: https://go.dev/doc/effective_go#control-structures
    title: Effective Go - Control structures
  - id: range-int
    resource: https://go.dev/blog/go1.22
    title: Go 1.22 is released! - The Go Blog
---

# Control Flow

*Go has one loop keyword. It turns out that's all you ever needed.*

> **In 30 seconds:** `for` is the only loop. It has a classic form, a "while" form, an infinite form and a `range` form. `if` and `switch` can run a short setup statement first. `switch` cases **don't fall through** by default and can match lists or arbitrary conditions. There are no parentheses around conditions, and the braces are mandatory.

## `for`: four loops in one keyword

```go
// 1. Classic C-style
for i := 0; i < 3; i++ {
	fmt.Println(i)
}

// 2. "while"
for n < 100 {
	n *= 3
}

// 3. Infinite (break or return to exit)
for {
	if done() {
		break
	}
}

// 4. range: over slices, arrays, strings, maps, channels, integers and iterator functions
for i, v := range []string{"a", "b"} {
	fmt.Println(i, v)
}
for i := range 3 { // Go 1.22+: 0, 1, 2
	fmt.Println(i)
}
```

| `range` over… | yields |
|---|---|
| slice / array | index, element |
| string | byte offset, rune |
| map | key, value (random order!) |
| channel | each received value until it's closed |
| integer `n` | `0` … `n-1` |

Don't need the index? Use the blank identifier: `for _, v := range xs`.

> **Go 1.22 fixed a classic bug:** each loop iteration now gets its **own** copy of the loop variable, so closures and goroutines that capture `i` behave the way you'd expect.[^range-int]

## `if` with a setup statement

```go
if err := save(doc); err != nil {
	return err // err exists only inside this if/else
}

if v, ok := cache[key]; ok {
	use(v)
} else {
	load(key)
}
```

This pattern keeps temporary variables in the smallest possible scope. You'll see it everywhere in Go code.[^effective-go-control]

## `switch`: the clean decision-maker

```go
switch day {
case "sat", "sun": // several values in one case
	fmt.Println("weekend")
case "mon", "tue", "wed", "thu", "fri":
	fmt.Println("weekday")
default:
	fmt.Println("unknown")
}
```

- **No fallthrough** unless you write the `fallthrough` keyword (rarely needed).
- Cases can be any expression, not only constants.

A **tagless switch** replaces a long if/else chain:

```go
switch {
case score >= 90:
	grade = "A"
case score >= 80:
	grade = "B"
default:
	grade = "C"
}
```

## Labeled `break` and `continue`

```go
outer:
	for i := range 3 {
		for j := range 3 {
			if i*j == 2 {
				break outer // leaves both loops
			}
		}
	}
```

## What about `goto`?

It exists, it's restricted (you can't jump over variable declarations or into blocks), and you'll almost never need it.

## Try it

From [`assets/code`](/assets/code/):

```bash
go run ./ch02/controlflow
```

```text
0 1 2 
n = 243
tick0 tick1 tick2 
odd
sat -> weekend
tue -> weekday
xyz -> unknown
grade B
found 1 2
```

## Watch out

- A plain `break` inside a `switch` or `select` inside a loop only exits the `switch`/`select`, **not** the loop. Use a label.
- Map iteration order is **deliberately randomized**. Never depend on it.
- `range` over a slice copies each element into `v`. Modifying `v` doesn't change the slice; write `xs[i]` instead.

## Key takeaways

- `for` does every kind of looping, and `range` covers most of them.
- `if x := ...; cond` keeps scopes tight.
- `switch` is safer than in C: no accidental fallthrough.

[^effective-go-control]: Effective Go - Control structures
[^range-int]: Go 1.22 is released! - The Go Blog

---
⬅️ [Variables and Types](/02-language-basics/variables-and-types.md) · 🏠 [Index](/index.md) · ➡️ [Functions](/02-language-basics/functions.md)
