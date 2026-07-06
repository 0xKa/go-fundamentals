# Functions in Go

A function is a named block of code that performs a task. Functions help split a program into small, reusable pieces.

Declare a function with the `func` keyword:

```go
func greet() {
	fmt.Println("Hello, Reda!")
}
```

The code inside a function runs only when the function is called:

```go
greet()
```

Output:

```text
Hello, Reda!
```

Functions are declared at package level, outside other functions. Go does not allow one named function declaration inside another function.

## Parameters and arguments

A parameter is a variable listed in a function declaration. An argument is the value supplied when the function is called.

```go
func greet(name string) {
	fmt.Println("Hello,", name)
}

greet("Mina")
greet("Omar")
```

Output:

```text
Hello, Mina
Hello, Omar
```

Here, `name` is a parameter. `"Mina"` and `"Omar"` are arguments.

Each parameter has a type. When adjacent parameters share a type, the type can be written once:

```go
func add(first, second int) int {
	return first + second
}
```

This is equivalent to writing `first int, second int`.

## Returning a value

The return type appears after the parameter list:

```go
func add(first, second int) int {
	return first + second
}
```

The `return` statement ends the function and sends a value back to the caller:

```go
total := add(7, 5)
fmt.Println("Total:", total)
```

Output:

```text
Total: 12
```

A returned value can be assigned to a variable, printed, or used as part of another expression:

```go
fmt.Println(add(10, 4))
result := add(2, 3) * 2
```

The returned value must match the declared return type.

## Returning multiple values

Go functions can return more than one value. List the return types inside parentheses:

```go
func divide(dividend, divisor float64) (float64, bool) {
	if divisor == 0 {
		return 0, false
	}

	return dividend / divisor, true
}
```

This function returns:

1. the quotient as a `float64`;
2. a `bool` that reports whether the division succeeded.

Receive both values in the same order:

```go
quotient, ok := divide(10, 2)

if ok {
	fmt.Println("Quotient:", quotient)
}
```

Output:

```text
Quotient: 5
```

Multiple return values are commonly used to return a result together with status information. The `value, ok` pattern lets the caller decide how to handle an unsuccessful operation:

```go
quotient, ok := divide(10, 0)

if !ok {
	fmt.Println("Cannot divide by zero.")
}
```

Output:

```text
Cannot divide by zero.
```

## Returning two related results

A function can calculate and return two useful results together:

```go
func minMax(numbers []int) (int, int) {
	smallest := numbers[0]
	largest := numbers[0]

	for _, number := range numbers[1:] {
		if number < smallest {
			smallest = number
		}
		if number > largest {
			largest = number
		}
	}

	return smallest, largest
}
```

Call the function by providing a non-empty slice:

```go
smallest, largest := minMax([]int{8, 3, 12, 5})
fmt.Println("Smallest:", smallest)
fmt.Println("Largest:", largest)
```

Output:

```text
Smallest: 3
Largest: 12
```

This example expects at least one number because it uses the first element as the starting minimum and maximum.

## Ignoring a returned value

Go requires every declared local variable to be used. If only one value from a multiple-value return is needed, assign the unwanted value to the blank identifier `_`:

```go
_, largest := minMax([]int{8, 3, 12, 5})
fmt.Println("Largest only:", largest)
```

Output:

```text
Largest only: 12
```

The blank identifier explicitly tells Go that the value is intentionally ignored.

## Arguments are passed by value

When an ordinary value such as an `int` is passed to a function, the function receives a copy:

```go
func double(number int) int {
	number *= 2
	return number
}

original := 6
doubled := double(original)

fmt.Println("Original:", original)
fmt.Println("Doubled:", doubled)
```

Output:

```text
Original: 6
Doubled: 12
```

Changing the parameter does not change `original`. The function returns the new value instead.

## Function scope

Parameters and variables declared inside a function belong to that function:

```go
func add(first, second int) int {
	total := first + second
	return total
}

// first, second, and total are not available here
```

Keeping variables in the smallest useful scope makes code easier to understand and prevents unrelated functions from changing them accidentally.

## Summary

Functions organize behavior into reusable units with explicit inputs, outputs, and local scope.

- Declare parameter and return types in the function signature.
- Use multiple return values when results naturally belong together.
- Use `_` to intentionally ignore an unwanted returned value.
- Remember that arguments are passed by value.
