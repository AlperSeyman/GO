package intermediate

import "fmt"

func main() {

	// sequence := adder() // adder returns function.

	// fmt.Println(sequence())
	// fmt.Println(sequence())
	// fmt.Println(sequence())
	// fmt.Println(sequence())

	subtracter := func() func(int) int {
		countDown := 99
		return func(i int) int {
			countDown -= i
			return countDown
		}
	}()
	// Using the closure substracter
	fmt.Println(subtracter(1))
	fmt.Println(subtracter(2))
	fmt.Println(subtracter(3))
	fmt.Println(subtracter(4))
	fmt.Println(subtracter(5))
}

func adder() func() int {

	i := 0

	fmt.Println("previous value of i: ", i)

	return func() int { // closure
		i++
		fmt.Println("added 1 to i")
		return i
	}
}
