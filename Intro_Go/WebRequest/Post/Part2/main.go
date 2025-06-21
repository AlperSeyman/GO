package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

const url = "https://postman-echo.com/post"

func main() {

	// The first step is to encode our JSON data so it can return data in byte format.
	// To do this, we use the Marshal function that Go’s JSON package provides.
	// Next, we convert the encoded JSON data to a type implemented by the io.Reader interface.
	// We simply use the NewBuffer function for this, passing in the encoded JSON data as an argument.
	// The NewBuffer function returns a value of type buffer, which we can then pass onto the Post function:

	postBody, _ := json.Marshal(map[string]string{
		"name":  "Toby",
		"email": "Tboy@example.com",
	})

	responseBody := bytes.NewBuffer(postBody)
	response, err := http.Post(url, "application/json", responseBody)
	if err != nil {
		panic(err)
	}

	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}

	sb := string(body)
	fmt.Println(sb)
}
