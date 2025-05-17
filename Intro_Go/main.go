package main

import (
	"fmt"
)

func main() {

	statePopulations := map[string]int{ // map[key]value
		"California": 39250017,
		"Texas":      27862589,
		"Florida":    20612439,
		"New York":   19745289,
	}

	for _, v := range statePopulations {
		fmt.Println(v)
	}
}
