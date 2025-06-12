package main

import (
	"fmt"
)

type Student struct {
	name   string
	grades []int
	age    int
}

func (s *Student) setName(name string) {
	s.name = name
}

func (s Student) getName() string {
	return s.name
}

func (s *Student) setAge(age int) {
	s.age = age
}

func (s Student) getAge() int {
	return s.age
}

func (s Student) getAverageGrade() float32 {
	sum := 0
	for _, grade := range s.grades {
		sum = sum + grade
	}
	return float32(sum) / float32(len(s.grades))
}

func (s *Student) getMaxGrade() int {
	max_grade := 0
	for _, grade := range s.grades {
		if grade > max_grade {
			max_grade = grade
		}
	}
	return max_grade
}

func main() {

	s1 := Student{grades: []int{70, 95, 78, 80, 95, 99, 85}}

	s1.setName("Tim")
	s1.setAge(19)

	name := s1.getName()
	age := s1.getAge()

	fmt.Println("Name :", name)
	fmt.Println("Age :", age)

	averageGrade := s1.getAverageGrade()
	fmt.Println("Average Grade :", averageGrade)

	max_grade := s1.getMaxGrade()
	fmt.Println("Max Grade :", max_grade)

}
