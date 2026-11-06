// Command variables tours declarations, zero values, constants and conversions.
//
//	go run ./ch02/variables
package main

import "fmt"

// Weekday is a named type; iota is an auto-incrementing constant generator.
type Weekday int

const (
	Sunday  Weekday = iota // 0
	Monday                 // 1
	Tuesday                // 2
)

// Pi is an untyped constant: it adapts to the context it is used in.
const Pi = 3.14159

func main() {
	// Three ways to declare.
	var a int        // zero value: 0
	var b = "gopher" // type inferred: string
	c := 42.5        // short declaration (inside functions only): float64
	fmt.Println(a, b, c)

	// Every type has a useful zero value.
	var (
		s   string
		ok  bool
		ptr *int
		xs  []int
	)
	fmt.Printf("%q %v %v %v\n", s, ok, ptr, xs == nil)

	// No implicit conversions: you must be explicit.
	n := 7
	f := float64(n) / 2
	fmt.Println(f)

	// Untyped constants are flexible: Pi becomes a float32 here.
	var radius float32 = 2
	fmt.Println(Pi * radius * radius)

	fmt.Println(Monday, Tuesday)
	fmt.Printf("%T %T %T\n", n, f, Monday)
}
