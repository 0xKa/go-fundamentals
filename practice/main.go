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
	var name1 string = "Reda"

	var name2 = "Ali"

	name3 := "Ahmed"

	var name4 string
	name4 = "Mohammed"

	fmt.Println(name1)
	fmt.Println(name2)
	fmt.Println(name3)
	fmt.Println(name4)

}
