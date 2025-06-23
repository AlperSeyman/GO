package main

import (
	"encoding/json"
	"fmt"
)

type course struct {
	Name     string   `json:"coursename"`
	Price    int      `json:"price"`
	Platform string   `json:"platform"`
	Password string   `json:"-"`
	Tags     []string `json:"tags,omitempty"`
}

func DecodeJson() {
	jsonDataFromWeb := []byte(`
		{
                "coursename": "Go Lang",
                "price": 100,
                "platform": "Udemy",
                "tags": ["web-dev","api"]
        }
	`)

	course := course{}

	checkValid := json.Valid(jsonDataFromWeb)
	if checkValid {
		fmt.Println("JSON was valid")
		json.Unmarshal(jsonDataFromWeb, &course)

		fmt.Printf("%+v\n", course)
	} else {
		fmt.Println("JSON was invalid")
	}
}

func main() {
	DecodeJson()
}
