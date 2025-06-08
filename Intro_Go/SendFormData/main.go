package main

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
)

func PerformPostFormRequest() {

	const myUrl = "http://localhost:8000/postform"

	// form data
	data := url.Values{}
	data.Add("firstname", "moin")
	data.Add("lastname", "ost")
	data.Add("email", "moin.ost@dev.com")

	response, err := http.PostForm(myUrl, data)

	if err != nil {
		panic(err)
	}

	defer response.Body.Close()

	content, _ := io.ReadAll(response.Body)

	fmt.Println(string(content))
}

func main() {
	PerformPostFormRequest()
}
