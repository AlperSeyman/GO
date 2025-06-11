package main

import "fmt"

func returnFunc(x string) func() {
	a := "hello"
	return func() {
		fmt.Println(a, x)
	}
}

func main() {

	result := returnFunc("world")

	result()
}
