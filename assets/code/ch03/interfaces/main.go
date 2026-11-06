// Interfaces: implicit satisfaction, type switches and the nil-interface trap.
//
// Run from assets/code:
//
//	go run ./ch03/interfaces
package main

import (
	"errors"
	"fmt"
	"io"
	"math"
	"os"
	"strings"
)

// Shape is satisfied by any type with these two methods. No "implements".
type Shape interface {
	Area() float64
	Perimeter() float64
}

type Rect struct{ W, H float64 }
type Circle struct{ R float64 }

func (r Rect) Area() float64      { return r.W * r.H }
func (r Rect) Perimeter() float64 { return 2 * (r.W + r.H) }

func (c Circle) Area() float64      { return math.Pi * c.R * c.R }
func (c Circle) Perimeter() float64 { return 2 * math.Pi * c.R }

// Compile-time proof that both types satisfy Shape.
var (
	_ Shape = Rect{}
	_ Shape = Circle{}
)

func describe(s Shape) string {
	// A type switch inspects the dynamic type stored in the interface.
	switch v := s.(type) {
	case Rect:
		return fmt.Sprintf("rect %gx%g", v.W, v.H)
	case Circle:
		return fmt.Sprintf("circle r=%g", v.R)
	default:
		return "unknown shape"
	}
}

// upper is an io.Writer that upper-cases everything passing through it.
type upper struct{ w io.Writer }

func (u upper) Write(p []byte) (int, error) {
	return u.w.Write([]byte(strings.ToUpper(string(p))))
}

type MyErr struct{}

func (*MyErr) Error() string { return "my error" }

// mayFail returns a typed nil pointer inside an error interface. Bug!
func mayFail() error {
	var p *MyErr // nil pointer
	return p     // interface now holds (type=*MyErr, value=nil): NOT nil
}

func main() {
	shapes := []Shape{Rect{3, 4}, Circle{1}}
	for _, s := range shapes {
		fmt.Printf("%-12s area=%6.2f perimeter=%6.2f\n", describe(s), s.Area(), s.Perimeter())
	}

	// Small interfaces compose: io.Copy works with any Reader and any Writer.
	src := strings.NewReader("interfaces are satisfied implicitly\n")
	if _, err := io.Copy(upper{os.Stdout}, src); err != nil {
		fmt.Println("copy failed:", err)
	}

	// The comma-ok type assertion never panics.
	var s Shape = Circle{2}
	if c, ok := s.(Circle); ok {
		fmt.Println("it is a circle with radius", c.R)
	}

	err := mayFail()
	fmt.Println("err == nil?", err == nil)
	var target *MyErr
	fmt.Println("errors.As finds *MyErr?", errors.As(err, &target))
}
