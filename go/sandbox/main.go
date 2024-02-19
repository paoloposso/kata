package main

import (
	"fmt"
	"math/rand"
)

func main() {
	c := make(chan int32)
	end := make(chan bool)

	go func() {
		for {
			r := int32(rand.Intn(1000))
			if r != 55 {
				c <- r
			} else {
				end <- true
			}
		}
	}()

	for {
		select {
		case value := <-c:
			fmt.Println(value)
		case value := <-end:
			if value {
				return
			}
		}
	}
}
