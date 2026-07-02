package main

import "fmt"

func NumbersInGo_Ex3() {
	// Signed integer types.
	var wholeNumber int = 21
	var smallNumber int8 = 127
	var mediumNumber int16 = 32_767
	var unicodeNumber int32 = 2_147_483_647
	var largeNumber int64 = 9_223_372_036_854_775_807

	// Unsigned integer types.
	var positiveNumber uint = 67
	var byteNumber uint8 = 255
	var smallPositiveNumber uint16 = 65_535
	var mediumPositiveNumber uint32 = 4_294_967_295
	var largePositiveNumber uint64 = 18_446_744_073_709_551_615

	// Floating-point types.
	var decimalNumber float32 = 3.14
	var preciseDecimalNumber float64 = 3.141592653589793

	// Complex number types.
	var complexNumber64 complex64 = complex(3, 4)
	var complexNumber128 complex128 = complex(5, 6)

	// Other numeric types.
	var character rune = 'G'
	var memoryAddress uintptr = 0x1234

	fmt.Printf("%-10s: %d\n", "int", wholeNumber)
	fmt.Printf("%-10s: %d\n", "int8", smallNumber)
	fmt.Printf("%-10s: %d\n", "int16", mediumNumber)
	fmt.Printf("%-10s: %d\n", "int32", unicodeNumber)
	fmt.Printf("%-10s: %d\n", "int64", largeNumber)

	fmt.Println()
	fmt.Printf("%-10s: %d\n", "uint", positiveNumber)
	fmt.Printf("%-10s: %d\n", "uint8", byteNumber)
	fmt.Printf("%-10s: %d\n", "uint16", smallPositiveNumber)
	fmt.Printf("%-10s: %d\n", "uint32", mediumPositiveNumber)
	fmt.Printf("%-10s: %d\n", "uint64", largePositiveNumber)

	fmt.Println()
	fmt.Printf("%-10s: %f\n", "float32", decimalNumber)
	fmt.Printf("%-10s: %.15f\n", "float64", preciseDecimalNumber)

	fmt.Println()
	fmt.Printf("%-10s: %v\n", "complex64", complexNumber64)
	fmt.Printf("%-10s: %v\n", "complex128", complexNumber128)

	fmt.Println()
	fmt.Printf("%-10s: %c (Unicode code point %d)\n", "rune", character, character)
	fmt.Printf("%-10s: %#x\n", "uintptr", memoryAddress)
}
