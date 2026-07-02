package main

import "fmt"

func Vars_Ex2() {
	// Explicit type declaration.
	var name1 string = "Reda"

	// Implicit type declaration (the type is inferred from the value).
	var name2 = "Ali"

	// Short variable declaration (only available inside functions).
	name3 := "Ahmed"

	// Declare a variable first, then assign a value.
	var name4 string
	name4 = "Mohammed"

	fmt.Println(name1)
	fmt.Println(name2)
	fmt.Println(name3)
	fmt.Println(name4)
}
