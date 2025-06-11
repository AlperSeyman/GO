package main

import "fmt"

func main() {

	var pointer *string
	str := "hello"
	pointer = &str
	fmt.Println(str)
	*pointer = "changed"
	fmt.Println(str)
}
