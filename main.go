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
	ROCK     = 0 // beats Scissors. (SCISSORS + 1) % 3 = 0
	PAPER    = 1 // beats Rock. (ROCK + 1) % 3 = 1
	SCISSORS = 2 // beats Paper (PAPER +1) % 3 = 2
)

func main() {
	rand.Seed(time.Now().UnixNano())
	playerChoice := ""
	playerValue := -1
	// *** added two variables to keep track of score
	playerScore := 0
	computerScore := 0

	reader := bufio.NewReader(os.Stdin)

	clearScreen()

	// *** print out some instructions
	fmt.Println("Rock, Paper & Scissors")
	fmt.Println("----------------------")
	fmt.Println("Game is played for three rounds, and best of three wins the game. Good luck!")
	fmt.Println()

	// *** added the for loop
	for i := 1; i <= 3; i++ {
		// *** print out round information
		fmt.Println()
		fmt.Println("Round", i)
		fmt.Println("-------")

		fmt.Print("Please enter rock, paper, or scissors -> ")
		playerChoice, _ = reader.ReadString('\n')
		playerChoice = strings.Replace(playerChoice, "\n", "", -1)

		// moved computer choice to for loop, so it's reset each time through
		computerValue := rand.Intn(3)

		if playerChoice == "rock" {
			playerValue = ROCK
		} else if playerChoice == "paper" {
			playerValue = PAPER
		} else if playerChoice == "scissors" {
			playerValue = SCISSORS
		} else {
			playerValue = -1
		}

		// *** print out player choice
		fmt.Println()
		fmt.Println("Player chose", strings.ToUpper(playerChoice))

		switch computerValue {
		case ROCK:
			fmt.Println("Computer chose ROCK")
			break
		case PAPER:
			fmt.Println("Computer chose PAPER")
			break
		case SCISSORS:
			fmt.Println("Computer chose SCISSORS")
			break
		default:
		}

		fmt.Println()

		if playerValue == computerValue {
			fmt.Println("It's a draw")
			// *** decrement i by 1, since we're repeating the round
			i--
		} else if playerValue == -1 {
			fmt.Println("Invalid Choice!")
			i--
		} else if playerValue == (computerValue +1) % 3 {
			playerScore = playerWins(playerScore)
		} else {
			computerScore = computerWins(computerScore)
		}

	fmt.Println("Final score")
	fmt.Println("-----------")
	fmt.Printf("Player: %d/3, Computer %d/3", playerScore, computerScore)
	fmt.Println()
	if playerScore > computerScore {
		fmt.Println("Player wins game!")
	} else {
		fmt.Println("Computer wins game!")
	}
}

func computerWins(d int) int {
	fmt.Println("Computer wins!")
	d++
	return d
}

func playerWins(d int) int {
	fmt.Println("Player wins!")
	d++
	return d
}

// clearScreen clears the screen
func clearScreen() {
	if strings.Contains(runtime.GOOS, "windows") {
		// windows
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
