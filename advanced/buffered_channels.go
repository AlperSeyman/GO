package main

import (
	"fmt"
	"time"
)

func main() {

	// variable := make(chan Type, capacity)
	ch := make(chan int, 2)

	ch <- 1 // send data to channel
	ch <- 2 // send data to channel

	go func() {
		time.Sleep(2 * time.Second)
		fmt.Println("Received :", <-ch)
	}()
	fmt.Println("Blocking starts")
	ch <- 3 // Blocks, because the buffer is full
	fmt.Println("Blocking ends")
	fmt.Println("Received :", <-ch)
	fmt.Println("Received :", <-ch)

	fmt.Println("Buffered Channels")
}
