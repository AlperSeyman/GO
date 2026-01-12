package main

import (
	"fmt"
	"io"
	"net/http"
)

func main() {

	// Create a new http client
	client := &http.Client{}

	response, err := client.Get("https://jsonplaceholder.typicode.com/posts/1")
	if err != nil {
		fmt.Println("Error making GET request: ", err)
	}

	defer response.Body.Close()

	// Read and print the response body
	body, err := io.ReadAll(response.Body)
	if err != nil {
		fmt.Println("Error reading response body: ", err)
	}
	fmt.Println(string(body))
}
