package main

import (
	"fmt"
	"strings"
)

func main() {
	myString := "This is a clear EXAMPLE of why we search in one case only."

	lowercase := strings.ToLower(myString)

	fmt.Println(strings.Title(lowercase))
}
