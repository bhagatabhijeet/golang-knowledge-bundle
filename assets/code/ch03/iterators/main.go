// Iterators: range-over-func with iter.Seq / iter.Seq2 (Go 1.23+).
//
// Run from assets/code:
//
//	go run ./ch03/iterators
package main

import (
	"fmt"
	"iter"
	"maps"
	"slices"
	"strings"
)

// Countdown is an iterator: a function that pushes values into yield.
func Countdown(from int) iter.Seq[int] {
	return func(yield func(int) bool) {
		for i := from; i > 0; i-- {
			if !yield(i) { // the loop body said "break"
				return
			}
		}
	}
}

// Fib yields Fibonacci numbers forever; the caller decides when to stop.
func Fib() iter.Seq[int] {
	return func(yield func(int) bool) {
		a, b := 0, 1
		for {
			if !yield(a) {
				return
			}
			a, b = b, a+b
		}
	}
}

// Filter is an iterator adapter: it lazily keeps values that match.
func Filter[T any](seq iter.Seq[T], keep func(T) bool) iter.Seq[T] {
	return func(yield func(T) bool) {
		for v := range seq {
			if keep(v) && !yield(v) {
				return
			}
		}
	}
}

// Words yields (index, word) pairs: an iter.Seq2.
func Words(s string) iter.Seq2[int, string] {
	return func(yield func(int, string) bool) {
		for i, w := range strings.Fields(s) {
			if !yield(i, w) {
				return
			}
		}
	}
}

func main() {
	for n := range Countdown(3) {
		fmt.Print(n, " ")
	}
	fmt.Println("liftoff!")

	for f := range Filter(Fib(), func(n int) bool { return n%2 == 0 }) {
		if f > 100 {
			break // safe: yield returns false and Fib stops
		}
		fmt.Print(f, " ")
	}
	fmt.Println()

	for i, w := range Words("range over func") {
		fmt.Printf("%d:%s ", i, w)
	}
	fmt.Println()

	// Standard library helpers speak iterators too.
	ages := map[string]int{"ada": 36, "grace": 85, "linus": 54}
	fmt.Println("sorted names:", slices.Sorted(maps.Keys(ages)))

	evens := slices.Collect(Filter(slices.Values([]int{1, 2, 3, 4, 5, 6}), func(n int) bool { return n%2 == 0 }))
	fmt.Println("evens:", evens)

	for i, v := range slices.Backward([]string{"a", "b", "c"}) {
		fmt.Print(i, v, " ")
	}
	fmt.Println()

	// Range over an integer (Go 1.22+).
	for i := range 3 {
		fmt.Print(i, " ")
	}
	fmt.Println()
}
