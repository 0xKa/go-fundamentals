package main

import (
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"

	"go-fundamentals/practice"

	"charm.land/huh/v2"
)

type Practice struct {
	Title string
	Run   func()
}

var practices = map[int]Practice{
	1:  {Title: "Go Introduction", Run: practice.GoIntroEx1},
	2:  {Title: "Variables", Run: practice.VarsEx2},
	3:  {Title: "Numeric Types in Go", Run: practice.NumbersInGoEx3},
	4:  {Title: "Printing and Formatting", Run: practice.FormattingEx4},
	5:  {Title: "Arrays and Slices", Run: practice.ArraysAndSlicesEx5},
	6:  {Title: "Go Standard Library", Run: practice.StandardLibraryEx6},
	7:  {Title: "Loops in Go", Run: practice.LoopsEx7},
	8:  {Title: "Booleans and Conditions in Go", Run: practice.BooleansAndConditionsEx8},
	9:  {Title: "Functions in Go", Run: practice.FunctionsEx9},
	10: {Title: "Maps in Go", Run: practice.MapsEx10},
	11: {Title: "Pass by Value and Pointers in Go", Run: practice.PointersEx11},
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

	// Map iteration is unordered, so sort the practice numbers.
	numbers := make([]int, 0, len(practices))
	for number := range practices {
		numbers = append(numbers, number)
	}
	sort.Ints(numbers)

	options := make([]huh.Option[int], 0, len(numbers))
	for _, number := range numbers {
		practice := practices[number]
		label := fmt.Sprintf("%02d: %s", number, practice.Title)
		options = append(options, huh.NewOption(label, number))
	}

	var selected int
	err := huh.NewSelect[int]().
		Title("Choose a practice").
		Options(options...).
		Value(&selected).
		Run()

	return selected, err
}
