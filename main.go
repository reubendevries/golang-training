package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/eiannone/keyboard"
)

var reader *bufio.Reader

type User struct {
	UserName        string
	Age             int
	FavouriteNumber float64
	OwnsADog        bool
}

func main() {
	reader = bufio.NewReader(os.Stdin)
	var user User
	user.UserName = readString("What is your name?")
	user.Age = readInteger("What is your age?")
	user.FavouriteNumber = readFloat("What is your favourite number")
	user.OwnsADog = readBool("Do you own a dog?")
	fmt.Printf("Your name is %s. You're %d years old, your favourite number is %f\n and you own a dog: %t", user.UserName, user.Age, user.FavouriteNumber, user.OwnsADog)
}

func prompt() {
	fmt.Print("-> ")
}

func readString(s string) string {
	fmt.Println(s)
	prompt()

	userInput, _ := reader.ReadString('\n')
	userInput = strings.TrimSpace(userInput)

	return userInput
}

func readInteger(s string) int {
	for {
		fmt.Println(s)
		prompt()

		userInput, _ := reader.ReadString('\n')
		userInput = strings.TrimSpace(userInput)

		num, err := strconv.Atoi(userInput)
		if err != nil {
			println("Please enter in a whole number.")
		} else {
			return num
		}
	}
}

func readFloat(s string) float64 {
	for {
		fmt.Println(s)
		prompt()

		userInput, _ := reader.ReadString('\n')
		userInput = strings.TrimSpace(userInput)

		num, err := strconv.ParseFloat(userInput, 64)
		if err != nil {
			println("Please enter in a number.")
		} else {
			return num
		}
	}
}

func readBool(s string) bool {
	for {
		fmt.Println(s)
		prompt()

		err := keyboard.Open()
		if err != nil {
			log.Fatal(err)
		}

		defer func() {
			_ = keyboard.Close()
		}()

		char, _, err := keyboard.GetSingleKey()
		if err != nil {
			log.Fatal(err)
		}
		if strings.ToLower(string(char)) != "y" || strings.ToLower(string(char)) != "n" {
			fmt.Println("Please select either Y or N")
		} else if char == 'y' || char == 'Y' {
			return true
		} else if char == 'n' || char == 'Y' {
			return false
		}
	}
}
