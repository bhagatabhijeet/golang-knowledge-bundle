//go:build go1.24

// The build constraint above lets this one file use Go 1.24 APIs
// even though the module's go.mod says go 1.23.

package calc

import "testing"

// BenchmarkJoinSlow uses the modern b.Loop (Go 1.24+): no b.N bookkeeping,
// setup outside the loop is not timed, and the body is never optimized away.
func BenchmarkJoinSlow(b *testing.B) {
	for b.Loop() {
		JoinSlow(words)
	}
}
