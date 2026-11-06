// Command functions demonstrates multiple returns, variadics, closures and defer.
//
//	go run ./ch02/functions
package main

import (
	"errors"
	"fmt"
)

// divide returns two values: the result and an error.
func divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, errors.New("division by zero")
	}
	return a / b, nil
}

// sum is variadic: it accepts any number of ints.
func sum(nums ...int) int {
	total := 0
	for _, n := range nums {
		total += n
	}
	return total
}

// counter returns a closure that remembers its own count.
func counter() func() int {
	count := 0
	return func() int {
		count++
		return count
	}
}

// apply takes a function as a value.
func apply(xs []int, f func(int) int) []int {
	out := make([]int, 0, len(xs))
	for _, x := range xs {
		out = append(out, f(x))
	}
	return out
}

// deferDemo shows that deferred calls run in LIFO order when the function returns.
func deferDemo() {
	defer fmt.Println("deferred 1 (runs last)")
	defer fmt.Println("deferred 2 (runs first)")
	fmt.Println("body runs")
}

func main() {
	if q, err := divide(10, 4); err == nil {
		fmt.Println("10/4 =", q)
	}
	if _, err := divide(1, 0); err != nil {
		fmt.Println("error:", err)
	}

	fmt.Println("sum:", sum(1, 2, 3), sum())
	nums := []int{4, 5, 6}
	fmt.Println("sum of slice:", sum(nums...))

	next := counter()
	fmt.Println(next(), next(), next())

	fmt.Println(apply(nums, func(x int) int { return x * x }))

	deferDemo()
}
