package main

import (
	"fmt"
)

type Person struct {
	firstName string
	lastName  string
	age       int
	address   Address
	PhoneHomeCell
}

type PhoneHomeCell struct {
	home string
	cell string
}

type Address struct {
	city    string
	country string
}

func main() {

	p1 := Person{
		firstName: "Joe",
		lastName:  "Doe",
		age:       30,
		address: Address{
			city:    "London",
			country: "U.K",
		},
		PhoneHomeCell: PhoneHomeCell{
			home: "3424234234",
			cell: "456546456456",
		},
	}

	p2 := Person{
		firstName: "Jane",
		age:       25,
	}

	p2.address.city = "New York"
	p2.address.country = "USA"

	// Anonymous Structs
	user := struct {
		username string
		email    string
	}{
		username: "user123",
		email:    "pseudoemail@example.org",
	}
	fmt.Println(user.username)
	fmt.Println(user.email)

	fmt.Println(p1.firstName)
	fmt.Println(p1.fullName())
	fmt.Println("P1 age before the incremenet: ", p1.age)
	p1.incrementAgeByOne()
	fmt.Println("P1 age after the increment: ", p1.age)

}

func (p Person) fullName() string {
	return p.firstName + " " + p.lastName
}

func (p *Person) incrementAgeByOne() {
	p.age++
}
