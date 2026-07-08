package cli

import (
	"bytes"
	"strings"
	"testing"
)

func TestFirstIncompletePractice(t *testing.T) {
	saved := progress{Completed: []int{1, 2}}

	number, found := firstIncomplete(saved)
	if !found || number != 3 {
		t.Fatalf("got number %d, found %t; want number 3", number, found)
	}
}

func TestFollowingPractice(t *testing.T) {
	next, found := followingPractice(3)
	if !found || next.number != 4 {
		t.Fatalf("got %+v, found %t; want practice 4", next, found)
	}

	if next, found := followingPractice(17); found {
		t.Fatalf("got practice %+v after final practice", next)
	}
}

func TestPracticeCatalogIsOrdered(t *testing.T) {
	for index := 1; index < len(practiceCatalog); index++ {
		previous := practiceCatalog[index-1]
		current := practiceCatalog[index]
		if previous.number >= current.number {
			t.Fatalf("practice %d appears before practice %d", previous.number, current.number)
		}
	}
}

func TestPrintPracticeListIncludesProgress(t *testing.T) {
	var output bytes.Buffer
	printPracticeList(&output, progress{Completed: []int{1}})

	text := output.String()
	for _, expected := range []string{
		"Practices (1/17 complete):",
		"✓ 01: Go Introduction",
		"○ 02: Variables",
	} {
		if !strings.Contains(text, expected) {
			t.Errorf("output does not contain %q:\n%s", expected, text)
		}
	}
}
