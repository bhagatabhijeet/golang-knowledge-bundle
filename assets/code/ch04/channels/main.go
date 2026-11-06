// Command channels tours unbuffered, buffered, directional and closed channels.
//
// Run from assets/code:  go run ./ch04/channels
package main

import "fmt"

// producer can only SEND on out (chan<-).
func producer(n int, out chan<- int) {
	for i := 1; i <= n; i++ {
		out <- i * i
	}
	close(out) // the sender closes, never the receiver
}

// summer can only RECEIVE from in (<-chan).
func summer(in <-chan int, done chan<- int) {
	total := 0
	for v := range in { // the loop ends when in is closed and drained
		total += v
	}
	done <- total
}

func main() {
	// 1. Unbuffered: a send waits for a receiver (a hand-off).
	ping := make(chan string)
	go func() { ping <- "ping" }()
	fmt.Println("unbuffered got:", <-ping)

	// 2. Buffered: sends succeed until the buffer is full.
	mailbox := make(chan string, 2)
	mailbox <- "letter 1"
	mailbox <- "letter 2" // no receiver yet, still fine
	fmt.Printf("buffered: len=%d cap=%d\n", len(mailbox), cap(mailbox))
	fmt.Println(<-mailbox, "|", <-mailbox)

	// 3. Directional channels + range + close.
	squares := make(chan int)
	done := make(chan int)
	go producer(5, squares)
	go summer(squares, done)
	fmt.Println("sum of squares 1..5 =", <-done)

	// 4. Receiving from a closed channel: buffered values first, then zero value with ok=false.
	c := make(chan int, 1)
	c <- 42
	close(c)
	v, ok := <-c
	fmt.Println("first receive:", v, ok)
	v, ok = <-c
	fmt.Println("after drain:  ", v, ok)
}
