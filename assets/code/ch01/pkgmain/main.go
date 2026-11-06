// Command pkgmain shows what "package main" means and the order in which a Go
// program starts: imported packages first, then package variables, then
// init functions, and finally main.
//
// Run from assets/code:
//
//	go run ./ch01/pkgmain
package main

import (
	"fmt"

	"gokb/ch01/pkgmain/shapes"
)

var banner = makeBanner()

func makeBanner() string {
	fmt.Println("  main: package variable initialized")
	return "=== area calculator ==="
}

func init() {
	fmt.Println("  main: init() runs")
}

func main() {
	fmt.Println("  main: main() starts. Your program begins here")
	fmt.Println(banner)
	fmt.Println("area of a 3x3 square:", shapes.Area(3))
}
