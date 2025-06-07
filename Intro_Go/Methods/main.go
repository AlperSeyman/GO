package main

import "fmt"

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}

func (user User) GetStatus() {
	fmt.Println("Is user active: ", user.Status)
}

func (user User) NewEmail() {
	user.Email = "test@go.dev"
	fmt.Println("Email of this user is:", user.Email)
}

func main() {

	user1 := User{
		Name:   "Wall",
		Email:  "wall@a.com",
		Status: true,
		Age:    34,
	}

	fmt.Println("Email of this user is:", user1.Email)
	user1.GetStatus()
	user1.NewEmail()
	fmt.Println("Email of this user is:", user1.Email)

}
