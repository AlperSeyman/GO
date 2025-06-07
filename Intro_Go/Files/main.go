package main

import (
	"fmt"
	"io"
	"os"
)

func readFile(fileName string) {

	databyte, err := os.ReadFile(fileName) // read file

	checkNilErr(err)

	fmt.Println("Text data inside file is \n", string(databyte))

}

func checkNilErr(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {

	content := "This needs to go in a file - LearnGo"

	file, err := os.Create("./mygofile.txt") // cread file current directory

	checkNilErr(err)

	length, err2 := io.WriteString(file, content) // write content into file

	checkNilErr(err2)

	fmt.Println("Length is :", length)
	defer file.Close()

	readFile("./mygofile.txt")
}
