package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func main() {

	wg.Add(2)

	go task1()
	go task2()

	wg.Wait()

}

func task1() {
	fmt.Println("Task 1")
	wg.Done()
}

func task2() {
	fmt.Println("Task 2")
	wg.Done()
}
