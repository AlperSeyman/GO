package main

import (
	"fmt"
	"time"
)

// Goroutines are just functions that leave the main thread and run in the background and come back to join the main
// thread once the functions are finished/ready to return any value.

//  Goroutines do not stop the program flow and are non blocking

func main() {

	var err error

	fmt.Println("Beginning Program.")
	go sayHello()
	fmt.Println("After sayHello() function.")

	// err = go doWork() // This is not accepted
	go printLetters()
	go printNumber()

	go func() {
		err = doWork()
	}()

	if err != nil {
		fmt.Println("Error: ", err)
	} else {
		fmt.Println("Work Completed Succesfully")
	}
	time.Sleep((2 * time.Second))

}

func sayHello() {
	time.Sleep(2 * time.Second)
	fmt.Println("Hello from Goroutine.")
}

func printNumber() {
	for i := 0; i < 5; i++ {
		fmt.Println("Number: ", i, time.Now())
		time.Sleep(100 * time.Millisecond)
	}
}

func printLetters() {
	for _, letter := range "abcde" {
		fmt.Println("Letter: ", string(letter), time.Now())
		time.Sleep(200 * time.Millisecond)
	}
}

func doWork() error {
	// Simulate work
	time.Sleep(1 * time.Second)
	return fmt.Errorf("an error occured in dowork")
}
