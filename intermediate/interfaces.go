package main

import (
	"fmt"
	"math"
)

type rect struct {
	width, heigth float64
}

type circle struct {
	radius float64
}

type geometry interface {
	area() float64
	perim() float64
}

func (r rect) area() float64 {
	return r.heigth * r.width
}

func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func (r rect) perim() float64 {
	return 2 * (r.heigth + r.width)
}

func (c circle) perim() float64 {
	return 2 * math.Pi * c.radius
}

func (c circle) diamater() float64 {
	return 2 * c.radius
}

func measure(g geometry) {
	fmt.Println(g)
	fmt.Println(g.area())
	fmt.Println(g.perim())
}

func value(i interface{}) { // interface{} means that any type-value
	fmt.Println(i)
}

func printType(i interface{}) {
	switch i.(type) {
	case int:
		fmt.Println("Type: Int")
	case string:
		fmt.Println("Type: String")
	default:
		fmt.Println("Type: Unknown")
	}
}

func myPrinter(i ...interface{}) { // ...interface{} means that any types-values
	for _, v := range i {
		fmt.Println(v)
	}
}

func main() {

	r := rect{width: 3, heigth: 4}
	c := circle{radius: 5}

	measure(r)
	measure(c)

	value("Nikola")

	printType(4)
	printType("John")
	printType(false)

	myPrinter("Tesla", 5, true, 2.2323)
}
