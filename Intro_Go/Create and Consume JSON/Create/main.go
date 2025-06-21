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

func EncodeJson() {

	myCourses := []course{
		{Name: "Go Lang", Price: 100, Platform: "Udemy", Password: "go123", Tags: []string{"web-dev", "api"}},
		{Name: "Python", Price: 99, Platform: "Udemy", Password: "py123", Tags: []string{"ml-ai", "numpy", "pandas"}},
		{Name: "Java", Price: 119, Platform: "Udemy", Password: "class123", Tags: []string{"oop", "spring-boots"}},
		{Name: "C/C++", Price: 0, Platform: "Youtube", Password: "", Tags: nil},
	}

	finalJson, err := json.MarshalIndent(myCourses, "", "\t")
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s\n", finalJson)
}

func main() {
	EncodeJson()
}
