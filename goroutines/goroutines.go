package main

import (
	"fmt"
	"time"
)

func say(s string) {
	for i := 0; i < 5; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(s)
	}
}

func sum(a []int, c chan int) {
	sum := 0
	for _, v := range a {
		sum += v
	}
	c <- sum // send sum to c
}

func main() {
	go say("world")
	say("hello")

	// Channels are a typed conduit through which you can
	// send and receive values with the channel operator, <-

	a := []int{7, 2, 8, -9, 4, 0}

	// Like maps and slices, channels must be created before use
	c := make(chan int)

	go sum(a[:len(a)/2], c)
	go sum(a[len(a)/2:], c)

	// The data flows in the direction of the arrow
	x, y := <-c, <-c // receive from c

	// By default, sends and receives block until the other
	// side is ready. This allows goroutines to synchronize
	// without explicit locks or condition variables.

	fmt.Println(x, y, x+y)

	// Channels can be buffered.

	ch := make(chan int, 2)
	ch <- 1
	ch <- 2

	// Sends to a buffered channel block only when the buffer
	// is full. Receives block when the buffer is empty
	fmt.Println(<-ch)
	fmt.Println(<-ch)
}
