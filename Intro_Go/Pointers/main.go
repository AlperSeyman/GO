package main

import "fmt"

func main() {

	thing2 := [5]float64{1, 2, 3, 4, 5}
	fmt.Println(square(&thing2))
	fmt.Println(thing2)

}

func square(thing2 *[5]float64) [5]float64 {
	for i, _ := range thing2 {
		thing2[i] = thing2[i] * thing2[i]
	}
	return *thing2
}
