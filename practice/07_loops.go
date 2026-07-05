package practice

import "fmt"

func LoopsEx7() {
	fmt.Println("Counted loop:")
	for number := 1; number <= 5; number++ {
		fmt.Println(number)
	}

	fmt.Println("\nCondition-only loop:")
	countdown := 3
	for countdown > 0 {
		fmt.Println(countdown)
		countdown--
	}
	fmt.Println("Go!")

	fmt.Println("\nRange over a slice:")
	languages := []string{"Go", "Python", "JavaScript"}
	for index, language := range languages {
		fmt.Printf("%d: %s\n", index, language)
	}

	fmt.Println("\nRange over a string:")
	for index, character := range "Go✓" {
		fmt.Printf("byte %d: %c\n", index, character)
	}

	fmt.Println("\nContinue and break:")
	for number := 1; number <= 10; number++ {
		if number > 6 {
			break
		}
		if number%2 != 0 {
			continue
		}

		fmt.Println(number)
	}
}
