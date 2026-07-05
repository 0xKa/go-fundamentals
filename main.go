package main

import (
	"fmt"
)

type StudentID int

type Address struct {
	Street string
	City   string
}

type Student struct {
	ID      StudentID
	Name    string
	Age     int
	Address Address
	Marks   map[string]float64
}

func (s Student) PrintDetails() {
	fmt.Println("Student ID:", s.ID)
	fmt.Println("Name:", s.Name)
	fmt.Println("Age:", s.Age)
	fmt.Println("Address:", s.Address.Street, ",", s.Address.City)
	fmt.Println("Marks:")
	for subject, mark := range s.Marks {
		fmt.Printf("%s: %.2f\n", subject, mark)
	}
}

func (s *Student) UpdateName(newName string) {
	s.Name = newName
}

func main() {

	s := Student{
		ID:   StudentID(222),
		Name: "Reda",
		Age:  21,
		Address: Address{
			Street: "123 Main St",
			City:   "Muscat",
		},
		Marks: map[string]float64{
			"Math":    85.5,
			"Science": 92.0,
			"English": 78.5,
		},
	}

	s2 := s
	s2.UpdateName("Ali")
	s2.PrintDetails()

	// cli.Start()
}
