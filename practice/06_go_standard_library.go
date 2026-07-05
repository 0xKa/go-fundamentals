package practice

import (
	"fmt"
	"sort"
	"strings"
)

func StandardLibraryEx6() {
	message := "  go makes small tools powerful  "
	clean := strings.TrimSpace(message)
	words := strings.Fields(clean)

	fmt.Printf("Raw: %q\n", message)
	fmt.Println("Clean:", clean)
	fmt.Println("Uppercase:", strings.ToUpper(clean))
	fmt.Println(`Contains "tools":`, strings.Contains(clean, "tools"))
	fmt.Println("Words:", words)
	fmt.Println("Joined:", strings.Join(words, " | "))

	scores := []int{42, 7, 19, 7}
	fmt.Println("\nScores before sorting:", scores)
	sort.Ints(scores)
	fmt.Println("Scores after sorting:", scores)

	names := []string{"Zoe", "Ahmed", "Mina"}
	fmt.Println("\nNames before sorting:", names)
	sort.Strings(names)
	fmt.Println("Names after sorting:", names)

	fruit := strings.Split("pear,apple,banana", ",")
	sort.Strings(fruit)
	fmt.Println("\nCombined example:", strings.Join(fruit, ", "))
}
