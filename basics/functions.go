package basic

import "fmt"

func basics() {

	// func <name>(parameters list) returnType {
	// code block
	// return value
	// }

	// fmt.Println(add(1, 2))

	// greet := func() {
	// 	fmt.Println("Hello Anonymous Function")
	// }
	// greet()

	// operation := add
	// result := operation(4, 3)
	// fmt.Println(result)

	// Passing a function as an argument
	result := applyOperation(3, 5, add) // applyOperation returns int
	fmt.Println("5 + 3 =", result)

	multiply := createMultiplier(2)     // createMultiplier returns func. multiply is a function now.
	fmt.Println("6 * 2 =", multiply(3)) // multiply returns int

}

func add(a, b int) int {
	return a + b
}

// Function that takes a function as an argument
func applyOperation(x int, y int, operation func(int, int) int) int {
	return operation(x, y)
}

// Function that returns a function
func createMultiplier(factor int) func(x int) int {
	return func(x int) int {
		return x * factor
	}
}
