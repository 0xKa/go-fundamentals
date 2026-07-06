package practice

import (
	"fmt"
	"strconv"
)

func TypeConversionsEx13() {
	fmt.Println("1. Convert an integer to a float")
	fmt.Println("Numeric types must be converted explicitly before mixing them.")
	fmt.Println("Code: count := 7; average := float64(count) / 2")
	count := 7
	average := float64(count) / 2
	fmt.Printf("Result: average = %.1f, type = %T\n", average, average)

	fmt.Println("\n2. Convert a float to an integer")
	fmt.Println("Converting to int discards the fractional part; it does not round.")
	fmt.Println("Code: price := 19.95; whole := int(price)")
	price := 19.95
	whole := int(price)
	fmt.Printf("Result: whole = %d\n", whole)

	fmt.Println("\n3. Format an integer as text")
	fmt.Println("strconv.Itoa returns the decimal text for an int.")
	fmt.Println("Code: ageText := strconv.Itoa(28)")
	ageText := strconv.Itoa(28)
	fmt.Printf("Result: ageText = %q, type = %T\n", ageText, ageText)

	fmt.Println("\n4. Parse numeric text")
	fmt.Println("strconv.Atoi returns an int and an error that must be checked.")
	fmt.Println(`Code: score, err := strconv.Atoi("42")`)
	score, err := strconv.Atoi("42")
	fmt.Printf("Result: score = %d, err = %v\n", score, err)

	fmt.Println("\n5. strconv function overview")
	fmt.Print(`Functions:

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
`)
}
