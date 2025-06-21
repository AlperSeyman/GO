package main

import (
	"fmt"
	"io"
	"net/http"
)

const url = "https://jsonplaceholder.typicode.com/posts"

func main() {

	fmt.Println("Web request")
	response, err := http.Get(url)
	if err != nil {
		panic(err)
	}

	body, err := io.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}
	sb := string(body)
	fmt.Println(sb)

}
