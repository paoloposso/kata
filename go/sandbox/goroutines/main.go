package main

import (
	"fmt"
	"time"
)

func main() {
	c := make(chan int, 1000)

	go func() {
		for i := 0; i < 10000; i++ {
			c <- i
		}
		close(c) // Close the channel after sending all values
	}()

	// Listener
	for value := range c {
		fmt.Println("Received:", value, "time:", time.Now())
	}
	fmt.Println("Channel is closed, listener exiting.")
}
