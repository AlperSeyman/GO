package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	// reader := bufio.NewReader(strings.NewReader("Hello, bufio packageee!\nHow are you doing?")) // return reader object

	// Reading byte slice
	// data := make([]byte, 20)
	// n, err := reader.Read(data)
	// if err != nil {
	// 	fmt.Println("Error reading: ", err)
	// 	return
	// }
	// fmt.Printf("Read %d bytes: %s\n", n, data[:n])

	// s, err := reader.ReadString('\n')
	// if err != nil {
	// 	fmt.Println("Error reading string:", err)
	// 	return
	// }
	// fmt.Println("Read string:", s)

	writer := bufio.NewWriter(os.Stdout) // os.Stdout --> Printing text to the screen.

	// Writing byte slice
	data := []byte("Hello, bufio package!\n")
	n, err := writer.Write(data)
	if err != nil {
		fmt.Println("Error writing:", err)
		return
	}
	fmt.Printf("Wrote %d bytes\n", n)

	err = writer.Flush()
	if err != nil {
		fmt.Println("Error flushing writer: ", err)
		return
	}

	// Writing string
	str := "This is a string\n"
	n, err = writer.WriteString(str)
	if err != nil {
		fmt.Println("Error writing string:", err)
		return
	}
	fmt.Printf("Wrote %d bytes\n", n)

	// Flush the buffer
	err = writer.Flush()
	if err != nil {
		fmt.Println("Error flushing writer: ", err)
		return
	}
}
