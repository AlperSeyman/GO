package main

import "fmt"

type PicnicItem struct {
	Name     string
	Quantity int
}

func countTotalItems(basket []PicnicItem) int {
	total := 0
	for _, item := range basket {
		total = total + item.Quantity
	}
	return total
}

func main() {

	basket := []PicnicItem{
		{Name: "Sandwich", Quantity: 3},
		{Name: "Apple", Quantity: 4},
		{Name: "Water Bottle", Quantity: 2},
	}

	totalItems := countTotalItems(basket)

	for i, item := range basket {
		fmt.Printf("Item %d: %s - Quantity: %d\n", i+1, item.Name, item.Quantity)
	}
	fmt.Printf("Total number of items: %d\n", totalItems)

}
