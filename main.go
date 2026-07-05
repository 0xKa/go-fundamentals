package main

import (
	"fmt"
)

func main() {

	score := new(int)
	fmt.Printf("Initial score: %d\n", *score)

	// startCLI()
}

func startCLI() {

	number, err := getPracticeNumber()
	if err != nil {
		fmt.Println("Invalid practice number.")
		return
	}

	runPractice(number)
}
