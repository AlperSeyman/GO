package basics

import "fmt"

func main() {

	process()
	fmt.Println("Returned from process")

}

func process() {
	defer func() {
		// if r := recover(); r != nil { // recover captures panic
		r := recover()
		if r != nil {
			fmt.Println("Recovered: ", r)
		}
	}()

	fmt.Println("Start Process")
	panic("Something went wrong")
	fmt.Println("This will never run")
}
