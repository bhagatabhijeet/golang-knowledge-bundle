// Command pointers shows value vs pointer semantics.
//
//	go run ./ch02/pointers
//	go build -gcflags=-m ./ch02/pointers   # print escape-analysis decisions
package main

import "fmt"

// Player is a small struct used to compare copies and pointers.
type Player struct {
	Name  string
	Score int
}

// bumpCopy receives a copy; the caller never sees the change.
func bumpCopy(p Player) { p.Score += 10 }

// bumpPtr receives an address; it mutates the caller's value.
func bumpPtr(p *Player) { p.Score += 10 }

// newPlayer returns a pointer to a local variable. That is safe in Go:
// the compiler notices the value outlives the call and puts it on the heap.
func newPlayer(name string) *Player {
	p := Player{Name: name}
	return &p
}

func main() {
	x := 1
	px := &x // px holds the address of x
	*px = 99 // write through the pointer
	fmt.Println("x =", x)

	p := Player{Name: "Ada"}
	bumpCopy(p)
	fmt.Println("after bumpCopy:", p.Score)
	bumpPtr(&p)
	fmt.Println("after bumpPtr: ", p.Score)

	q := newPlayer("Grace")
	q.Score = 7 // fields auto-dereference: no (*q).Score needed
	fmt.Printf("%+v\n", *q)

	var nothing *Player
	fmt.Println("nil pointer?", nothing == nil)

	// new(T) allocates a zeroed T and returns *T.
	n := new(int)
	*n = 5
	fmt.Println(*n)
}
