package practice

import "fmt"

type pointerPlayer struct {
	Name  string
	Score int
}

func PointersEx11() {
	fmt.Println("1. A non-pointer argument is copied")
	score := 70
	fmt.Println("The function changes only its copy.")
	fmt.Println("Code: func addBonus(score int) { score += 10 }")
	fmt.Println("Code: score := 70; addBonus(score)")
	addBonus(score)
	fmt.Printf("Result: score = %d\n", score)

	fmt.Println("\n2. Return a changed value")
	fmt.Println("Store a returned value to update the caller.")
	fmt.Println("Code: func withBonus(score int) int { return score + 10 }")
	fmt.Println("Code: score = withBonus(score)")
	score = withBonus(score)
	fmt.Printf("Result: score = %d\n", score)

	fmt.Println("\n3. Take an address and dereference it")
	scorePointer := &score
	fmt.Println("& takes an address; * accesses its value.")
	fmt.Println("Code: scorePointer := &score")
	fmt.Println("Code: *scorePointer = 85")
	*scorePointer = 85
	fmt.Printf("Result: score = %d, *scorePointer = %d\n", score, *scorePointer)

	fmt.Println("\n4. Pass a pointer value")
	fmt.Println("A copied pointer still leads to the original value.")
	fmt.Println("Code: func addBonusThroughPointer(score *int) { *score += 10 }")
	fmt.Println("Code: addBonusThroughPointer(&score)")
	addBonusThroughPointer(&score)
	fmt.Printf("Result: score = %d\n", score)

	fmt.Println("\n5. Update a struct through a pointer")
	player := pointerPlayer{Name: "Amina", Score: 90}
	fmt.Println("A struct pointer lets a function update the original struct.")
	fmt.Println(`Code: player := pointerPlayer{Name: "Amina", Score: 90}`)
	fmt.Println(`Code: func renamePlayer(player *pointerPlayer) { player.Name = "Mona" }`)
	fmt.Println("Code: renamePlayer(&player)")
	renamePlayer(&player)
	fmt.Printf("Result: player.Name = %q\n", player.Name)

	fmt.Println("\n6. Handle a nil pointer")
	var missingScore *int
	fmt.Println("Check for nil before dereferencing.")
	fmt.Println("Code: var missingScore *int")
	fmt.Println("Code: if score == nil { return false }")
	fmt.Println("Code: changed := addBonusSafely(missingScore)")
	changed := addBonusSafely(missingScore)
	fmt.Printf("Result: changed = %t, missingScore == nil is %t\n", changed, missingScore == nil)

	fmt.Println("\n7. Slices are also passed by value")
	numbers := []int{1, 2, 3}
	fmt.Println("A copied slice shares its underlying array.")
	fmt.Println("Code: func replaceFirst(numbers []int) { numbers[0] = 99 }")
	fmt.Println("Code: numbers := []int{1, 2, 3}; replaceFirst(numbers)")
	replaceFirst(numbers)
	fmt.Printf("Result: numbers = %v\n", numbers)

	fmt.Println("\n8. new returns a pointer to a zero value")
	count := new(int)
	fmt.Println("Code: count := new(int)")
	fmt.Printf("Result: *count = %d\n", *count)
	fmt.Println("Code: *count = 3")
	*count = 3
	fmt.Printf("Result: *count = %d\n", *count)
}

func addBonus(score int) {
	score += 10
}

func withBonus(score int) int {
	return score + 10
}

func addBonusThroughPointer(score *int) {
	*score += 10
}

func renamePlayer(player *pointerPlayer) {
	player.Name = "Mona"
}

func addBonusSafely(score *int) bool {
	if score == nil {
		return false
	}

	*score += 10
	return true
}

func replaceFirst(numbers []int) {
	numbers[0] = 99
}
