// Package calc is a tiny library used to demonstrate Go testing:
// table-driven tests, subtests, benchmarks, examples and fuzzing.
//
// Run from assets/code:
//
//	go test ./ch05/calc -v
//	go test ./ch05/calc -bench=. -benchmem
//	go test ./ch05/calc -cover
//	go test ./ch05/calc -fuzz=FuzzReverse -fuzztime=10s
package calc

import (
	"errors"
	"strings"
)

// ErrDivideByZero is returned by Divide when the divisor is zero.
var ErrDivideByZero = errors.New("calc: divide by zero")

// Add returns a + b.
func Add(a, b int) int { return a + b }

// Divide returns a / b, or ErrDivideByZero when b == 0.
func Divide(a, b int) (int, error) {
	if b == 0 {
		return 0, ErrDivideByZero
	}
	return a / b, nil
}

// Reverse returns s with its runes in reverse order.
// It is rune-aware, so "héllo" becomes "olléh", not mangled bytes.
func Reverse(s string) string {
	r := []rune(s)
	for i, j := 0, len(r)-1; i < j; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}
	return string(r)
}

// JoinSlow concatenates words with + in a loop (quadratic copying).
func JoinSlow(words []string) string {
	out := ""
	for i, w := range words {
		if i > 0 {
			out += ","
		}
		out += w
	}
	return out
}

// JoinFast concatenates words using strings.Builder (one growing buffer).
func JoinFast(words []string) string {
	var b strings.Builder
	for i, w := range words {
		if i > 0 {
			b.WriteByte(',')
		}
		b.WriteString(w)
	}
	return b.String()
}
