package main

import (
	"fmt"
	"go-fundamentals/practice"
	"os"
	"strconv"
	"strings"
)

type Practice struct {
	Title string
	Run   func()
}

var practices = map[int]Practice{
	1: {Title: "Go Introduction", Run: practice.GoIntroEx1},
	2: {Title: "Variables", Run: practice.VarsEx2},
	3: {Title: "Numeric Types in Go", Run: practice.NumbersInGoEx3},
	4: {Title: "Printing and Formatting", Run: practice.FormattingEx4},
	5: {Title: "Arrays and Slices", Run: practice.ArraysAndSlicesEx5},
	6: {Title: "Go Standard Library", Run: practice.StandardLibraryEx6},
	7: {Title: "Loops in Go", Run: practice.LoopsEx7},
	8: {Title: "Booleans and Conditions in Go", Run: practice.BooleansAndConditionsEx8},
	9: {Title: "Functions in Go", Run: practice.FunctionsEx9},
}

func runPractice(number int) {
	practice, exists := practices[number]
	if !exists {
		fmt.Printf("Practice %d does not exist.\n", number)
		return
	}

	heading := fmt.Sprintf("========== Practice %02d: %s ==========", number, practice.Title)

	fmt.Printf("%s\n\n", heading)
	practice.Run()
	fmt.Printf("\n%s\n", strings.Repeat("=", len(heading)))
}

func getPracticeNumber() (int, error) {
	// check whether a number was provided as a cli argument.
	if len(os.Args) > 1 {
		return strconv.Atoi(os.Args[1])
	}

	// otherwise, prompt the user for a number.
	var number int
	fmt.Print("Enter a practice number: ")
	_, err := fmt.Scan(&number)

	return number, err
}
