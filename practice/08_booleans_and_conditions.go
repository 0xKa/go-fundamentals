package practice

import "fmt"

func BooleansAndConditionsEx8() {
	isLearning := true
	var hasFinished bool

	fmt.Println("Boolean values:")
	fmt.Println("Learning:", isLearning)
	fmt.Println("Finished:", hasFinished)

	age := 20
	hasTicket := true
	isMember := false

	fmt.Println("\nComparisons and logical operators:")
	fmt.Println("Adult:", age >= 18)
	fmt.Println("Can enter:", age >= 18 && hasTicket)
	fmt.Println("Gets discount:", age < 18 || isMember)
	fmt.Println("Ticket missing:", !hasTicket)

	score := 74
	fmt.Println("\nConditional grade:")
	if score >= 90 {
		fmt.Println("Grade: A")
	} else if score >= 75 {
		fmt.Println("Grade: B")
	} else if score >= 60 {
		fmt.Println("Grade: C")
	} else {
		fmt.Println("Grade: F")
	}

	fmt.Println("\nShort statement:")
	if length := len("gopher"); length > 5 {
		fmt.Println("Long word:", length)
	} else {
		fmt.Println("Short word:", length)
	}

	day := "Saturday"
	fmt.Println("\nSwitch:")
	switch day {
	case "Saturday", "Sunday":
		fmt.Println("Weekend")
	case "Friday":
		fmt.Println("Almost the weekend")
	default:
		fmt.Println("Weekday")
	}
}
