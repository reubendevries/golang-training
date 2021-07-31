package main

import "fmt"

type Car struct {
	NumberOfTires int
	IsLuxery      bool
	BucketSeats   bool
	Make          string
	Model         string
	Year          int
}

func main() {
	// this is an example of a struct
	myCar := Car{
		NumberOfTires: 4,
		IsLuxery:      true,
		BucketSeats:   true,
		Make:          "Volvo",
		Model:         "XC90",
		Year:          2019,
	}

	// an example of how we would print out an struct.
	fmt.Printf("My car is a %d %s %s\n", myCar.Year, myCar.Make, myCar.Model)

	// This an array
	var myStrings [3]string
	myStrings[0] = "cat"
	myStrings[1] = "dog"
	myStrings[2] = "fish"

	// this is a for loop for an array
	i := 0
	for i < len(myStrings) {
		fmt.Printf("The element is %s\n", myStrings[i])
		i++
	}
}
