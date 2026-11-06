// Package shapes is a library package: it has no main function and cannot run
// on its own. Other packages import it.
package shapes

import "fmt"

// Registered counts shapes; it is initialized before any init function runs.
var Registered = register("square")

func register(name string) int {
	fmt.Println("  shapes: package variable initialized (registered", name+")")
	return 1
}

func init() {
	fmt.Println("  shapes: init() runs")
}

// Area returns the area of a square with the given side. Capitalized, so exported.
func Area(side float64) float64 { return side * side }
