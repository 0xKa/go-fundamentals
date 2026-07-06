# Type Conversions in Go

A type conversion creates a value of one type from a value of another compatible type.

Go requires most conversions to be explicit:

```go
count := 7
average := float64(count) / 2
```

The conversion syntax is `TargetType(value)`. Here, `float64(count)` creates a `float64` value from `count`.

## Converting numeric types

Numeric types do not mix automatically:

```go
count := 7
half := 2.0

// result := count / half // invalid: int and float64 are different types
result := float64(count) / half
```

Output:

```text
3.5
```

Making the conversion visible prevents the program from silently choosing how values should change.

Converting a floating-point value to an integer discards its fractional part:

```go
price := 19.95
whole := int(price)

fmt.Println(whole)
```

Output:

```text
19
```

This truncates toward zero; it does not round. Numeric conversions can also lose precision or exceed the destination type's range, so choose the destination type deliberately.

## Converting numbers and strings

Text containing digits is not automatically a number. Use the `strconv` package to convert between numeric values and their text representations.

Use `strconv.Itoa` to format an `int` as a string:

```go
age := 28
ageText := strconv.Itoa(age)

fmt.Printf("%q\n", ageText)
```

Output:

```text
"28"
```

Use `strconv.Atoi` to parse a string as an `int`:

```go
input := "42"
score, err := strconv.Atoi(input)
if err != nil {
	fmt.Println("Please enter a valid number")
	return
}

fmt.Println(score)
```

Output:

```text
42
```

Parsing can fail because user input might not contain a valid number, so always check the returned error.

`string(65)` is not the way to produce `"65"`. It converts the number to the Unicode character with that code point:

```go
fmt.Println(string(65))
```

Output:

```text
A
```

Use `strconv.Itoa(65)` when the desired result is `"65"`.

###  `strconv` functions overview


```go
Atoi         = string → int
Itoa         = int → string

ParseInt     = string → int64, more control
FormatInt    = int64 → string, more control

ParseFloat   = string → float64
FormatFloat  = float64 → string

ParseBool    = string → bool
FormatBool   = bool → string

ParseUint    = string → uint64
FormatUint   = uint64 → string

ParseX       = string → value
FormatX      = value → string
```

## Summary

Conversions make a deliberate change between compatible types, while `strconv` parses and formats the textual representation of values.

- Use `TargetType(value)` for explicit compatible conversions.
- Converting a float to an integer truncates the fractional part.
- Use `strconv` when translating between numbers, booleans, and strings.
- Always check errors returned by parsing functions.
