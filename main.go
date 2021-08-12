package main

import (
	"errors"
	"fmt"
)

func main() {
	a := 12
	b := 6
	c, err := divideTwoNumbers(a, b)
	if err != nil {
		fmt.Print(err)
	} else {
		if c == 2 {
			fmt.Println("We found 2.")
		}
	}
}

func divideTwoNumbers(x, y int) (int, error) {
	if y == 0 {
		return 0, errors.New("cannot divde by zero.\n")
	}
	return x / y, nil
}
