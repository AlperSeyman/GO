package main

import (
	"fmt"
)

// func printAnything[T any](item T) { // [T any] makes this function "Generic"
// 	fmt.Println(item)
// }

// func swap[T any](a, b T) (T, T) {
// 	return b, a
// }

// type Box[T any] struct {
// 	Content T
// }

type Stack[T any] struct {
	elements []T // Slice
}

func (s *Stack[T]) push(element T) {
	s.elements = append(s.elements, element)
}

func (s *Stack[T]) pop() (T, bool) {

	if len(s.elements) == 0 {
		var zeroValue T
		return zeroValue, false
	}
	lastIndex := len(s.elements) - 1
	lastElement := s.elements[lastIndex]

	s.elements = s.elements[:lastIndex]

	return lastElement, true
}

func (s *Stack[T]) isEmpty() bool {
	return len(s.elements) == 0
}

func (s Stack[T]) printAll() {

	if len(s.elements) == 0 {
		fmt.Println("The stack is empty")
		return
	}
	fmt.Print("Stack Elements: ")
	for _, v := range s.elements {
		fmt.Print(v)
	}
	fmt.Println()

}

func main() {

	intStack := Stack[int]{}
	intStack.push(1)
	intStack.push(2)
	intStack.push(3)
	intStack.printAll()
	fmt.Println(intStack.pop())
	intStack.printAll()

	stringStack := Stack[string]{}
	stringStack.push("Hello")
	stringStack.push("World")
	stringStack.push("John")
	stringStack.printAll()
	fmt.Println(stringStack.pop())
	fmt.Println(stringStack.pop())
	fmt.Println("Is stringStack empty: ", stringStack.isEmpty())

	// printAnything(100)     // T becomes an 'int'
	// printAnything("Hello") // T becomes a 'string'
	// printAnything(3.14)    // T becomes a 'float'

	// x1, y1 := 1, 2
	// fmt.Println(x1, y1)
	// x1, y1 = swap(x1, y1)
	// fmt.Println(x1, y1)

	// x2, y2 := "John", "Jane"
	// x2, y2 = swap(x2, y2)
	// fmt.Println(x2, y2)

	// b1 := Box[int]{
	// 	Content: 500,
	// }

	// b2 := Box[string]{
	// 	Content: "Goold Metal",
	// }
	// fmt.Println(b1.Content)
	// fmt.Println(b2.Content)

}
