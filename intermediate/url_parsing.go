package main

import (
	"fmt"
	"net/url"
)

func main() {

	// [scheme://][userinfo@]host[:port][?query][#fragment]

	rawURL := "https://example.com:8080/path?query=params#fragment"

	parseURL, err := url.Parse(rawURL)

	if err != nil {
		fmt.Println("Error parsing URL: ", err)
		return
	}
	fmt.Println("Scheme: ", parseURL.Scheme)
	fmt.Println("Host: ", parseURL.Host)
	fmt.Println("Port: ", parseURL.Port())
	fmt.Println("Path: ", parseURL.Path)
	fmt.Println("Raw Query: ", parseURL.RawQuery)
	fmt.Println("Fragment: ", parseURL.Fragment)

	rawURL1 := "https://example.com/path?name=John&age=30"

	parseURL1, err := url.Parse(rawURL1)
	if err != nil {
		fmt.Println("Error parsing URL: ", err)
		return
	}
	queryParams := parseURL1.Query()
	fmt.Println(queryParams)
	fmt.Println("Name: ", queryParams.Get("name"))
	fmt.Println("Age: ", queryParams.Get("age"))

	// Building URL
	baseURL := &url.URL{
		Scheme: "https",
		Host:   "example.com",
		Path:   "path",
	}

	query := baseURL.Query()
	query.Set("name", "John")
	query.Set("age", "37")
	baseURL.RawQuery = query.Encode()

	fmt.Println("Built URL: ", baseURL.String())

	// Add key-value pairs to the values object
	values := url.Values{}
	values.Add("name", "Jane")
	values.Add("age", "30")
	values.Add("city", "London")
	values.Add("country", "UK")

	fmt.Println(values)
	// Encode
	encodedQuery := values.Encode()
	fmt.Println(encodedQuery)

	// Build URL
	myURL := &url.URL{
		Scheme: "https",
		Host:   "learningGO.com",
		Path:   "info",
	}
	fullUrl := myURL.String() + "?" + encodedQuery
	fmt.Println(fullUrl)
}
