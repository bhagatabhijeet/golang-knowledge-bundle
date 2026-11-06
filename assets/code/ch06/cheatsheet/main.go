// Command cheatsheet exercises a little of everything from the Go Cheatsheet.
//
// Run from assets/code:
//
//	go run ./ch06/cheatsheet
package main

import (
	"cmp"
	"errors"
	"fmt"
	"slices"
)

// ErrNotFound is a sentinel error callers can match with errors.Is.
var ErrNotFound = errors.New("not found")

// divmod shows multiple named return values.
func divmod(a, b int) (q, r int) { return a / b, a % b }

// Max works for any ordered type thanks to generics.
func Max[T cmp.Ordered](a, b T) T {
	if a > b {
		return a
	}
	return b
}

func find(key string) error {
	// Wrap with %w so callers can still see the sentinel underneath.
	return fmt.Errorf("lookup %q: %w", key, ErrNotFound)
}

func main() {
	q, r := divmod(17, 5)
	fmt.Println("divmod(17, 5) =", q, r)

	nums := []int{5, 3, 1, 4, 2}
	slices.Sort(nums)
	fmt.Println("sorted:", nums)

	fmt.Println("max:", Max("apple", "banana"))

	// Fan work out to goroutines and collect results over a channel.
	results := make(chan [2]int)
	for i := range 5 {
		go func() { results <- [2]int{i, (i + 1) * (i + 1)} }()
	}
	squares := make([]int, 5)
	for range 5 {
		res := <-results
		squares[res[0]] = res[1]
	}
	fmt.Println("squares:", squares)

	err := find("gopher")
	fmt.Println("err wraps ErrNotFound:", errors.Is(err, ErrNotFound))
}
