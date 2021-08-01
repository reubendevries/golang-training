package main

import "fmt"

type Animal struct {
	Name         string
	Sound        string
	NumberOfLegs int
}

func (a *Animal) Says() {
	fmt.Printf("A %s says %s.\n", a.Name, a.Sound)
}

func (a *Animal) HowManyLegs() {
	fmt.Printf("A %s has %d number of legs.\n", a.Name, a.NumberOfLegs)
}

func main() {
	var dog Animal
	dog.Name = "Dog"
	dog.Sound = "Woof"
	dog.NumberOfLegs = 4
	dog.Says()
	dog.HowManyLegs()

	cat := Animal{
		Name:         "Cat",
		Sound:        "Meow",
		NumberOfLegs: 4,
	}
	cat.Says()
	cat.HowManyLegs()

	sumMany(1, 2, 3, 4, 5, 6, 7, 8, 9)
}

func sumMany(nums ...int) int {

	total := 0

	for _, x := range nums {
		total = total + x
	}
	return total
}
