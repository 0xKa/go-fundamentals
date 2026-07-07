package gobyexample

import (
	"errors"
	"fmt"
)

func errorHandlingF(arg int) (int, error) {
	if arg == 42 {

		return -1, errors.New("can't work with 42")
	}

	return arg + 3, nil
}

var errorHandlingErrOutOfTea = errors.New("no more tea available")
var errorHandlingErrPower = errors.New("can't boil water")

func errorHandlingMakeTea(arg int) error {
	switch arg {
	case 2:
		return errorHandlingErrOutOfTea
	case 4:
		return fmt.Errorf("making tea: %w", errorHandlingErrPower)
	}
	return nil
}

func ShowErrorHandling() {
	for _, i := range []int{7, 42} {

		if r, e := errorHandlingF(i); e != nil {
			fmt.Println("f failed:", e)
		} else {
			fmt.Println("f worked:", r)
		}
	}

	for i := range 5 {
		if err := errorHandlingMakeTea(i); err != nil {

			if errors.Is(err, errorHandlingErrOutOfTea) {
				fmt.Println("We should buy new tea!")
			} else if errors.Is(err, errorHandlingErrPower) {
				fmt.Println("Now it is dark.")
			} else {
				fmt.Printf("unknown error: %s\n", err)
			}
			continue
		}

		fmt.Println("Tea is ready!")
	}
}
