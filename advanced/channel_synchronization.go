package main

import (
	"fmt"
	"time"
)

// func main() {

// 	ch := make(chan struct{})

// 	go func() {
// 		fmt.Println("Working...")
// 		time.Sleep(2 * time.Second)
// 		ch <- struct{}{}
// 	}()

// 	<-ch
// 	fmt.Println("Finished")

// }

// func main() {

// 	ch := make(chan int)

// 	go func() {
// 		ch <- 9 // Blocking until the value is received.
// 		time.Sleep(1 * time.Second)
// 		fmt.Println("Sent value")
// 	}()
// 	value := <-ch // Blocking until a value is sent.
// 	fmt.Println(value)
// 	time.Sleep(2 * time.Second)
// }

// ============== Synchronization Multiple Goroutines ==============
// func main() {

// 	numGoroutines := 3
// 	ch := make(chan int, 3)

// 	for i := range numGoroutines {
// 		go func(id int) {
// 			fmt.Printf("Goroutine %d working...\n", id)
// 			time.Sleep(time.Second)
// 			ch <- id // Sending signal of completion
// 		}(i)
// 	}

// 	for range numGoroutines {
// 		<-ch // wait for each goroutine to finish, wait for all goroutines to signal completion.
// 	}
// 	fmt.Println("All goroutines are complete.")
// }

// ======== Synchronizing Data Exchange ========
// func main() {

// 	data := make(chan string)

// 	go func() {
// 		for i := range 5 {
// 			data <- "Hello " + string('0'+i)
// 			time.Sleep(100 * time.Millisecond)
// 		}
// 		close(data)
// 	}()
// 	// close(data) // Channel closed before Goroutine could send a value to the channel

// 	for value := range data {
// 		fmt.Println("Received value: ", value, ": ", time.Now())
// 	}
// }

func main() {

	coffeChan := make(chan string)

	go func() {
		for i := range 3 {
			time.Sleep(time.Second)
			message := "Coffe" + string('0'+i)

			coffeChan <- message
		}
		close(coffeChan)
	}()

	fmt.Println("Customer is waiting")

	for cup := range coffeChan {
		fmt.Println("Customer received ", cup)
	}

	time.Sleep(500 * time.Millisecond)
	fmt.Println("Customer is leaving the shop.")
}
