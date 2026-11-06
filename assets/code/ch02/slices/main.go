// Command slices explores slice headers, append growth, and maps.
//
//	go run ./ch02/slices
package main

import (
	"fmt"
	"maps"
	"slices"
)

func main() {
	// A slice is a view: pointer + len + cap over a backing array.
	arr := [5]int{10, 20, 30, 40, 50}
	s := arr[1:3] // [20 30], len 2, cap 4
	fmt.Println(s, len(s), cap(s))

	// Slices share memory with their backing array.
	s[0] = 99
	fmt.Println("arr after s[0]=99:", arr)

	// append writes into spare capacity...
	s = append(s, 77)
	fmt.Println("arr after append:  ", arr)

	// ...and reallocates when capacity runs out.
	var grow []int
	prevCap := -1
	for i := range 20 {
		grow = append(grow, i)
		if cap(grow) != prevCap {
			fmt.Printf("len=%-2d cap=%d\n", len(grow), cap(grow))
			prevCap = cap(grow)
		}
	}

	// make preallocates; slices.Clone gives an independent copy.
	buf := make([]int, 0, 8)
	buf = append(buf, 3, 1, 2)
	clone := slices.Clone(buf)
	slices.Sort(clone)
	fmt.Println(buf, clone, slices.Contains(clone, 2))

	// Maps: hash tables with built-in syntax.
	ages := map[string]int{"ada": 36, "linus": 54}
	ages["grace"] = 85
	if age, ok := ages["alan"]; !ok {
		fmt.Println("alan missing, zero value:", age)
	}
	delete(ages, "linus")

	// Map iteration order is random; sort the keys for stable output.
	for _, k := range slices.Sorted(maps.Keys(ages)) {
		fmt.Println(k, ages[k])
	}
}
