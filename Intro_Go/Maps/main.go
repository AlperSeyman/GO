package main

import (
	"fmt"
)

func getNameCounts(names []string) map[rune]map[string]int {
	counts := make(map[rune]map[string]int)
	for _, name := range names {
		if len(name) == 0 {
			continue
		}
		firstChar := rune(name[0])
		_, ok := counts[firstChar]
		if !ok {
			counts[firstChar] = make(map[string]int)
		}
		counts[firstChar][name]++
	}
	return counts
}

func main() {

	names := []string{"Bily", "Bily", "Bob", "Joe"}

	fmt.Print(getNameCounts(names))
}
