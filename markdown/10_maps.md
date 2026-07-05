# Maps in Go

A map stores key-value pairs. Each key identifies one value, which makes maps useful for looking up information quickly.

For example, a map can connect student names to scores:

```go
scores := map[string]int{
	"Amina": 92,
	"Khaled":  84,
}
```

The type is written as `map[keyType]valueType`. In this example:

- each key is a `string`;
- each value is an `int`;
- the complete type is `map[string]int`.

Keys in a map are unique. Assigning another value to an existing key replaces its previous value.

## Creating a map with a literal

A map literal creates a map and provides its initial entries:

```go
scores := map[string]int{
	"Amina": 92,
	"Khaled":  84,
}

fmt.Println("Amina:", scores["Amina"])
```

Output:

```text
Amina: 92
```

Use square brackets and a key to retrieve its value:

```go
value := scores["Amina"]
```

## Creating an empty map with `make`

Use `make` when a map should start empty:

```go
inventory := make(map[string]int)
inventory["notebooks"] = 4
inventory["pens"] = 10
```

The map type still describes both parts:

```go
map[string]int
```

The key type is inside the brackets, and the value type follows them.

## Adding and updating entries

The same syntax adds a new key or updates an existing key:

```go
scores["Reda"] = 95 // add
scores["Khaled"] = 88 // update
```

After these assignments, `"Reda"` has a score of `95`, and `"Khaled"` has a new score of `88`.

Map elements are not variables that can be addressed directly. Update a numeric value by assigning the result back to its key:

```go
scores["Reda"] = scores["Reda"] + 1
```

The shorter increment form also works:

```go
scores["Reda"]++
```

## Looking up a key

A basic lookup returns the value:

```go
score := scores["Reda"]
fmt.Println(score)
```

Output:

```text
95
```

If a key does not exist, a lookup returns the zero value of the map's value type:

```go
fmt.Println(scores["Nora"]) // 0
```

This creates an ambiguity: a score of `0` could be stored in the map, or the key could be absent.

Use the two-value lookup to tell the difference:

```go
score, found := scores["Nora"]

if found {
	fmt.Println("Score:", score)
} else {
	fmt.Println("Nora was not found.")
}
```

Output:

```text
Nora was not found.
```

The second value is a boolean. It is commonly named `ok`, `found`, or `exists`:

```go
score, ok := scores["Reda"]
```

This is another common use of the multiple-return pattern from the functions lesson.

## Deleting entries

Use the built-in `delete` function with a map and a key:

```go
delete(scores, "Khaled")
```

Deleting a key that does not exist is safe and does nothing:

```go
delete(scores, "Unknown")
```

Use `len` to count the current entries:

```go
fmt.Println("Students:", len(scores))
```

Use `clear` to remove every entry:

```go
clear(scores)
fmt.Println(len(scores)) // 0
```

## Iterating over a map

Use `range` to visit each key-value pair:

```go
for name, score := range scores {
	fmt.Println(name, score)
}
```

Map iteration order is not guaranteed. A program should not depend on entries appearing in insertion order or any other fixed order.

When stable output is needed, collect and sort the keys first:

```go
names := make([]string, 0, len(scores))

for name := range scores {
	names = append(names, name)
}

sort.Strings(names)

for _, name := range names {
	fmt.Printf("%s: %d\n", name, scores[name])
}
```

For the complete example, the file needs both packages:

```go
import (
	"fmt"
	"sort"
)
```

## Counting words with a map

A map is a natural way to count how often each word appears:

```go
func countWords(text string) map[string]int {
	counts := make(map[string]int)

	for _, word := range strings.Fields(text) {
		counts[word]++
	}

	return counts
}
```

Call the function and look up each count:

```go
counts := countWords("go makes go simple")

fmt.Println("go:", counts["go"])
fmt.Println("makes:", counts["makes"])
fmt.Println("simple:", counts["simple"])
```

Output:

```text
go: 2
makes: 1
simple: 1
```

Looking up a new word starts with the zero value `0`, so `counts[word]++` works whether the key already exists or not.

## Map key rules

Map keys must be comparable with `==` and `!=`.

Common key types include:

- strings;
- integers;
- booleans;
- arrays;
- structs whose fields are all comparable.

Slices, maps, and functions cannot be map keys because they are not comparable:

```go
// invalid: slices cannot be map keys
// values := map[[]int]string{}
```

Map values do not have the same restriction. A map value can be a slice, another map, or any other type.

## Maps share their underlying data

Assigning a map to another variable does not copy all its entries. Both variables refer to the same underlying map data:

```go
original := map[string]int{"pens": 3}
shared := original

shared["pens"] = 8
fmt.Println(original["pens"])
```

Output:

```text
8
```

For the same reason, a function can update entries in a map passed as an argument, and the caller will see those updates.

## Nil maps

A map's zero value is `nil`:

```go
var scores map[string]int
```

Reading from a nil map is safe and returns the value type's zero value:

```go
fmt.Println(scores["Amina"]) // 0
```

Writing to a nil map causes a runtime panic:

```go
// scores["Amina"] = 92 // panic
```

Initialize the map before adding entries:

```go
scores = make(map[string]int)
scores["Amina"] = 92
```

## Common mistakes

- Writing to a nil map before initializing it.
- Assuming a missing key is present because its zero value was returned.
- Depending on map iteration order.
- Trying to use a slice, map, or function as a key.
- Expecting assignment to create an independent copy of a map.
- Forgetting that assigning to an existing key replaces its value.

Use a map when values need to be found by meaningful, unique keys rather than by numeric positions.
