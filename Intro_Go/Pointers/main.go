package main

import "fmt"

type Circle struct {
	X, Y, Radius int
}

func (c *Circle) grow() {
	c.Radius = c.Radius * 2
}

func main() {

	c := Circle{
		X:      1,
		Y:      2,
		Radius: 4,
	}

	fmt.Println(c.Radius)
	c.grow()
	fmt.Println(c.Radius)

}
