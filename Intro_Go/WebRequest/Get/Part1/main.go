package main

import (
	"fmt"
	"io"
	"net/http"
	"strings"
)

func PerformGetRequest(url string) {

	response, err := http.Get(url)

	if err != nil {
		panic(err)
	}

	defer response.Body.Close()

	fmt.Println("Status Code:", response.StatusCode)
	fmt.Println("Content Length:", response.ContentLength)

	var responseString strings.Builder

	content, _ := io.ReadAll(response.Body)

	byteCount, _ := responseString.Write(content)

	fmt.Println("ByteCount is:", byteCount)
	fmt.Println(responseString.String())

	//fmt.Println(content)
	//fmt.Println(string(content))

}

func main() {
	fmt.Println("Web Request GET")
	PerformGetRequest("http://localhost:8000/get")

}
