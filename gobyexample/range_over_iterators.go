package gobyexample

import (
	"fmt"
	"iter"
)

// yield is just a keyword for the function that will be called to give the next number in the sequence. The function will return true if it wants to continue yielding numbers, or false if it wants to stop.
func rangeOverIteratorsNumbers(yield func(int) bool) {
	yield(11)
	yield(22)
	yield(33)
}

func rangeOverIteratorsGenFib() iter.Seq[int] {
	return func(yield func(int) bool) {
		a, b := 0, 1

		for {
			if !yield(a) {
				return
			}
			a, b = b, a+b
		}
	}
}

func ShowRangeOverIterators() {

	for n := range rangeOverIteratorsNumbers {
		fmt.Println(n)
	}

	for n := range rangeOverIteratorsGenFib() {
		if n > 1000 { // get feb series nums until 1000
			break
		}
		fmt.Printf("%d ", n)
	}
}
