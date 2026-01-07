package main

import (
	"fmt"
	"time"
)

// Goroutines are just functions that leave the main thread and run in the background and come back to join the main
// thread once the functions are finished/ready to return any value.

//  Goroutines do not stop the program flow and are non blocking

func main() {

	fmt.Println("Beginning Program.")
	go sayHello()
	fmt.Println("After sayHello() function.")

	time.Sleep((2 * time.Second))
	fmt.Println("Ending Program.")
}

func sayHello() {
	time.Sleep(2 * time.Second)
	fmt.Println("Hello from Goroutine.")
}
