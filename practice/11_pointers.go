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
	fmt.Println("---")
	fmt.Println("func addBonus(score int) {")
	fmt.Println("    score += 10")
	fmt.Println("}")
	fmt.Println("score := 70")
	fmt.Println("addBonus(score)")
	addBonus(score)
	fmt.Println("---")
	fmt.Println("Output:")
	fmt.Printf(">> score = %d\n", score)

	fmt.Println("\n2. Return a changed value")
	fmt.Println("Store a returned value to update the caller.")
	fmt.Println("---")
	fmt.Println("func withBonus(score int) int {")
	fmt.Println("    return score + 10")
	fmt.Println("}")
	fmt.Println("score = withBonus(score)")
	score = withBonus(score)
	fmt.Println("---")
	fmt.Println("Output:")
	fmt.Printf(">> score = %d\n", score)

	fmt.Println("\n3. Take an address and dereference it")
	scorePointer := &score
	fmt.Println("& takes an address; * accesses its value.")
	fmt.Println("---")
	fmt.Println("scorePointer := &score")
	fmt.Println("*scorePointer = 85")
	*scorePointer = 85
	fmt.Println("---")
	fmt.Println("Output:")
	fmt.Printf(">> score = %d, *scorePointer = %d\n", score, *scorePointer)

	fmt.Println("\n4. Pass a pointer value")
	fmt.Println("A copied pointer still leads to the original value.")
	fmt.Println("---")
	fmt.Println("func addBonusThroughPointer(score *int) {")
	fmt.Println("    *score += 10")
	fmt.Println("}")
	fmt.Println("addBonusThroughPointer(&score)")
	addBonusThroughPointer(&score)
	fmt.Println("---")
	fmt.Println("Output:")
	fmt.Printf(">> score = %d\n", score)

	fmt.Println("\n5. Update a struct through a pointer")
	player := pointerPlayer{Name: "Amina", Score: 90}
	fmt.Println("A struct pointer lets a function update the original struct.")
	fmt.Println("---")
	fmt.Println(`player := pointerPlayer{Name: "Amina", Score: 90}`)
	fmt.Println("func renamePlayer(player *pointerPlayer) {")
	fmt.Println(`    player.Name = "Mona"`)
	fmt.Println("}")
	fmt.Println("renamePlayer(&player)")
	renamePlayer(&player)
	fmt.Println("---")
	fmt.Println("Output:")
	fmt.Printf(">> player.Name = %q\n", player.Name)

	fmt.Println("\n6. Handle a nil pointer")
	var missingScore *int
	fmt.Println("Check for nil before dereferencing.")
	fmt.Println("---")
	fmt.Println("var missingScore *int")
	fmt.Println("if score == nil {")
	fmt.Println("    return false")
	fmt.Println("}")
	fmt.Println("changed := addBonusSafely(missingScore)")
	changed := addBonusSafely(missingScore)
	fmt.Println("---")
	fmt.Println("Output:")
	fmt.Printf(">> changed = %t, missingScore == nil is %t\n", changed, missingScore == nil)

	fmt.Println("\n7. Slices are also passed by value")
	numbers := []int{1, 2, 3}
	fmt.Println("A copied slice shares its underlying array.")
	fmt.Println("---")
	fmt.Println("func replaceFirst(numbers []int) {")
	fmt.Println("    numbers[0] = 99")
	fmt.Println("}")
	fmt.Println("numbers := []int{1, 2, 3}")
	fmt.Println("replaceFirst(numbers)")
	replaceFirst(numbers)
	fmt.Println("---")
	fmt.Println("Output:")
	fmt.Printf(">> numbers = %v\n", numbers)

	fmt.Println("\n8. new returns a pointer to a zero value")
	count := new(int)
	fmt.Println("---")
	fmt.Println("count := new(int)")
	fmt.Println("---")
	fmt.Println("Output:")
	fmt.Printf(">> *count = %d\n", *count)
	fmt.Println()
	fmt.Println("---")
	fmt.Println("*count = 3")
	*count = 3
	fmt.Println("---")
	fmt.Println("Output:")
	fmt.Printf(">> *count = %d\n", *count)
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
