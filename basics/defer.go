package main

import "fmt"

func main() {

	process(10)
}

func process(i int) {
	defer fmt.Println("Deffered i value: ", i)
	defer fmt.Println("First deferred statement executed")
	defer fmt.Println("Second deferred statement executed")
	defer fmt.Println("Third deferred statement executed")
	i++
	fmt.Println("Normal execution statement")
	fmt.Println("Value of i: ", i)
}
