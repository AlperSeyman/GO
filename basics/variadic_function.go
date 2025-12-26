package basics

import "fmt"

func main() {

	// ... Ellipsis
	// func functionName(param1 type1, param2 type2, param3 ...type3) returnType{
	// function body
	// }

	// result := sum(1, 2, 3)
	// statement, total := sum("Sum: ", 1, 2, 3)
	// fmt.Println(statement, total)

	numbers := []int{1, 2, 3, 4, 5, 9}
	sequence, total := sum(1, numbers...)
	fmt.Printf("Sequnce: %d, Total: %d", sequence, total)
}

func sum(sequence int, nums ...int) (int, int) {
	total := 0
	for _, num := range nums {
		total += num
	}
	return sequence, total
}
