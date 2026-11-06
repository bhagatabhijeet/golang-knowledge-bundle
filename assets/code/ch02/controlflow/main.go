// Command controlflow shows Go's small set of control structures.
//
//	go run ./ch02/controlflow
package main

import "fmt"

func main() {
	// for is Go's only loop. Classic three-part form:
	for i := 0; i < 3; i++ {
		fmt.Print(i, " ")
	}
	fmt.Println()

	// "while" form: just a condition.
	n := 1
	for n < 100 {
		n *= 3
	}
	fmt.Println("n =", n)

	// Range over an integer (Go 1.22+).
	for i := range 3 {
		fmt.Print("tick", i, " ")
	}
	fmt.Println()

	// if with a short statement: v is scoped to the if/else.
	if v := n % 2; v == 0 {
		fmt.Println("even")
	} else {
		fmt.Println("odd")
	}

	// switch: no fallthrough by default; cases can list several values.
	for _, day := range []string{"sat", "tue", "xyz"} {
		switch day {
		case "sat", "sun":
			fmt.Println(day, "-> weekend")
		case "mon", "tue", "wed", "thu", "fri":
			fmt.Println(day, "-> weekday")
		default:
			fmt.Println(day, "-> unknown")
		}
	}

	// A tagless switch is a tidy if/else chain.
	score := 82
	switch {
	case score >= 90:
		fmt.Println("grade A")
	case score >= 80:
		fmt.Println("grade B")
	default:
		fmt.Println("keep going")
	}

	// Labeled break escapes nested loops.
outer:
	for i := range 3 {
		for j := range 3 {
			if i*j == 2 {
				fmt.Println("found", i, j)
				break outer
			}
		}
	}
}
