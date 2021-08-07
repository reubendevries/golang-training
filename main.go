package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"os/exec"
	"runtime"
	"strings"
	"time"
)

const (
	ROCK     = 0
	PAPER    = 1
	SCISSORS = 2
)

func main() {
	rand.Seed(time.Now().UnixNano())
	playerChoice := ""

	computerValue := rand.Intn(2)

	reader := bufio.NewReader(os.Stdin)

	clearScreen()

	fmt.Print("Please enter rock, paper, or scissors -> ")
	playerChoice, _ = reader.ReadString('\n')
	playerChoice = strings.TrimSpace(playerChoice)

	playerValue := playersChoice(playerChoice)

	fmt.Println()
	fmt.Println("You have chosen: ", playerChoice, "and value is: ", playerValue)
	fmt.Println("Computer has chosen: ", computerValue)

}

// clearScreen clears the screen
func clearScreen() {
	// windows
	if strings.Contains(runtime.GOOS, "Windows") {
		cmd := exec.Command("cmd", "/c", "cls")
		cmd.Stdout = os.Stdout
		cmd.Run()
	} else {
		// linux or mac
		cmd := exec.Command("clear")
		cmd.Stdout = os.Stdout
		cmd.Run()
	}
}

func playersChoice(str string) int {
	playerValue := -1

	if str == "rock" {
		playerValue = ROCK
	} else if str == "paper" {
		playerValue = PAPER
	} else {
		playerValue = SCISSORS
	}
	return playerValue
}
