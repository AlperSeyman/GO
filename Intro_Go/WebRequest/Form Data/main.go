package main

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
)

func PerformPostFormRequest(myUrl string) {

	// formdata
	data := url.Values{}
	data.Add("firstname", "nikola")
	data.Add("lastname", "tesla")
	data.Add("country", "serbia")
	data.Add("job", "engineer")

	response, err := http.PostForm(myUrl, data)
	if err != nil {
		panic(err)
	}
	defer response.Body.Close()

	content, _ := io.ReadAll(response.Body)

	fmt.Println(string(content))
}

func main() {
	PerformPostFormRequest("http://localhost:8080/postform")
}
