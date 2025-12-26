package basics

import "fmt"

func main() {

	// panic(interface{})

	// Example of a valid input
	//process(10)

	// Example of a invalid input
	process(-3)
}

func process(input int) {

	defer fmt.Println("Deferred 1")
	defer fmt.Println("Deferred 2")

	if input < 0 {
		fmt.Println("Before panic")
		panic("input must be non-negative number")
	}
	fmt.Println("Processing input: ", input)
}
