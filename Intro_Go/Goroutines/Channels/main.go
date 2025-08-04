package main

import (
	"fmt"
	"math"
)

type circle struct {
	r float64
}

func (c circle) display(ch chan string) {
	ch <- "A Circle"

}

func (c circle) area(ch chan float64) {
	msg := math.Pi * c.r * c.r
	ch <- msg
}

func main() {

	c1 := circle{r: 5}

	ch1 := make(chan float64)
	ch2 := make(chan string)

	go c1.display(ch2)
	go c1.area(ch1)

	fmt.Printf("%.2f\n", <-ch1)
	fmt.Println(<-ch2)
}
