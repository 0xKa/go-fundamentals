package main

import (
	"fmt"
)

func main() {

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
