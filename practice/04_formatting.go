package practice

import "fmt"

func FormattingEx4() {
	fmt.Print("Print: Hello", " World! ", 1, 2.5)
	fmt.Println()
	fmt.Println("Println:", "adds spaces between its arguments.")

	name := "Ahmed"
	age := 30

	fmt.Printf("Printf: Hello, %s! You are %d years old.\n", name, age)

	formattedGreeting := fmt.Sprintf(
		"Sprintf: Hello, %s! You are %d years old.",
		name,
		age,
	)
	fmt.Println(formattedGreeting)

	formattedLine := fmt.Sprintln("Sprintln:", name, "is", age, "years old.")
	fmt.Print(formattedLine)

	person := Person{Name: name, Age: age}
	number := 42
	decimal := 3.14159
	letter := 'G'

	fmt.Println("\nFormatting verbs:")
	fmt.Printf("%-6s: %s\n", "%s", "Go")
	fmt.Printf("%-6s: %d\n", "%d", number)
	fmt.Printf("%-6s: %f\n", "%f", decimal)
	fmt.Printf("%-6s: %.2f\n", "%.2f", decimal)
	fmt.Printf("%-6s: %t\n", "%t", true)
	fmt.Printf("%-6s: %v\n", "%v", person)
	fmt.Printf("%-6s: %+v\n", "%+v", person)
	fmt.Printf("%-6s: %#v\n", "%#v", person)
	fmt.Printf("%-6s: %T\n", "%T", person)
	fmt.Printf("%-6s: 100%%\n", "%%")
	fmt.Printf("%-6s: %p\n", "%p", &number)
	fmt.Printf("%-6s: %e\n", "%e", decimal)
	fmt.Printf("%-6s: %g\n", "%g", decimal)
	fmt.Printf("%-6s: %x\n", "%x", number)
	fmt.Printf("%-6s: %o\n", "%o", number)
	fmt.Printf("%-6s: %c\n", "%c", letter)
	fmt.Printf("%-6s: %q\n", "%q", "Go")
	fmt.Printf("%-6s: %U\n", "%U", letter)
	fmt.Printf("%-6s: %b\n", "%b", number)
}
