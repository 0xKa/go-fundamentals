# Variables, Strings, and Numbers in Go

[Corresponding Go file](../practice/02_vars.go)

### Variables

4 ways to declare variables in Go:

```go
    // 1. Explicit type declaration
	var name1 string = "Reda"

    // 2. Implicit type declaration (type inferred from the value)
	var name2 = "Ali"

    // 3. Short variable declaration (only inside functions)
	name3 := "Ahmed"

    // 4. Declare variable first, then assign a value
	var name4 string
	name4 = "Mohammed"

	fmt.Println(name1) 
	fmt.Println(name2)
	fmt.Println(name3)
	fmt.Println(name4)

```
