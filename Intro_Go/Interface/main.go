package main

import (
	"fmt"
	"math"
)

type Shape interface {
	area() float64
}

type Circle struct {
	radius float64
}

type Rect struct {
	width  float64
	height float64
}

func (r Rect) area() float64 {
	return r.width * r.height
}

func (c Circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func getArea(s Shape) float64 {
	return s.area()
}

func main() {

	c1 := Circle{radius: 4.5}
	r1 := Rect{width: 5, height: 7}

	shapes := []Shape{c1, r1}

	fmt.Println(getArea(shapes[0]))
}
