package basics

import "fmt"

func main() {

	// 	Simple iteration over a range
	// for i := 0; i < 5; i++ {
	// 	fmt.Println(i)
	//  }

	// iterate over collection
	// numbers := []int{1, 2, 3, 4, 5, 6}
	// for index, value := range numbers {
	// 	fmt.Printf("Index: %d, Value: %d\n", index, value)
	// }

	// for i := 0; i < 10; i++ {
	// 	if i%2 == 0 {
	// 		continue -> contiune the loop but skip the rest of lines/statements
	// 	}
	// 	fmt.Println("Odd Number:", i)
	// 	if i == 5 {
	// 		break -> break out of the loop
	// 	}
	// }

	rows := 5
	// Outer loop
	for i := 1; i <= rows; i++ {
		// Inner loop for spaces before starts
		for j := 1; j <= rows-i; j++ {
			fmt.Print(" ")
		}
		// inner loop for starts
		for k := 1; k <= 2*i-1; k++ {
			fmt.Print("*")
		}
		fmt.Println() // Move to the next line
	}
}
