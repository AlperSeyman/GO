package main

import "fmt"

// Mutable --> changeable
// Immutable --> not changeable

func changeFirst(slice []int) {
	slice[0] = 1000
}

func main() {
	x := []int{1, 3, 4, 6}
	fmt.Println(x)
	changeFirst(x)
	fmt.Println(x)
}
