# Pass by Value and Pointers in Go

A pointer is a value that stores the memory address of another value. Pointers let code work with the same value from more than one place.

The most important rule is:

> Go always passes arguments by value.

That includes pointer arguments. A function receives a copy of a pointer, but the copied pointer still holds the same address. The function can therefore change the value at that address.

## Passing a non-pointer value

When an `int`, `string`, `bool`, array, or struct is passed to a function, the parameter receives a copy:

```go
func addBonus(score int) {
	score += 10
}

score := 70
addBonus(score)
fmt.Println(score)
```

Output:

```text
70
```

The function changes only its local `score` parameter. Returning the new value is often the clearest approach when the function calculates a result:

```go
func withBonus(score int) int {
	return score + 10
}

score = withBonus(score)
fmt.Println(score)
```

Output:

```text
80
```

## Creating and reading a pointer

Use `&` to get a value's address:

```go
score := 70
scorePointer := &score
```

The type of `scorePointer` is `*int`, read as “pointer to an `int`.”

Use `*` to follow, or dereference, a pointer and access the value stored at its address:

```go
fmt.Println(*scorePointer)
*scorePointer = 85
fmt.Println(score)
```

Output:

```text
70
85
```

The two uses of `*` have related meanings:

- `*int` describes a pointer type;
- `*scorePointer` accesses the pointed-to value.

## Passing a pointer to a function

A function can accept an address and update the value at that address:

```go
func addBonusThroughPointer(score *int) {
	*score += 10
}

score := 70
addBonusThroughPointer(&score)
fmt.Println(score)
```

Output:

```text
80
```

`&score` creates a pointer containing the address of `score`. The function receives a copy of that pointer. Both pointer values contain the same address, so `*score` inside the function refers to the caller's variable.

Go is not switching to “pass by reference” here. The pointer itself is still passed by value.

Reassigning the copied pointer does not redirect a pointer in the caller:

```go
func redirect(pointer *int) {
	other := 100
	pointer = &other
}

value := 10
pointer := &value
redirect(pointer)
fmt.Println(*pointer)
```

Output:

```text
10
```

Changing `*pointer` changes the shared value. Changing `pointer` changes only the function's local pointer copy.

## Pointers to structs

Passing a struct value copies the struct:

```go
type Player struct {
	Name  string
	Score int
}

func renameCopy(player Player) {
	player.Name = "Mona"
}
```

Use a pointer when the function should update the original struct:

```go
func rename(player *Player) {
	player.Name = "Mona"
}

player := Player{Name: "Amina", Score: 90}
rename(&player)
fmt.Println(player.Name)
```

Output:

```text
Mona
```

Go allows `player.Name` instead of the more explicit `(*player).Name` when selecting a field through a struct pointer.

## Nil pointers

The zero value of a pointer is `nil`, which means it does not point to a value:

```go
var scorePointer *int
fmt.Println(scorePointer == nil)
```

Output:

```text
true
```

Dereferencing a nil pointer causes a runtime panic. Check it before using `*` when nil is possible:

```go
func addBonusSafely(score *int) bool {
	if score == nil {
		return false
	}

	*score += 10
	return true
}
```

A pointer can represent an optional value, but a separate boolean or a small result type is sometimes clearer.

## Creating a pointer with `new`

The built-in `new` function allocates a zero value and returns its address:

```go
count := new(int)
fmt.Println(*count)

*count = 3
fmt.Println(*count)
```

Output:

```text
0
3
```

The type of `count` is `*int`. In everyday Go, taking the address of a local variable or struct literal is often more readable:

```go
count := 0
countPointer := &count

playerPointer := &Player{Name: "Amina"}
```

It is also safe for a function to return the address of a local variable. Go keeps the value alive for as long as the returned pointer needs it:

```go
func startingScore() *int {
	score := 50
	return &score
}
```

## Slices and maps still follow pass by value

Slices and maps can seem as if they are passed by reference, but their values are copied like every other argument.

A copied slice value still refers to the same underlying array. Changing an existing element is visible to the caller:

```go
func replaceFirst(numbers []int) {
	numbers[0] = 99
}

numbers := []int{1, 2, 3}
replaceFirst(numbers)
fmt.Println(numbers)
```

Output:

```text
[99 2 3]
```

A copied map value refers to the same underlying map data, so adding or updating an entry is also visible to the caller.

The slice variable itself is still copied. If a function appends and changes its local slice length, the caller does not receive that new slice value. Return it instead:

```go
func appendNumber(numbers []int, number int) []int {
	return append(numbers, number)
}

numbers = appendNumber(numbers, 4)
```

Arrays behave differently from slices: passing an array copies all its elements unless a pointer to the array is passed.

## When to use a pointer

Use a pointer when:

- a function needs to mutate the caller's value;
- several parts of a program must share the identity of one value;
- copying a large struct would be needlessly expensive;
- `nil` has a useful, clearly documented meaning.

Prefer a non-pointer value when:

- the value is small and naturally copied;
- the function only needs to read it;
- returning a transformed value makes the data flow clearer;
- independent copies are safer and easier to reason about.

Strings, slices, maps, functions, channels, and interfaces are small descriptor-like values. Passing them directly is normal; a pointer to one of these types is rarely needed.

## Pointer limitations and common mistakes

Go keeps pointer operations deliberately small and safe:

- pointers can be compared with `nil`;
- pointers of the same type can be compared for equality;
- Go does not support pointer arithmetic;
- a pointer must point to a value of the correct type.

Common mistakes include:

- expecting a function to change a non-pointer argument;
- calling `*pointer` when `pointer` might be nil;
- confusing `*T`, a pointer type, with `*pointer`, a dereference;
- passing a pointer when returning a value would be simpler;
- assuming slices and maps break the pass-by-value rule;
- taking pointers to values without a clear mutation, identity, or optional-value need.

Pointers are most useful when they make shared mutation or identity explicit. For ordinary calculations, values and return statements usually keep the code simpler.
