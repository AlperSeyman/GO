package main

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	FirstName string  `json:"first_name"`
	Age       int     `json:"age,omitempty"`
	Email     string  `json:"email,omitempty"`
	Address   Address `json:"address"`
}

type Address struct {
	City  string `json:"city"`
	State string `json:"state"`
}

type Employee struct {
	Fullname string  `json:"full_name"`
	EmpID    string  `json:"emp_id"`
	Age      int     `json:"age"`
	Address  Address `json:"address"`
}

func main() {

	person1 := Person{
		FirstName: "John",
	}

	// Marshalling
	jsonData1, err := json.Marshal(person1)
	if err != nil {
		fmt.Println("Error marshalling to JSON: ", err)
		return
	}
	fmt.Println(string(jsonData1))

	person2 := Person{
		FirstName: "Jane",
		Age:       29,
		Email:     "jane@example.com",
		Address: Address{
			City:  "New York",
			State: "NY",
		},
	}
	jsonData2, err := json.Marshal(person2)
	if err != nil {
		fmt.Println("Error marshalling to JSON: ", err)
		return
	}

	fmt.Println(string(jsonData2))

	// Unmarshalling
	jsonData3 := `{"full_name":"Jenny Doe", "emp_id":"0004", "age":30, "address":{"city":"San Jose", "state":"CA"}}`

	var employee Employee
	err = json.Unmarshal([]byte(jsonData3), &employee)
	if err != nil {
		fmt.Println("Error Unmarshalling JSON: ", err)
	}
	fmt.Println("Employee JSON data: ", employee)

	listOfCityState := []Address{
		{City: "New York", State: "NY"},
		{City: "San Jose", State: "CA"},
		{City: "Las Vegas", State: "NV"},
		{City: "Clearwater", State: "FL"},
	}

	fmt.Println(listOfCityState)
	city_data, err := json.Marshal(listOfCityState)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("City Data: ", string(city_data))

	// Handling unknown JSON structures
	jsonData4 := `{"name":"Jon", "age":30, "address":{"city":"New York", "state":"NY"}}`

	var unknown_data map[string]interface{} // it can be replaced by any
	err = json.Unmarshal([]byte(jsonData4), &unknown_data)
	if err != nil {
		fmt.Println("Handling unknown JSON structures error: ", err)
	}
	fmt.Println("Decoded Unknown JSON: ", unknown_data)
	fmt.Println("Decoded Unknown JSON: ", unknown_data["name"])
	fmt.Println("Decoded Unknown JSON: ", unknown_data["age"])
	fmt.Println("Decoded Unknown JSON: ", unknown_data["address"])
}
