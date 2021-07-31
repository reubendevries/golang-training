package main

import (
	"fmt"
	"sort"
)

func main() {
	var animals []string
	animals = append(animals, "dog")
	animals = append(animals, "fish")
	animals = append(animals, "cat")
	animals = append(animals, "horse")

	fmt.Println(animals)

	// a way to iterate through and print all animals found in slice.
	for _, i := range animals {
		fmt.Println(i)

	}

	// print a specific slice.
	fmt.Println("The first element is", animals[0])

	// print more then one element.
	fmt.Println("The first two elements are", animals[0:2])

	// print the lengeth of the slice.
	fmt.Println("the length of the slice is", len(animals), "elements long ")

	// check to see if the slice is sorted.
	fmt.Println("Is the slice sorted?", sort.StringsAreSorted(animals))

	// sort the slice.
	sort.Strings(animals)

	// check again to see if the slice is now sorted.
	fmt.Println("Is the slice sorted?", sort.StringsAreSorted(animals))

	// print out animals to show that it's been sorted.
	fmt.Println(animals)

	deleteFromSlice(animals, 3)

	fmt.Println(animals)

}

func deleteFromSlice(a []string, i int) []string {
	a[i] = a[len(a)-1]
	a[len(a)-1] = ""
	a = a[:len(a)-1]
	return a
}
