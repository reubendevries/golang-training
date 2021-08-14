package main

import (
	"fmt"
	"strings"
)

func main() {

	//						1		  2			3		  4
	//			  0123456789012345678901234567890123456789012345
	newString := "Go is a great programming language. Go for it!"

	if strings.Contains(newString, "Go") {
		newString = strings.ReplaceAll(newString, "Go", "Golang")
		//newString = strings.Replace(newString, "Go", "Golang", 1)
	}

	fmt.Println(newString)

	// String Comparision -- works due to run value, a string is just
	if "Alpha" > "Absolute" {
		fmt.Println("Alpha is greater than Absolute")
	} else {
		fmt.Println("Alpha is not greater than Absolute")
	}

	badEmail := " me@here.com "
	badEmail = strings.TrimSpace(badEmail)
	fmt.Printf("=%s=\n", badEmail)
	fmt.Println()
}
