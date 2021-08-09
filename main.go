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
	playerValue := -1

	computerValue := rand.Intn(3)

	reader := bufio.NewReader(os.Stdin)

	clearScreen()
	gameCount := 0
	playerScore := 0
	computerScore := 0

	for gameCount < 3 {
		if gameCount == 0 {
			fmt.Println("Round One")
			fmt.Println("_________")
			fmt.Println()
		} else if gameCount == 1 {
			fmt.Println("Round Two")
			fmt.Println("_________")
			fmt.Println()
		} else {
			fmt.Println("Round Three")
			fmt.Println("___________")
			fmt.Println()
		}
		fmt.Print("Please enter rock, paper or scissors -> ")
		playerChoice, _ = reader.ReadString('\n')
		playerChoice = strings.Replace(playerChoice, "\n", "", -1)
		if playerChoice == "rock" {
			playerValue = ROCK
		} else if playerChoice == "paper" {
			playerValue = PAPER
		} else if playerChoice == "scissors" {
			playerValue = SCISSORS
		}

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
		} else {
			switch playerValue {
			case ROCK:
				if computerValue == PAPER {
					fmt.Println("Computer Wins")
					gameCount++
					computerScore++
				} else {
					fmt.Println("Player Wins")
					gameCount++
					playerScore++
				}
				break
			case PAPER:
				if computerValue == SCISSORS {
					fmt.Println("Computer Wins")
					gameCount++
					computerScore++
				} else {
					fmt.Println("Player Wins")
					gameCount++
					playerScore++
				}
				break
			case SCISSORS:
				if computerValue == ROCK {
					fmt.Println("Computer Wins")
					gameCount++
					computerScore++
				} else {
					fmt.Println("Player Wins")
					gameCount++
					playerScore++
				}
				break
			default:
				fmt.Println("Invalid Choice!")
			}
		}
	}
	clearScreen()
	fmt.Println("Final Score")
	fmt.Println("___________")
	fmt.Println()
	fmt.Printf("Player Scored: %d/3\n", playerScore)
	fmt.Printf("Computer Scored: %d/3\n", computerScore)
	if playerScore > computerScore {
		fmt.Println("Congratulations, You Win!")
		fmt.Println()
	} else {
		fmt.Println("Sorry You Lose, Better Luck Next Time!")
		fmt.Println()
	}
}

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
