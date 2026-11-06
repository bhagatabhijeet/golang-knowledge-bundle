# Chapter 5 · Production Go

* [Testing in Go](/05-production-go/testing.md) - Go ships a test runner, coverage tool, benchmark harness, and fuzzer in the box, so learn the table-driven pattern once and every test you write after it is one line long.
* [Benchmarks and Profiling](/05-production-go/benchmarks-and-profiling.md) - Stop guessing where your program is slow; Go's benchmark harness, benchstat, pprof, and the execution tracer show you exactly where the time and memory go.
* [HTTP Services](/05-production-go/http-services.md) - Build a production-shaped JSON API with nothing but the standard library, covering method-and-wildcard routing, middleware, timeouts, and graceful shutdown.
* [Project Layout](/05-production-go/project-layout.md) - How to organize a Go repository with go.mod, cmd/ and internal/, based on the official guidance, and why you should start flat and add structure only when the code asks for it.
* [Idioms and Style](/05-production-go/idioms-and-style.md) - The habits that make Go code feel like Go, including gofmt, naming, small interfaces, useful zero values, early returns, the Go proverbs, and structured logging with log/slog.

# Code for this chapter

* [calc](/assets/code/ch05/calc/) - Table-driven tests, subtests, an Example, a fuzz target and two benchmarks (`go test ./ch05/calc`)
* [httpserver](/assets/code/ch05/httpserver/) - Stdlib JSON API with routing, middleware, timeouts and graceful shutdown (`go run ./ch05/httpserver`)
* [profiling](/assets/code/ch05/profiling/) - Writes CPU and heap profiles for `go tool pprof` (`go run ./ch05/profiling`)
* [slogdemo](/assets/code/ch05/slogdemo/) - Structured logging with log/slog, in text and JSON (`go run ./ch05/slogdemo`)

# Where next

* [Chapter 6 · Reference](/06-reference/) - Cheatsheet, gotchas and glossary to keep open while you code
