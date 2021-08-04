package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	i := 1000

	for i > 100 {
		// get a random number between 1 and 1001
		i = rand.Intn(1000) + 1
		if i > 100 {
			fmt.Println("i is", i, "so loop keeps going")
		}
	}

	fmt.Println("Got", i, "and broke out of loop.")
}
