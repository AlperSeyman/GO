package basics

import (
	"errors"
	"fmt"
)

func main() {

	// func functionName(parameter1 type1, parameter2 type2, ....) return (returnType1, returnType2, ...){
	// code block
	// return returnType1, returnType2
	// }

	q, r := divide(10, 3)
	fmt.Println("quotient: ", q)
	fmt.Println("remainder: ", r)

	result, err := compare(3, 2)
	if err != nil {
		fmt.Println("Error: ", err)
	} else {
		fmt.Println(result)
	}

}

func divide(a, b int) (quotient int, remainder int) {
	quotient = a / b
	remainder = a % b
	return
}

func compare(a, b int) (string, error) {
	if a > b {
		return "a is greater than b", nil
	} else if b > a {
		return "b is greater than a", nil
	} else {
		return "", errors.New("Unable to compare which is greater")
	}

}
