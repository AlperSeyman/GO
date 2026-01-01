package main

import "fmt"

type person struct {
	name string
	age  int
}

type employee struct {
	employeeInfo person // Embedded struct Named field (inheritance)
	person              // Embedded struct Anonymous (inheritance)
	empID        string
	salary       float64
}

func (p person) introduce() {
	fmt.Printf("Hi! I'm %s and I'm %d years old.\n", p.name, p.age)
}

func (e employee) introduce() {
	fmt.Printf("Hi! I'm %s, employee ID: %s, and I earn %.2f.\n", e.name, e.empID, e.salary)
}

func main() {

	emp := employee{
		person: person{
			name: "John",
			age:  30,
		},
		empID:  "E001",
		salary: 50000,
	}

	fmt.Println("Name: ", emp.name) // Accessing the embedded struct field emp.person.name
	fmt.Println("Age: ", emp.age)   // Same as above
	fmt.Println("Employee: ", emp.empID)
	fmt.Println("Salary: ", emp.salary)

	emp.introduce()
}
