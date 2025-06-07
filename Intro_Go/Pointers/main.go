package main

import "fmt"

func main() {

	myNumber := 23

	var ptr = &myNumber

	fmt.Println("Value of actual pointer", ptr)
	fmt.Println("Value of actual pointer", *ptr)

	*ptr = *ptr * 2

	fmt.Println("New value is: ", myNumber)

}
