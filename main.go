package main

import "fmt"

func main() {

	nums := [5]int{1, 2, 3, 4, 5}
	part := nums[1:3] // slice from index 1 to 3 (4 is exclusive)
	fmt.Println("Numbers:", nums)
	fmt.Println("Part:", part)

	part = append(part, 6, 7, 8, 9) // append 6 to the slice
	fmt.Println("After appending 6:")
	fmt.Println("Numbers:", nums)
	fmt.Println("Part:", part)

	// startCLI()
}

func sum(nums [5]int) int {
	total := 0
	for _, n := range nums {
		total += n
	}
	return total
}

func startCLI() {

	number, err := getPracticeNumber()
	if err != nil {
		fmt.Println("Invalid practice number.")
		return
	}

	runPractice(number)
}
