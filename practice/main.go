package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

func (p Person) Hi() string {
	return "Hello, my name is " + p.Name
}

func main() {
	p := Person{Name: "Reda", Age: 21}
	fmt.Println(p.Hi())
}
