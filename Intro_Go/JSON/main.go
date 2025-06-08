package main

import (
	"encoding/json"
	"fmt"
)

type Course struct {
	Name     string   `json:"coursename"`
	Price    int      `json:"price"`
	Platform string   `json:"website"`
	Password string   `json:"-"`
	Tags     []string `json:"tags,omitempty"`
}

func encodeJson() {

	courses := []Course{
		{Name: "React", Price: 299, Platform: "Udemy", Password: "abc123", Tags: []string{"web-dev", "js"}},
		{Name: "Python", Price: 300, Platform: "OnlineCode", Password: "xyz456", Tags: []string{"machine-learning", "data-science"}},
		{Name: "Angular", Price: 249, Platform: "Udemy", Password: "işlkjh123", Tags: nil},
	}

	//package this data as JSON data
	finalJson, err := json.MarshalIndent(courses, "", "\t")

	if err != nil {
		panic(err)
	}

	fmt.Printf("%s\n", finalJson)
}

func main() {
	encodeJson()
}
