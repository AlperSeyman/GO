package main

import "fmt"

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}

func main() {

	user := User{
		Name:   "Sich",
		Email:  "moin@mion.com",
		Status: true,
		Age:    25,
	}

	fmt.Println(user)
	fmt.Printf("User details are %+v\n", user)
	fmt.Println("Name:", user.Name)

}
