package main

import "fmt"

func main() {

	// var ptr *int

	var ptr *int
	var a = 10
	ptr = &a // referencing

	fmt.Println(a)
	fmt.Println(ptr) //  dereferencing a pointer

	modifyValue(ptr)
	fmt.Println(a)
}

func modifyValue(ptr *int) {
	*ptr++
}
