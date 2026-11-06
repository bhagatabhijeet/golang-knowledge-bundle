// Generics: type parameters, constraints, ~ and comparable.
//
// Run from assets/code:
//
//	go run ./ch03/generics
package main

import (
	"cmp"
	"fmt"
)

// Map transforms a slice of T into a slice of U.
func Map[T, U any](in []T, f func(T) U) []U {
	out := make([]U, 0, len(in))
	for _, v := range in {
		out = append(out, f(v))
	}
	return out
}

// Max works for any ordered type: ints, floats, strings.
func Max[T cmp.Ordered](a, b T) T {
	if a > b {
		return a
	}
	return b
}

// Number uses ~ so named types such as Celsius are accepted too.
type Number interface {
	~int | ~int64 | ~float64
}

func Sum[T Number](xs []T) T {
	var total T
	for _, x := range xs {
		total += x
	}
	return total
}

type Celsius float64

// Set is a generic type. comparable allows == and use as a map key.
type Set[T comparable] struct{ m map[T]struct{} }

func NewSet[T comparable](items ...T) *Set[T] {
	s := &Set[T]{m: make(map[T]struct{})}
	for _, it := range items {
		s.Add(it)
	}
	return s
}

func (s *Set[T]) Add(v T) { s.m[v] = struct{}{} }

func (s *Set[T]) Has(v T) bool {
	_, ok := s.m[v]
	return ok
}

func (s *Set[T]) Len() int { return len(s.m) }

func main() {
	lengths := Map([]string{"go", "gopher", "generics"}, func(s string) int { return len(s) })
	fmt.Println("lengths:", lengths)

	fmt.Println("max int:", Max(3, 7), "| max string:", Max("apple", "banana"))

	temps := []Celsius{21.5, 19.0, 23.5}
	fmt.Println("sum of temps:", Sum(temps)) // works thanks to ~float64

	langs := NewSet("go", "rust", "go", "zig")
	fmt.Println("set size:", langs.Len(), "| has go?", langs.Has("go"), "| has c?", langs.Has("c"))

	// Explicit instantiation produces an ordinary, non-generic function value.
	toStr := Map[int, string]
	fmt.Println(toStr([]int{1, 2}, func(i int) string { return fmt.Sprint("#", i) }))
}
