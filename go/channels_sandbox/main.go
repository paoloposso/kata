package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int32, 4)
	ch2 := make(chan int32, 4)
	receiver := make(chan int32, 4)
	unbufRouter := make(chan int32)
	quit := make(chan int32)

	defer close(ch)
	defer close(ch2)
	defer close(unbufRouter)

	go func() {
		ch <- 4
	}()

	go func() {
		ch <- 8
	}()

	go func() {
		time.Sleep(time.Second * 2)
		ch <- 10
	}()

	go func() {
		unbufRouter <- 10
		unbufRouter <- 11
		unbufRouter <- 14
		unbufRouter <- 18
		unbufRouter <- 10
		unbufRouter <- 107
		unbufRouter <- 22
	}()

	go func() {
		for i := int32(0); i < 20; i++ {
			ch2 <- i
		}
	}()

	go func() {
		for {
			select {
			case msg := <-ch:
				fmt.Println("received from ch", msg)
			case msg2 := <-ch2:
				fmt.Println("received from ch2", msg2)
			case msg3 := <-receiver:
				fmt.Println("received from not buf", msg3)
				unbufRouter <- msg3
			case msgBuffff := <-unbufRouter:
				fmt.Println("ROUTED: ", msgBuffff)
			}
		}
	}()

	go func() {
		time.Sleep(time.Second * 6)
		close(quit)
	}()

	fmt.Println(<-quit) //will quit when this channel closes
}
