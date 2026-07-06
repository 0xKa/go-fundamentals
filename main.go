package main

import (
	"bufio"
	"fmt"
	"strings"
)

func getInput(r *bufio.Reader, prompt string) (string, error) {
	fmt.Print(prompt)

	input, err := r.ReadString('\n')
	if err != nil {
		return "", err
	}

	return strings.TrimSpace(input), nil
}

func main() {

	reader := bufio.NewReader(strings.NewReader("Reda\nDubai\n"))

	name, err := getInput(reader, "Enter your name: ")
	if err != nil {
		fmt.Println("Error reading input:", err)
		return
	}
	fmt.Printf("Hello, %s!\n", strings.TrimSpace(name))
	// cli.Start()
}
