package gobyexample

import (
	"fmt"
	"os"
)

func ShowEnvironmentVariables() {
	// set an env using os.Setenv
	os.Setenv("ZZ-a", "1")
	os.Setenv("ZZ-b", "2")
	os.Setenv("ZZ-c", "3")

	// get an env using os.Getenv
	fmt.Println("ZZ-a:", os.Getenv("ZZ-a"))
	fmt.Println("ZZ-b:", os.Getenv("ZZ-b"))
	fmt.Println("ZZ-c:", os.Getenv("ZZ-c"))

	// Setenv returns an error if the key is empty or contains an '=' character. It also returns an error if the value is empty and the key does not exist.
	ok := os.Setenv("ZZ-test-key", "sk-...")

	if ok != nil {
		fmt.Println("Failed to set environment variable")
		return
	}

	env, exists := os.LookupEnv("ZZ-test-key")
	if !exists {
		fmt.Println("Environment variable ZZ-test-key does not exist")
		return
	}
	fmt.Println("Value of ZZ-test-key:", env)
	// LookupEnv vs Getenv: LookupEnv returns the value and a boolean indicating whether the key exists, while Getenv returns an empty string if the key does not exist.

	// Print all environment variables
	// fmt.Println(os.Environ())

	// Format os.Environ() output
	// for _, e := range os.Environ() {
	// 	// pair := strings.SplitN(e, "=", 2)
	// 	pair := e
	// 	fmt.Println(pair)
	// }
}
