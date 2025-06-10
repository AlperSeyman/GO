package main

import "fmt"

func test(x int, y int) (z1 int, z2 int) {
	defer fmt.Println("hello")
	z1 = x + y
	z2 = x - y
	fmt.Println("before return")
	return // return z1 and z2
}

func main() {
	ans1, ans2 := test(14, 7)
	fmt.Println(ans1, ans2)
}
