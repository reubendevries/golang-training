package main

import "fmt"

func main() {
	second := 31
	minute := 1

	if minute < 59 && second+1 > 59 {
		minute++
	}

	fmt.Println(minute)
}
