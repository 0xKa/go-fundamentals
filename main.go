package main

import "fmt"

func main() {

	nums := []int{1, 2, 3, 4, 5}
	fmt.Println("Numbers:", nums)

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
