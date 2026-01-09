package main

import (
	"fmt"
)

func main() {

	// variable := make(chan type)
	greeting := make(chan string)

	greetString := "Hello"

	// greeting <- greetString // send data to channel

	go func() {
		greeting <- greetString
		greeting <- "World"

		for _, letter := range "abcde" {
			greeting <- "Alphabet: " + string(letter)
		}
	}()

	// go func() {
	// 	reciever := <-greeting // recieve data from channel
	// 	fmt.Println(reciever)

	// 	reciever = <-greeting
	// 	fmt.Println(reciever)
	// }()

	reciever := <-greeting // recieve data from channel
	fmt.Println(reciever)

	reciever = <-greeting
	fmt.Println(reciever)

	for _ = range 5 {
		rcvr := <-greeting
		fmt.Println(rcvr)
	}

	// time.Sleep(2 * time.Second)
	fmt.Println("End of the program")

}
