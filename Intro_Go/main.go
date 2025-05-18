package main

import "fmt"

func main() {

	g := greeter{
		greeting: "Hello",
		name:     "Alper",
	}
	g.greet()
}

type greeter struct {
	greeting string
	name     string
}

func (g greeter) greet() {
	fmt.Println(g.greeting, g.name)
}
