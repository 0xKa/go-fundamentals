package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

func (p Person) Hi() string {
	return "Hello, my name is " + p.Name
}

func GoIntro_Ex1() {
	person := Person{
		Name: "Reda",
		Age:  21,
	}

	fmt.Println(person.Hi())
}
