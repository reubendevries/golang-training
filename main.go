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
	// randomizes the computerValue
	rand.Seed(time.Now().UnixNano())

	computerValue := rand.Intn(2)

	// create a var called reader that reads from Standard Input.
	reader := bufio.NewReader(os.Stdin)
	// sets var gameCount int to 0
	gameCount := 0
	for gameCount <= 3 {
		// runs clear screen
		clearScreen()

		// sets the playerChoice to blank, before player can choose
		playerChoice := ""

		// prints the following line
		fmt.Print("Please enter rock, paper, or scissors -> ")

		// lets input know it's finished when carriage return is processed by keyboard.
		playerChoice, _ = reader.ReadString('\n')

		// trims the carriage return from the players choice.
		playerChoice = strings.TrimSpace(playerChoice)

		// function that returns the playerValue based on choice inputed.
		playerValue := playersChoice(playerChoice)

		// function that return the winner of rock, paper, scissors
		winner := determineWinner(computerValue, playerValue)

		// prints out the winner
		fmt.Printf("Result: %s\n", winner)

		// add the game count so that loop will end after three games.
		gameCount++
	}
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

// function that returns the playerValue depending on what
// the player selected originally.
func playersChoice(str string) int {
	// sets playerValue to -1
	playerValue := -1
	// if string is rock then assigns playerValue to 0
	if str == "rock" {
		playerValue = ROCK
		// if string is paper then assigns playerValue to 1
	} else if str == "paper" {
		playerValue = PAPER
		// if string is scissors then assigns playerValue to 2
	} else {
		playerValue = SCISSORS
	}
	// returns player value that has been set.
	return playerValue
}

// a function to determine the winner
func determineWinner(n, m int) string {
	var winner string
	// if tie, then it's a draw
	if n == m {
		winner = "It's A Draw"
		// else if computer has scissors and player has rock
		// then player wins
	} else if n == 2 && m == 0 {
		winner = "Congratulations, You Win!"
		// else if computer is greater than players choice then
		// computer wins
	} else if n > m {
		winner = "Sorry, You Lose."
		// else player wins
	} else {
		winner = "Congratulations, You Win!"
	}
	// returns the winner
	return winner
}
