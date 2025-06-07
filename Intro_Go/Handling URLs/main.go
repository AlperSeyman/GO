package main

import (
	"fmt"
	"net/url"
)

const myUrl string = "https://example.com:3000/learn?coursename=goLang&paymentid=ghb123ghb"

func main() {
	fmt.Println("Handling URLs")
	fmt.Println(myUrl)

	// parsing URL
	result, _ := url.Parse(myUrl)
	fmt.Println(result.Scheme)
	fmt.Println(result.Host)
	fmt.Println(result.Path)
	fmt.Println(result.Port())
	fmt.Println(result.RawQuery)

	qparams := result.Query()
	fmt.Printf("The type of query params are: %T\n", qparams)

	fmt.Println(qparams["coursename"])

	for _, val := range qparams {
		fmt.Println("Param is :", val)
	}

	partsOfUrl := &url.URL{
		Scheme:  "https",
		Host:    "example.com",
		Path:    "/learn",
		RawPath: "coursename=goLang",
	}

	anotherURL := partsOfUrl.String()

	fmt.Println(anotherURL)
}
