package main

import "fmt"

func proAdder(values ...int) (int, string) {
	total := 0

	for _, value := range values {
		total += value
	}
	return total, "Hi Pro result function"
}

func main() {

	result, myMessage := proAdder(2, 3, 4, 5, 6)
	fmt.Println(result)
	fmt.Println(myMessage)

}
