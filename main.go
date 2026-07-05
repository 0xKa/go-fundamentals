package main

import (
	"fmt"
)

func main() {

	attempts := 0
	password := "secret"
	attemptsAllowed := 3

	fmt.Println("Enter the correct password to continue:")
	for {
		print("Password: ")
		var input string
		fmt.Scanln(&input)

		if input == password {
			fmt.Println("Access granted.")
			break
		} else {
			fmt.Println("Access denied. Try again, you have", attemptsAllowed-1-attempts, "attempts left.")
		}

		attempts++

		if attempts == attemptsAllowed {
			fmt.Println("Too many failed attempts. Access denied.")
			return
		}
	}

	startCLI()
}

func startCLI() {

	number, err := getPracticeNumber()
	if err != nil {
		fmt.Println("Invalid practice number.")
		return
	}

	runPractice(number)
}
