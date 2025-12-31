package main

import "fmt"

type Rectangle struct {
	length float64
	width  float64
}

type MyInt int

func (m MyInt) IsPositive() bool {
	return m > 0
}

func main() {

	r := Rectangle{length: 10, width: 9}

	area := r.Area()
	fmt.Println("Area: ", area)

	fmt.Println("L before:", r.length)
	r.Scale(2)
	fmt.Println("L after:", r.length)

	num1 := MyInt(-5)
	num2 := MyInt(8)

	fmt.Println(num1.IsPositive())
	fmt.Println(num2.IsPositive())

}

// Method with value receiver
func (r Rectangle) Area() float64 {
	return r.length * r.width
}

// Method with pointer receiver
func (r *Rectangle) Scale(factor float64) {
	r.length = r.length * factor
	r.width = r.width * factor
}
