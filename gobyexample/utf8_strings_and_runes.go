package gobyexample

import (
	"fmt"
	"unicode/utf8"
)

func ShowUTF8StringsAndRunes() {
	const s = "السلام عليكم 😊👍"

	fmt.Println(s)

	fmt.Println("Bytes:")
	for i := range len(s) {
		fmt.Printf("%x ", s[i])
	}

	fmt.Println("\nlength:", len(s))
	fmt.Println("Rune count:", utf8.RuneCountInString(s))

	for idx, runeValue := range s {
		fmt.Printf("%#U starts at %d\n", runeValue, idx)
	}
}
