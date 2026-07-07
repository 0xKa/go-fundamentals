package gobyexample

import (
	"fmt"
	"os"
)

func ShowEnvironmentVariables() {
	os.Setenv("ZZ-a", "1")
	os.Setenv("ZZ-b", "2")
	os.Setenv("ZZ-c", "3")
	fmt.Println("ZZ-a:", os.Getenv("ZZ-a"))
	fmt.Println("ZZ-b:", os.Getenv("ZZ-b"))
	fmt.Println("ZZ-c:", os.Getenv("ZZ-c"))

	ok := os.Setenv("ZZ-test-key", "sk-...")

	if ok != nil {
		fmt.Println("Failed to set environment variable")
		return
	}

	// Get the value of the environment variable
	value := os.Getenv("ZZ-test-key")
	fmt.Println("Value of ZZ-test-key:", value)

	// Print all environment variables
	// fmt.Println(os.Environ())

	// Format os.Environ() output
	// for _, e := range os.Environ() {
	// 	// pair := strings.SplitN(e, "=", 2)
	// 	pair := e
	// 	fmt.Println(pair)
	// }
}
