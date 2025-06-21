package main

import (
	"fmt"
	"io"
	"net/http"
	"strings"
)

func PerformPostJsonRequest(url string) {

	// fake json paylod
	requestBody := strings.NewReader(`
		{
			"coursename":"Go Lang",
			"price" : 0,
			"platform" : "udemy"
		}
	`)

	response, err := http.Post(url, "application/json", requestBody)

	if err != nil {
		panic(err)
	}

	defer response.Body.Close()

	content, _ := io.ReadAll(response.Body)

	fmt.Println(string(content))
}

func main() {
	fmt.Println("Web Request POST")
	PerformPostJsonRequest("http://localhost:8000/post")
}
