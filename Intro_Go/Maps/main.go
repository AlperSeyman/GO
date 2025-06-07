package main

import (
	"fmt"
)

func main() {

	languages := make(map[string]string)

	languages["JS"] = "Javascript"
	languages["RB"] = "Ruby"
	languages["PY"] = "Python"

	fmt.Println("List of all Languages: ", languages)
	fmt.Println("JS shorts for: ", languages["JS"])
	fmt.Println("RB shorts for: ", languages["RB"])
	fmt.Println("PY shorts for: ", languages["PY"])

	delete(languages, "RB") // delete RB from map

	fmt.Println("List of all Languages: ", languages)

	for _, value := range languages {
		fmt.Println(value)
	}
}
