package main

import "fmt"

func main() {
	// does one number divide exactly into another?

	// x := 12
	// y := 5

	// if x%y == 0 {
	// 	fmt.Println(y, "divides exactly into", x)
	// } else {
	// 	fmt.Println(y, "doesn't divide exactly into", x)
	// }

	// thisMonth := 4
	// fmt.Println("The month after", thisMonth, "is", thisMonth+1)

	for m := 1; m <= 12; m++ {
		fmt.Println("The month after", m, "is", m%12+1)
	}
}
