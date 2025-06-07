package main

import "fmt"

func main() {
	defer fmt.Println("World")
	defer fmt.Println("One")
	defer fmt.Println("Two")
	fmt.Println("Hello")
	myDefer()

}

// World, One, Two
// 0, 1, 2, 3, 4
// hello, 4321, Two, One, World

func myDefer() {
	for i := range 5 {
		defer fmt.Print(i)
	}
}
