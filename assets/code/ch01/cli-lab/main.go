// Command cli-lab is the practice project for the Go CLI lessons.
// It greets someone, and it is small enough that you can focus on the go command, not the code.
//
// Run from assets/code:
//
//	go run ./ch01/cli-lab -name Gopher
//	go build -o bin/ ./ch01/cli-lab
//	go install ./ch01/cli-lab
//	go test ./ch01/cli-lab/...
package main

import (
	"flag"
	"fmt"

	"gokb/ch01/cli-lab/greet"
)

// version is "dev" by default. Override it at build time:
//
//	go build -ldflags="-X main.version=1.0.0" -o bin/ ./ch01/cli-lab
var version = "dev"

func main() {
	name := flag.String("name", "World", "who to greet")
	shout := flag.Bool("shout", false, "greet in capitals")
	showVersion := flag.Bool("version", false, "print the version and exit")
	flag.Parse()

	if *showVersion {
		fmt.Println("cli-lab", version)
		return
	}
	fmt.Println(greet.Hello(*name, *shout))
}
