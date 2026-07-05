package practice

import "fmt"

func FunctionsEx9() {
	fmt.Println("Calling a function:")
	greet("Reda")

	fmt.Println("\nParameters and a return value:")
	total := add(7, 5)
	fmt.Println("Total:", total)

	fmt.Println("\nMultiple return values:")
	quotient, ok := divide(10, 2)
	if ok {
		fmt.Println("Quotient:", quotient)
	}

	quotient, ok = divide(10, 0)
	if !ok {
		fmt.Println("Cannot divide by zero.")
	}

	fmt.Println("\nTwo related results:")
	numbers := []int{8, 3, 12, 5}
	smallest, largest := minMax(numbers)
	fmt.Println("Smallest:", smallest)
	fmt.Println("Largest:", largest)

	_, largestOnly := minMax(numbers)
	fmt.Println("Largest only:", largestOnly)

	fmt.Println("\nArguments are passed by value:")
	original := 6
	doubled := double(original)
	fmt.Println("Original:", original)
	fmt.Println("Doubled:", doubled)
}

func greet(name string) {
	fmt.Println("Hello,", name)
}

func add(first, second int) int {
	return first + second
}

func divide(dividend, divisor float64) (float64, bool) {
	if divisor == 0 {
		return 0, false
	}

	return dividend / divisor, true
}

func minMax(numbers []int) (int, int) {
	smallest := numbers[0]
	largest := numbers[0]

	for _, number := range numbers[1:] {
		if number < smallest {
			smallest = number
		}
		if number > largest {
			largest = number
		}
	}

	return smallest, largest
}

func double(number int) int {
	number *= 2
	return number
}
