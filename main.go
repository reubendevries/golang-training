package main

import (
	"fmt"
	"log"
	"strconv"

	"github.com/eiannone/keyboard"
)

var coffees = []string{
	"Cappucino",
	"Latte",
	"Americano",
	"Mocha",
	"Macchiato",
	"Espresso",
}

func main() {

	// opens keyboard to listen to single keys.
	err := keyboard.Open()
	if err != nil {
		log.Fatal(err)
	}

	// defer function to close keyboard.
	defer func() {
		_ = keyboard.Close()
	}()

	// sets the for while loop to ON
	var answer bool
	answer = true

	// for while loop, which will print the menu out, let you select the coffe
	// brew your coffee and give you an option to brew another coffee if you
	// select no then it will set the anwer to false, terminiating the loop.
	for answer {
		printMenu(coffees, 6)
		char := selectCoffee()
		brewCoffee(char, coffees, 6)
		answer = stillThirsty()
	}
}

func printMenu(arr []string, size int) {

	// function that prints the menu
	fmt.Println("")
	fmt.Println("Please select which type of Coffee you would like:")
	fmt.Println("")
	fmt.Println("MENU")
	fmt.Println("----")
	// for loop to print out the coffee menuß
	for i, v := range arr {
		fmt.Printf("%d - %s\n", (i + 1), v)
	}
}

func selectCoffee() rune {

	// function that listens to return a single key.
	char, _, err := keyboard.GetSingleKey()
	if err != nil {
		log.Fatal(err)
	}
	// returns the char that you pressed on your keyboard
	return char
}

func brewCoffee(r rune, arr []string, size int) {

	// function that accept the char converts it into a string and then
	// evaulates the string that then prints out the coffee you've chosen.
	i, _ := strconv.Atoi(string(r))
	fmt.Println("")
	fmt.Printf("You chose %s \n", arr[i-1])
	fmt.Println("")
	fmt.Printf("Now Brewing your cup of %s\n", arr[i-1])
	fmt.Println("")
}

func stillThirsty() bool {

	// function that asks if you want to brew another coffee.
	// if you select yes returns true, if you select no returns false.
	fmt.Println("")
	fmt.Println("Would you like another coffee? Press 'Y' for another coffee or 'N' to quit.")
	fmt.Println("")
	char, _, err := keyboard.GetSingleKey()
	if err != nil {
		log.Fatal(err)
	}
	if char == 'n' || char == 'N' {
		println("Thanks, have a great day!")
		return false
	}
	return true
}
