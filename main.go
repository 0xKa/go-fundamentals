package main

import (
	"fmt"
	"strings"
)

func getInitialsOfEachWord(name string, separator string) (initials []string, count int) {
	if separator == "" {
		separator = " "
	}

	words := strings.Split(name, separator)
	count = 0
	for _, word := range words {
		if word != "" {
			initials = append(initials, string(word[0]))
			count++
		}
	}
	return initials, count
}

func main() {
	name := "Mohammed Ahmed Khaled"
	initials, count := getInitialsOfEachWord(name, " ")
	fmt.Printf("Initials: %s, Count: %d\n", initials, count)

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
