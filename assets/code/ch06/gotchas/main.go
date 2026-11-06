// Command gotchas reproduces classic Go surprises so you can see them happen.
//
// Run from assets/code:
//
//	go run ./ch06/gotchas
package main

import (
	"errors"
	"fmt"
)

type MyErr struct{}

func (*MyErr) Error() string { return "my error" }

// mayFail returns a typed nil pointer inside an error interface. Bug!
func mayFail() error {
	var p *MyErr // nil pointer
	return p     // interface now holds (type=*MyErr, value=nil): NOT nil
}

func main() {
	// 1. The nil interface trap.
	err := mayFail()
	fmt.Println("1. err != nil:", err != nil)

	// 2. Slices share backing arrays.
	a := []int{1, 2, 3, 4}
	b := a[:2]
	b = append(b, 99) // fits in cap, so it overwrites a[2]
	fmt.Println("2. a after append to b:", a)

	// 3. Writing to a nil map panics. Recover to show the message.
	func() {
		defer func() { fmt.Println("3. recovered:", recover()) }()
		var m map[string]int
		m["boom"] = 1
	}()

	// 4. Deferred call arguments are evaluated immediately.
	func() {
		x := 1
		defer fmt.Println("4. deferred x =", x)
		x = 2
		_ = x
	}()

	// 5. Since Go 1.22 each loop iteration gets a fresh variable.
	var funcs []func() int
	for i := range 3 {
		funcs = append(funcs, func() int { return i })
	}
	fmt.Println("5. closures see:", funcs[0](), funcs[1](), funcs[2]())

	// 6. Comparing errors: use errors.Is, not ==, once wrapping is involved.
	base := errors.New("base")
	wrapped := fmt.Errorf("ctx: %w", base)
	fmt.Println("6. == :", wrapped == base, "| errors.Is:", errors.Is(wrapped, base))
}
