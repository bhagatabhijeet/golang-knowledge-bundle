---
type: Glossary
title: Glossary
description: Plain-English definitions of every Go term in this book, from blank identifier to zero value.
tags: [go, reference, glossary, terminology]
status: stable
chapter: 6
order: 3
level: beginner
reading_time: 7 min
generated: { by: claude_code/claude-opus-5-5, at: 2026-11-06T00:00:00Z }
sources:
  - id: go-spec
    resource: https://go.dev/ref/spec
    title: The Go Programming Language Specification
  - id: go-mem
    resource: https://go.dev/ref/mem
    title: The Go Memory Model
---

# Glossary

*Words are tools too. Here are Go's, sharpened.*

> **In 30 seconds:** Each term gets a one-line definition and a link to the
> lesson that teaches it. Formal definitions live in the Go spec.[^go-spec]

| Term | Meaning | Learn more |
|---|---|---|
| **Blank identifier `_`** | A write-only placeholder that discards a value: `_, err := f()`. | [Variables and Types](/02-language-basics/variables-and-types.md) |
| **Buffered channel** | A channel with capacity. Sends block only when the buffer is full. | [Channels](/04-concurrency/channels.md) |
| **Closure** | A function literal that captures variables from its surrounding scope. | [Functions](/02-language-basics/functions.md) |
| **Constraint** | An interface that limits which types a type parameter accepts, e.g. `cmp.Ordered`. | [Generics](/03-types-and-abstraction/generics.md) |
| **Context** | A `context.Context` carrying deadlines, cancellation, and request-scoped values across API boundaries. | [Context](/04-concurrency/context.md) |
| **Data race** | Two goroutines access the same memory at the same time, with at least one write and no synchronization.[^go-mem] | [Sync Primitives](/04-concurrency/sync-primitives.md) |
| **defer** | Schedules a call to run when the surrounding function returns, in LIFO order. | [Functions](/02-language-basics/functions.md) |
| **Embedding** | Putting a type inside a struct or interface without a field name, which promotes its methods. | [Embedding and Composition](/03-types-and-abstraction/embedding-and-composition.md) |
| **Escape analysis** | The compiler pass that decides whether a value lives on the stack or the heap. | [Pointers](/02-language-basics/pointers.md) |
| **Executable package** | A `package main` with a `func main()`. It builds a runnable program. | [What package main Means](/01-getting-started/go-packages/package-main.md) |
| **Exported** | An identifier starting with an uppercase letter. It is visible outside its package. | [Modules and Packages](/01-getting-started/go-packages/modules-and-packages.md) |
| **Goroutine** | A lightweight thread managed by the Go runtime. It starts with a few KB of stack. | [Goroutines](/04-concurrency/goroutines.md) |
| **gopls** | The official Go language server, which powers editor features like completion and go-to-definition. | [Configuring VS Code for Go](/01-getting-started/environment-setup/configuring-vscode-go.md) |
| **GOPATH** | The legacy workspace directory. It now mainly holds the module cache and `go install` binaries. | [Installing Go](/01-getting-started/environment-setup/installing-go.md) |
| **Import statement** | `import "fmt"` gives a file access to another package's exported names, used as `fmt.Println`. | [Import Statements](/01-getting-started/go-packages/import-statements.md) |
| **Interface** | A set of method signatures. Any type with those methods satisfies it implicitly. | [Interfaces](/03-types-and-abstraction/interfaces.md) |
| **iota** | A constant counter that resets in each `const` block and increments per line. | [Variables and Types](/02-language-basics/variables-and-types.md) |
| **Iterator** | A function of type `iter.Seq[V]` / `iter.Seq2[K,V]` that you can `range` over (Go 1.23+). | [Iterators](/03-types-and-abstraction/iterators.md) |
| **Method set** | The methods callable on a type. It decides which interfaces the type satisfies. | [Structs and Methods](/03-types-and-abstraction/structs-and-methods.md) |
| **Module** | A versioned collection of packages defined by `go.mod`. | [Modules and Packages](/01-getting-started/go-packages/modules-and-packages.md) |
| **Package** | A directory of `.go` files compiled together. It is Go's unit of encapsulation. | [Modules and Packages](/01-getting-started/go-packages/modules-and-packages.md) |
| **panic / recover** | A panic unwinds the stack for unrecoverable errors. `recover` inside a deferred function stops it. | [Errors](/03-types-and-abstraction/errors.md) |
| **Receiver** | The value a method is called on, either a value receiver `(t T)` or a pointer receiver `(t *T)`. | [Structs and Methods](/03-types-and-abstraction/structs-and-methods.md) |
| **Reusable package** | Any package not named `main`: helper code (a library) that other packages import. | [What package main Means](/01-getting-started/go-packages/package-main.md) |
| **Rune** | An `int32` holding one Unicode code point. | [Strings and Runes](/02-language-basics/strings-and-runes.md) |
| **Scheduler (GMP)** | The runtime's model of goroutines (G) run on OS threads (M) through processors (P). | [Goroutines](/04-concurrency/goroutines.md) |
| **select** | Waits on several channel operations and runs the first one that is ready. | [Select](/04-concurrency/select.md) |
| **Sentinel error** | A package-level error value, like `io.EOF`, that callers compare against with `errors.Is`. | [Errors](/03-types-and-abstraction/errors.md) |
| **Slice** | A (pointer, len, cap) view over an underlying array. | [Slices and Maps](/02-language-basics/slices-and-maps.md) |
| **Table-driven test** | A test that loops over a slice of cases, often with `t.Run` subtests. | [Testing](/05-production-go/testing.md) |
| **Type parameter** | A placeholder type in a generic function or type: `func Map[T, U any]`. | [Generics](/03-types-and-abstraction/generics.md) |
| **Type switch** | `switch v := x.(type)` branches on an interface's dynamic type. | [Interfaces](/03-types-and-abstraction/interfaces.md) |
| **Workspace** | An optional `go.work` file listing local modules to build together. Not the same as a module or a package. | [Package vs Module vs Workspace](/01-getting-started/go-packages/package-module-workspace.md) |
| **Wrapping** | Adding context to an error with `fmt.Errorf("...: %w", err)` while keeping the original. | [Errors](/03-types-and-abstraction/errors.md) |
| **Zero value** | The default value of every type (`0`, `""`, `false`, `nil`). Good Go makes it useful. | [Idioms and Style](/05-production-go/idioms-and-style.md) |

## Key takeaways

- If a term confuses you mid-chapter, look it up here, then follow the link back.
- The spec is short and readable. Look up the formal definition there when you need precision.

[^go-spec]: The Go Programming Language Specification
[^go-mem]: The Go Memory Model

---
⬅️ [Gotchas](/06-reference/gotchas.md) · 🏠 [Index](/index.md) · ➡️ [Back to the cover](/README.md)
