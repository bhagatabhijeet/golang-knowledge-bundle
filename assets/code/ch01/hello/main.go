// Command hello is the first program of the book.
//
// Run it from the assets/code directory:
//
//	go run ./ch01/hello
//	go run ./ch01/hello Gopher
package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	// os.Args[0] is the program itself; everything after it is user input.
	name := "World"
	if len(os.Args) > 1 {
		name = strings.Join(os.Args[1:], " ")
	}
	fmt.Printf("Hello, %s! Welcome to Go.\n", name)
}
