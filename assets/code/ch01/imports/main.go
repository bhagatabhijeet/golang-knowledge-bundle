// Command imports shows every form of the import statement: grouped imports,
// standard library and local packages, an alias, and a blank import.
//
// Run from assets/code:
//
//	go run ./ch01/imports
package main

import (
	// Standard library: short paths, no domain name.
	"fmt"
	"math"
	str "strconv" // alias: refer to strconv as str in this file
	"strings"

	// Your own module: module path + folder.
	"gokb/ch01/pkgmain/shapes"

	// Blank import: run the package's init(), use nothing from it.
	_ "gokb/ch01/imports/plugin"
)

func main() {
	fmt.Println("fmt:     ", "Println comes from package fmt")
	fmt.Println("strings: ", strings.ToUpper("shout"), strings.Repeat("go", 3))
	fmt.Println("math:    ", math.Sqrt(16), math.Pi)
	fmt.Println("strconv: ", str.Itoa(42)+" is now a string")
	fmt.Println("shapes:  ", shapes.Area(4))
}
