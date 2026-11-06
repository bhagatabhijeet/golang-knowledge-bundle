// Command platform reports what the Go toolchain built it for.
// Use it to watch cross-compilation work:
//
//	go run ./ch01/platform
//	GOOS=linux GOARCH=arm64 go build -o platform-linux ./ch01/platform
package main

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Println("Go version:", runtime.Version())
	fmt.Println("OS/Arch:   ", runtime.GOOS+"/"+runtime.GOARCH)
	fmt.Println("CPUs:      ", runtime.NumCPU())
}
