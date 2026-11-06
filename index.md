---
okf_version: "0.2"
---

# The Go Knowledge Bundle

A book-style field guide to Go. Read the chapters in order as a course, or jump to any concept. Every lesson is an OKF concept with frontmatter, runnable code in `/assets/code`, and diagrams in `/assets/images`.

# Start here

* [Cover page](/README.md) - The book cover: what this guide is, how to read it, and the full table of contents.
* [Update log](/log.md) - Chronological history of changes to this bundle, newest first.

# Chapter 1 · Getting Started

* [Chapter overview](/01-getting-started/) - Why Go exists, a professional environment in 15 minutes, and your first programs.
* [Why Go?](/01-getting-started/why-go.md) - The origin story of a language built at Google to make large-scale software boring to build, and how it quietly came to run much of the cloud.
* [Environment Setup](/01-getting-started/environment-setup/) - Install Go, install VS Code, and configure the official Go extension, from a blank machine to a debugger-ready editor in about 15 minutes.
  * [Installing Go](/01-getting-started/environment-setup/installing-go.md) - Go from zero to a working `go version` on Windows, macOS or Linux in about five minutes, then learn what GOROOT, GOPATH and PATH really mean.
  * [Installing VS Code](/01-getting-started/environment-setup/installing-vscode.md) - Install Visual Studio Code on any OS, put the `code` command on your PATH, and learn the dozen shortcuts that make it feel like an extension of your hands.
  * [Configuring VS Code for Go](/01-getting-started/environment-setup/configuring-vscode-go.md) - Turn VS Code into a Go powerhouse with the official extension, gopls, Delve and a battle-tested settings.json, so you get autocomplete, formatting on save, linting, debugging and one-click tests.
* [Hello, World!](/01-getting-started/hello-world.md) - Write, run and dissect your first Go program, and see why every line in it is there on purpose.
* [Go Packages](/01-getting-started/go-packages/) - What package main means, the two types of package (executable and reusable), how imports work, and why package, module and workspace are not the same thing.
  * [What package main Means](/01-getting-started/go-packages/package-main.md) - The two types of Go package, executable (package main, which builds a runnable file) and reusable (helper code you import), and how a Go program starts up.
  * [Import Statements](/01-getting-started/go-packages/import-statements.md) - An import statement opens a door into another package, like "fmt", and gives your code its exported functions, types and constants, from the standard library, your own module, or the wider Go world.
  * [Package vs Module vs Workspace](/01-getting-started/go-packages/package-module-workspace.md) - Is a package the same as a project? Is a module a workspace? No, and once you see the four nesting levels of Go code, every import path and go.mod makes sense.
  * [Modules and Packages](/01-getting-started/go-packages/modules-and-packages.md) - Learn how Go organizes code (files into packages, packages into modules) and how a single capital letter decides what the world can see.
* [Go CLI](/01-getting-started/go-cli/) - The six commands you'll type every day: run, build, install, fmt, test and get, practiced on one small project.
  * [go run](/01-getting-started/go-cli/go-run.md) - Compile and launch your program in one breath, the fastest feedback loop in Go and the command you will type more than any other.
  * [go build](/01-getting-started/go-cli/go-build.md) - Turn your source into one self-contained binary you can copy to any machine and run, with no runtime, no installer and no dependencies.
  * [go install](/01-getting-started/go-cli/go-install.md) - Compile a program and drop it into your personal bin folder, so the command works from any terminal, whether it's your own tool or someone else's.
  * [go fmt](/01-getting-started/go-cli/go-fmt.md) - One command, one style, zero arguments, so all Go code everywhere looks like it was written by the same careful person.
  * [go test](/01-getting-started/go-cli/go-test.md) - Go's built-in test runner finds, compiles and runs your tests, examples and coverage with no framework to install and nothing to configure.
  * [go get](/01-getting-started/go-cli/go-get.md) - Add, upgrade, downgrade and remove your project's dependencies with one command, recorded in go.mod and locked by checksums in go.sum.
* [A Tour of the Toolchain](/01-getting-started/toolchain-tour.md) - Meet the `go` command, the Swiss Army knife that builds, tests, formats, vets, documents and cross-compiles your code with no plugins required.

# Chapter 2 · Language Basics

* [Chapter overview](/02-language-basics/) - Values, loops, functions, pointers, collections and text: the vocabulary of everyday Go.
* [Variables and Types](/02-language-basics/variables-and-types.md) - Declare variables three ways, meet Go's built-in types, and discover why "zero values" mean Go never hands you garbage memory.
* [Control Flow](/02-language-basics/control-flow.md) - One loop, a smarter if, and a switch that doesn't fall through. Go's control flow fits on an index card and still covers everything.
* [Functions](/02-language-basics/functions.md) - Multiple return values, variadic parameters, first-class functions, closures that remember, and `defer`, the cleanup button you press in advance.
* [Pointers](/02-language-basics/pointers.md) - Pointers without the fear. Learn when Go copies, when it shares, and how escape analysis quietly decides between the stack and the heap.
* [Slices and Maps](/02-language-basics/slices-and-maps.md) - Go's two workhorse collections, explained from the inside out. See the three-word slice header that makes append fast and occasionally surprising.
* [Strings and Runes](/02-language-basics/strings-and-runes.md) - Why `len("世界")` is 6, what a rune really is, and how to slice, iterate and build text in Go without mangling a single emoji.

# Chapter 3 · Types and Abstraction

* [Chapter overview](/03-types-and-abstraction/) - Structs, interfaces, composition, generics, errors and iterators: designing with Go's type system.
* [Structs and Methods](/03-types-and-abstraction/structs-and-methods.md) - Go has no classes, just structs with methods attached, and choosing between a value and a pointer receiver is the most important design call you will make about a type.
* [Interfaces](/03-types-and-abstraction/interfaces.md) - Go interfaces are satisfied silently, with no "implements" keyword, which makes the smallest abstractions in Go some of the most powerful in any language.
* [Embedding and Composition](/03-types-and-abstraction/embedding-and-composition.md) - Go swaps inheritance trees for Lego bricks, letting you embed one type inside another and have its methods promoted for free.
* [Generics](/03-types-and-abstraction/generics.md) - Type parameters let you write a function or data structure once and use it with any type the constraint allows, checked fully at compile time.
* [Errors](/03-types-and-abstraction/errors.md) - In Go an error is just a value you return, wrap and inspect, and learning to build and read error chains is what separates scripts from production code.
* [Iterators](/03-types-and-abstraction/iterators.md) - Since Go 1.23 your own functions can feed a for-range loop, which turns lazy sequences, infinite streams and custom collections into one-line loops.

# Chapter 4 · Concurrency

* [Chapter overview](/04-concurrency/) - Goroutines, channels, select, locks, context and the patterns that tie them together.
* [Goroutines](/04-concurrency/goroutines.md) - Launch thousands of lightweight concurrent tasks with one keyword and learn how Go's scheduler juggles them across your CPU cores.
* [Channels](/04-concurrency/channels.md) - Typed pipes that let goroutines hand each other data safely, so they share memory by communicating instead of fighting over it.
* [Select](/04-concurrency/select.md) - The switch statement for channels that lets one goroutine wait on many events at once and react to whichever happens first.
* [Sync Primitives](/04-concurrency/sync-primitives.md) - Mutexes, read-write locks, Once, atomics and sync.Map are the precision tools for guarding shared state, and the race detector checks that you used them right.
* [Context](/04-concurrency/context.md) - The one parameter that carries deadlines, cancellation signals and request-scoped values through every layer of a Go program.
* [Concurrency Patterns](/04-concurrency/concurrency-patterns.md) - Pipelines, fan-out and fan-in, worker pools, semaphores and errgroup are the five patterns behind most concurrent Go code in production.

# Chapter 5 · Production Go

* [Chapter overview](/05-production-go/) - Testing, profiling, HTTP services, project layout and idiomatic style for code that ships.
* [Testing in Go](/05-production-go/testing.md) - Go ships a test runner, coverage tool, benchmark harness, and fuzzer in the box, so learn the table-driven pattern once and every test you write after it is one line long.
* [Benchmarks and Profiling](/05-production-go/benchmarks-and-profiling.md) - Stop guessing where your program is slow; Go's benchmark harness, benchstat, pprof, and the execution tracer show you exactly where the time and memory go.
* [HTTP Services](/05-production-go/http-services.md) - Build a production-shaped JSON API with nothing but the standard library, covering method-and-wildcard routing, middleware, timeouts, and graceful shutdown.
* [Project Layout](/05-production-go/project-layout.md) - How to organize a Go repository with go.mod, cmd/ and internal/, based on the official guidance, and why you should start flat and add structure only when the code asks for it.
* [Idioms and Style](/05-production-go/idioms-and-style.md) - The habits that make Go code feel like Go, including gofmt, naming, small interfaces, useful zero values, early returns, the Go proverbs, and structured logging with log/slog.

# Chapter 6 · Reference

* [Chapter overview](/06-reference/) - Cheatsheet, gotchas and glossary for fast lookups.
* [Go Cheatsheet](/06-reference/cheatsheet.md) - Every everyday Go construct on one page — syntax, idioms, and commands you will reach for a hundred times a week.
* [Gotchas](/06-reference/gotchas.md) - The traps that catch every Go programmer at least once, each with a reproduction, an explanation, and a fix.
* [Glossary](/06-reference/glossary.md) - Plain-English definitions of every Go term in this book, from blank identifier to zero value.

# Assets

* [Runnable examples](/assets/code/) - One Go module (`gokb`) with a `package main` per lesson; run any of them with `go run ./chNN/<name>` from `assets/code`.
* [Diagrams](/assets/images/) - Self-contained SVG illustrations used throughout the book.
