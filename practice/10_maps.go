package practice

import (
	"fmt"
	"sort"
	"strings"
)

func MapsEx10() {
	scores := map[string]int{
		"Amina":  92,
		"Khaled": 84,
	}

	fmt.Println("1. Map literal")
	fmt.Println("A map pairs each unique key with a value.")
	fmt.Println("---")
	fmt.Println(`scores := map[string]int{"Amina": 92, "Khaled": 84}`)
	fmt.Println("---")
	fmt.Println("Output:")
	fmt.Printf(">> Amina = %d, Khaled = %d\n", scores["Amina"], scores["Khaled"])

	fmt.Println("\n2. Add and update")
	fmt.Println("Assigning a new key adds it; assigning an existing key replaces its value.")
	fmt.Println("---")
	fmt.Println(`scores["Reda"] = 95`)
	scores["Reda"] = 95
	fmt.Println(`scores["Khaled"] = 88`)
	scores["Khaled"] = 88
	fmt.Println("---")
	fmt.Println("Output:")
	fmt.Printf(">> Reda = %d, Khaled = %d\n", scores["Reda"], scores["Khaled"])

	fmt.Println("\n3. Comma-ok lookup")
	fmt.Println("The second value reports whether the key exists.")
	score, found := scores["Reda"]
	fmt.Println("---")
	fmt.Println(`score, found := scores["Reda"]`)
	fmt.Println("---")
	fmt.Println("Output:")
	fmt.Printf(">> score = %d, found = %t\n", score, found)
	score, found = scores["Nora"]
	fmt.Println()
	fmt.Println("---")
	fmt.Println(`score, found = scores["Nora"]`)
	fmt.Println("---")
	fmt.Println("Output:")
	fmt.Printf(">> score = %d, found = %t\n", score, found)

	fmt.Println("\n4. Delete")
	fmt.Println("delete removes a key and its value from the map.")
	fmt.Println("---")
	fmt.Println(`delete(scores, "Khaled")`)
	delete(scores, "Khaled")
	fmt.Println("---")
	fmt.Println("Output:")
	fmt.Printf(">> len(scores) = %d\n", len(scores))

	fmt.Println("\n5. Iterate with sorted keys")
	names := make([]string, 0, len(scores))
	for name := range scores {
		names = append(names, name)
	}
	sort.Strings(names)

	fmt.Println("Map order is not guaranteed, so sort the keys for stable output.")
	fmt.Println("---")
	fmt.Println("sort.Strings(names)")
	fmt.Println("for _, name := range names {")
	fmt.Println(`    fmt.Println(name, scores[name])`)
	fmt.Println("}")
	fmt.Println("---")
	fmt.Println("Output:")
	for _, name := range names {
		fmt.Printf(">> %s %d\n", name, scores[name])
	}

	fmt.Println("\n6. Count words")
	fmt.Println("Incrementing a missing key starts from its zero value, 0.")
	text := "go makes go simple"
	fmt.Println("---")
	fmt.Printf("text := %q\n", text)
	counts := countWords(text)
	fmt.Println("counts := countWords(text)")
	fmt.Println("---")
	fmt.Println("Output:")
	fmt.Printf(">> go = %d, makes = %d, simple = %d\n", counts["go"], counts["makes"], counts["simple"])

	fmt.Println("\n7. Make and clear")
	fmt.Println("make creates an empty writable map; clear removes all its entries.")
	fmt.Println("---")
	fmt.Println("inventory := make(map[string]int)")
	inventory := make(map[string]int)
	fmt.Println(`inventory["notebooks"] = 4`)
	fmt.Println(`inventory["pens"] = 10`)
	inventory["notebooks"] = 4
	inventory["pens"] = 10
	fmt.Println("clear(inventory)")
	clear(inventory)
	fmt.Println("---")
	fmt.Println("Output:")
	fmt.Printf(">> len(inventory) = %d\n", len(inventory))
}

func countWords(text string) map[string]int {
	counts := make(map[string]int)

	for _, word := range strings.Fields(text) {
		counts[word]++
	}

	return counts
}
