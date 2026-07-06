# Booleans and Conditions in Go

A boolean represents one of two values: `true` or `false`. Go uses the `bool` type for boolean values.

Booleans let a program ask questions and choose which code to run:

```go
isLearning := true
hasFinished := false

fmt.Println(isLearning)
fmt.Println(hasFinished)
```

Output:

```text
true
false
```

The zero value of a `bool` is `false`:

```go
var isReady bool
fmt.Println(isReady) // false
```

## Comparison operators

Comparisons produce boolean results.

| Operator | Meaning |
| --- | --- |
| `==` | equal to |
| `!=` | not equal to |
| `<` | less than |
| `<=` | less than or equal to |
| `>` | greater than |
| `>=` | greater than or equal to |

For example:

```go
age := 20

fmt.Println(age >= 18)
fmt.Println(age == 21)
fmt.Println(age != 16)
```

Output:

```text
true
false
true
```

Use `==` to compare values. A single `=` assigns a value to a variable:

```go
score := 80     // assignment
passed := score >= 60 // comparison
```

Values being compared must have compatible types. Ordered comparisons such as `<` and `>` work with numbers and strings, while booleans can only be compared with `==` and `!=`.

## Logical operators

Logical operators combine or reverse boolean expressions.

| Operator | Meaning | Result is `true` when |
| --- | --- | --- |
| `&&` | AND | both expressions are true |
| `\|\|` | OR | at least one expression is true |
| `!` | NOT | the expression is false |

Example:

```go
age := 20
hasTicket := true
isMember := false

canEnter := age >= 18 && hasTicket
getsDiscount := age < 18 || isMember

fmt.Println(canEnter)
fmt.Println(getsDiscount)
fmt.Println(!hasTicket)
```

Output:

```text
true
false
false
```

Parentheses can make a longer condition easier to understand:

```go
canJoin := (age >= 18 && hasTicket) || isMember
```

Go evaluates `&&` and `||` from left to right and stops when the result is already known. This is called short-circuit evaluation.

## Choosing with `if`

An `if` statement runs a block only when its condition is `true`:

```go
temperature := 32

if temperature > 30 {
	fmt.Println("It is hot.")
}
```

Go does not require parentheses around the condition, but braces are required.

The condition must be a boolean expression. Go does not treat values such as `1` or non-empty strings as `true`:

```go
if temperature > 30 {
	fmt.Println("valid condition")
}

// if 1 {
//     fmt.Println("invalid in Go")
// }
```

## Adding `else`

Use `else` to run a different block when the condition is `false`:

```go
score := 74

if score >= 60 {
	fmt.Println("Passed")
} else {
	fmt.Println("Try again")
}
```

Output:

```text
Passed
```

In Go, `else` must appear on the same line as the closing brace of the preceding block.

## Checking several conditions

Use `else if` when there are several mutually exclusive choices:

```go
score := 74

if score >= 90 {
	fmt.Println("Grade: A")
} else if score >= 75 {
	fmt.Println("Grade: B")
} else if score >= 60 {
	fmt.Println("Grade: C")
} else {
	fmt.Println("Grade: F")
}
```

Output:

```text
Grade: C
```

Conditions are checked from top to bottom. Go runs the first block whose condition is `true`, then skips the remaining branches. Put more specific or higher thresholds first.

## A short statement in `if`

An `if` statement can begin with a short statement followed by a semicolon:

```go
if length := len("gopher"); length > 5 {
	fmt.Println("Long word:", length)
} else {
	fmt.Println("Short word:", length)
}
```

Output:

```text
Long word: 6
```

The `length` variable is available in the `if` condition and every branch, but it is not available after the complete `if` statement. This keeps a temporary value close to where it is used.

## Choosing with `switch`

A `switch` statement is useful when one value is compared against several choices:

```go
day := "Saturday"

switch day {
case "Saturday", "Sunday":
	fmt.Println("Weekend")
case "Friday":
	fmt.Println("Almost the weekend")
default:
	fmt.Println("Weekday")
}
```

Output:

```text
Weekend
```

Go automatically stops after the matching case, so cases do not need `break`. Multiple values can share a case by separating them with commas.

A `switch` without an expression can replace a long chain of boolean conditions:

```go
temperature := 32

switch {
case temperature > 30:
	fmt.Println("Hot")
case temperature >= 20:
	fmt.Println("Warm")
default:
	fmt.Println("Cool")
}
```

As with an `if` chain, the first matching case runs.

## Summary

Boolean expressions let a program compare values and choose which code should run.

- Conditions in Go must evaluate to `bool`.
- Combine conditions with `&&`, `||`, and `!`.
- Use `if` for ranges or unrelated expressions.
- Use `switch` when several branches form a clear set of choices.
