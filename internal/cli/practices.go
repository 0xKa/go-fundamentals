package cli

import (
	"fmt"
	"io"
	"strings"

	"go-fundamentals/practice"
)

type practiceItem struct {
	number int
	title  string
	run    func()
}

var practiceCatalog = []practiceItem{
	{number: 1, title: "Go Introduction", run: practice.GoIntroEx1},
	{number: 2, title: "Variables", run: practice.VarsEx2},
	{number: 3, title: "Numeric Types in Go", run: practice.NumbersInGoEx3},
	{number: 4, title: "Printing and Formatting", run: practice.FormattingEx4},
	{number: 5, title: "Arrays and Slices", run: practice.ArraysAndSlicesEx5},
	{number: 6, title: "Go Standard Library", run: practice.StandardLibraryEx6},
	{number: 7, title: "Loops in Go", run: practice.LoopsEx7},
	{number: 8, title: "Booleans and Conditions in Go", run: practice.BooleansAndConditionsEx8},
	{number: 9, title: "Functions in Go", run: practice.FunctionsEx9},
	{number: 10, title: "Maps in Go", run: practice.MapsEx10},
	{number: 11, title: "Pass by Value and Pointers in Go", run: practice.PointersEx11},
	{number: 12, title: "Structs and Custom Types in Go", run: practice.StructsAndCustomTypesEx12},
	{number: 13, title: "Type Conversions in Go", run: practice.TypeConversionsEx13},
	{number: 14, title: "Saving Files in Go", run: practice.SavingFilesEx14},
	{number: 15, title: "Interfaces in Go", run: practice.InterfacesEx15},
	{number: 16, title: "HTTP Servers and JSON Routes in Go", run: practice.HTTPServersAndJSONRoutesEx16},
	{number: 17, title: "Context in Go", run: practice.ContextInGoEx17},
}

func executePractice(number int) error {
	selected, found := findPractice(number)
	if !found {
		return fmt.Errorf("practice %d does not exist", number)
	}

	heading := fmt.Sprintf(
		"========== Practice %02d: %s ==========",
		selected.number,
		selected.title,
	)

	fmt.Printf("%s\n\n", heading)
	selected.run()
	fmt.Printf("\n%s\n", strings.Repeat("=", len(heading)))
	return nil
}

func findPractice(number int) (practiceItem, bool) {
	for _, practice := range practiceCatalog {
		if practice.number == number {
			return practice, true
		}
	}
	return practiceItem{}, false
}

func firstIncomplete(saved progress) (int, bool) {
	for _, practice := range practiceCatalog {
		if !saved.isCompleted(practice.number) {
			return practice.number, true
		}
	}
	return 0, false
}

func followingPractice(number int) (practiceItem, bool) {
	for _, practice := range practiceCatalog {
		if practice.number > number {
			return practice, true
		}
	}
	return practiceItem{}, false
}

func completedPracticeCount(saved progress) int {
	count := 0
	for _, practice := range practiceCatalog {
		if saved.isCompleted(practice.number) {
			count++
		}
	}
	return count
}

func printPracticeList(output io.Writer, saved progress) {
	fmt.Fprintf(
		output,
		"Practices (%d/%d complete):\n",
		completedPracticeCount(saved),
		len(practiceCatalog),
	)

	for _, practice := range practiceCatalog {
		marker := "○"
		if saved.isCompleted(practice.number) {
			marker = "✓"
		}
		fmt.Fprintf(output, "%s %02d: %s\n", marker, practice.number, practice.title)
	}
}
