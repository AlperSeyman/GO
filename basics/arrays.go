package basic

import "fmt"

func main() {

	// var arrayName [size]elementType

	// var numbers [5]int

	// fruits := [4]string{"Apple", "Banana", "Orange", "Grapes"}

	// for i := 0; i < len(fruits); i++ {
	// 	fmt.Println(fruits[i])
	// }

	// for index, value := range fruits {
	// 	fmt.Println(index, value)
	// }

	// for _, value := range fruits {
	// 	fmt.Println(value)
	// }

	// var matrix [3][3]int = [3][3]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}
	// for i := 0; i < len(matrix); i++ {
	// 	for j := 0; j < len(matrix[i]); j++ {
	// 		fmt.Print(matrix[i][j], " ")
	// 	}
	// }

	orginalArray := [3]int{1, 2, 3}
	copiedArray := orginalArray

	fmt.Println(copiedArray)

}
