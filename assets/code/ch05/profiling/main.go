// Command profiling writes a CPU profile and a heap profile of a deliberately
// wasteful workload, ready to explore with `go tool pprof`.
// Profiles go to your temp directory so they never clutter the repo;
// the program prints their exact paths.
//
// Run from assets/code:
//
//	go run ./ch05/profiling
//	go tool pprof -top <printed cpu.prof path>
//	go tool pprof -http=:8081 <printed cpu.prof path>   # flame graph in the browser
//	go tool pprof -sample_index=alloc_space -top <printed mem.prof path>
package main

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime"
	"runtime/pprof"
	"strings"
)

// buildSlow concatenates with + : every step copies the whole string.
func buildSlow(n int) string {
	s := ""
	for i := 0; i < n; i++ {
		s += "x"
	}
	return s
}

// buildFast uses strings.Builder: amortized appends into one buffer.
func buildFast(n int) string {
	var b strings.Builder
	for i := 0; i < n; i++ {
		b.WriteByte('x')
	}
	return b.String()
}

func main() {
	cpuPath := filepath.Join(os.TempDir(), "cpu.prof")
	cpu, err := os.Create(cpuPath)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	defer cpu.Close()
	if err := pprof.StartCPUProfile(cpu); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	total := 0
	for i := 0; i < 20; i++ {
		total += len(buildSlow(20_000)) // the hot spot pprof should find
		total += len(buildFast(20_000))
	}
	pprof.StopCPUProfile()

	memPath := filepath.Join(os.TempDir(), "mem.prof")
	mem, err := os.Create(memPath)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	defer mem.Close()
	runtime.GC() // get up-to-date statistics before the heap snapshot
	if err := pprof.WriteHeapProfile(mem); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	fmt.Println("built", total, "bytes")
	fmt.Println("cpu profile:", cpuPath)
	fmt.Println("heap profile:", memPath)
}
