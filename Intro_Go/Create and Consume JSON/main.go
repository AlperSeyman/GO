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

// Create json data
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

// consume json data
func decodeJson() {

	jsonDataFromWeb := []byte(`
		{
        	"coursename": "React",
            "price": 299,
            "website": "Udemy",
            "tags": ["web-dev","js"]
        }
	`)

	var courses Course

	// validate that data is json or not
	checkValid := json.Valid(jsonDataFromWeb) // return bool, true or false

	if checkValid { // if json is valid
		fmt.Println("JSON was valid")
		json.Unmarshal(jsonDataFromWeb, &courses)
		fmt.Printf("%#v", courses)
		fmt.Println()
		fmt.Println("****************")
	} else {
		fmt.Println("JSON WAS NOT VALID")
	}

	// some cases where we just want to add data to key value
	var myOnlineData map[string]any // value can be int, string, array, slice, struct
	json.Unmarshal(jsonDataFromWeb, &myOnlineData)
	fmt.Printf("%#v", myOnlineData)

	fmt.Println()
	for k, v := range myOnlineData {
		fmt.Printf("Key is %v and Value is %v and Type is: %T\n", k, v, v)
	}
}

func main() {
	//encodeJson()
	decodeJson()
}
