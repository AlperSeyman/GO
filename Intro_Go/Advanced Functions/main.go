package main

import "fmt"

type EmailBill struct {
	CostInPennies int
}

func adder() func(int) int {
	sum := 0
	return func(i int) int {
		sum = sum + i
		return sum
	}

}

func main() {

	price := EmailBill{
		CostInPennies: 2,
	}

	calculate_price := adder()
	fmt.Println(calculate_price(price.CostInPennies))

}
