// Embedding: composition with automatic method promotion.
//
// Run from assets/code:
//
//	go run ./ch03/embedding
package main

import (
	"fmt"
	"sync"
)

// Logger is a small reusable part.
type Logger struct{ Prefix string }

func (l Logger) Log(msg string) { fmt.Println(l.Prefix + msg) }

// Counter embeds a sync.Mutex and a Logger. Their methods are promoted:
// c.Lock(), c.Unlock() and c.Log() all work directly on a Counter.
type Counter struct {
	sync.Mutex
	Logger
	hits map[string]int
}

func NewCounter() *Counter {
	return &Counter{Logger: Logger{Prefix: "[counter] "}, hits: map[string]int{}}
}

func (c *Counter) Hit(page string) {
	c.Lock()
	defer c.Unlock()
	c.hits[page]++
	c.Log(fmt.Sprintf("%s -> %d", page, c.hits[page]))
}

// Shadowing: an outer method with the same name wins over the promoted one.
type LoudLogger struct{ Logger }

func (l LoudLogger) Log(msg string) { l.Logger.Log(msg + "!!!") }

// Interfaces embed too: io.ReadWriter is just Reader + Writer.
type Speaker interface{ Speak() string }
type Walker interface{ Walk() string }
type Robot interface {
	Speaker
	Walker
}

type R2 struct{}

func (R2) Speak() string { return "beep" }
func (R2) Walk() string  { return "roll" }

func main() {
	c := NewCounter()
	c.Hit("/home")
	c.Hit("/home")
	c.Hit("/about")

	LoudLogger{Logger{Prefix: "[loud] "}}.Log("shadowed")

	var r Robot = R2{}
	fmt.Println(r.Speak(), "and", r.Walk())
}
