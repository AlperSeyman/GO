package main

import "fmt"

func update(name *string) {
	*name = "Alper"
}

func main() {

	name := "Tesla"
	fmt.Println(name)

	update(&name)
	fmt.Println(name)

}
