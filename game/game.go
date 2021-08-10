package game

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

var reader = bufio.NewReader(os.Stdin)

type Game struct {
	DisplayChan chan string
	RoundChan   chan int
	Round       Round
}

type Round struct {
	RoundNumber   int
	PlayerScore   int
	ComputerScore int
}

func (g *Game) Rounds() {
	// use select to process input in channels
	// print to screen
	// keep trace of the round number
	for {
		select {
		case round := <-g.RoundChan:
			g.Round.RoundNumber = g.Round.RoundNumber + round
			g.RoundChan <- 1
		case msg := <-g.DisplayChan:
			fmt.Println(msg)
			g.DisplayChan <- ""
		}
	}
}

//Clear Screen clears the screen
func (g *Game) ClearScreen() {
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

func (g *Game) PrintIntro() {
	g.DisplayChan <- fmt.Sprintf(`
Rock, Paper & Scisors
---------------------
Game is played for three rounds, and best of three wins the game. Good luck!	
`)
	<-g.DisplayChan
}

func (g *Game) PlayRound() bool {
	rand.Seed(time.Now().UnixNano())
	playerValue := -1

	g.DisplayChan <- fmt.Sprintf(`
	
Round %d
--------
`, g.Round.RoundNumber)
	<-g.DisplayChan

	fmt.Print("Please enter rock, paper or scissors -> ")

	playerChoice, _ := reader.ReadString('\n')
	playerChoice = strings.TrimSpace(playerChoice)

	computerValue := rand.Intn(3)

	if playerChoice == "rock" {
		playerValue = ROCK
	} else if playerChoice == "paper" {
		playerValue = PAPER
	} else if playerChoice == "scissors" {
		playerValue = SCISSORS
	}

	g.DisplayChan <- ""
	<-g.DisplayChan

	g.DisplayChan <- fmt.Sprintf("Player chose %s\n", strings.ToUpper(playerChoice))
	<-g.DisplayChan
	switch computerValue {
	case ROCK:
		g.DisplayChan <- "Computer chose ROCK"
		<-g.DisplayChan
	case PAPER:
		g.DisplayChan <- "Computer chose PAPER"
		<-g.DisplayChan
	case SCISSORS:
		g.DisplayChan <- "Computer chose SCISSORS"
		<-g.DisplayChan
	default:
	}

	if playerValue == computerValue {
		g.DisplayChan <- "It's a draw!"
		<-g.DisplayChan
		return false
	} else {
		switch playerValue {
		case ROCK:
			if computerValue == PAPER {
				g.computerWins()
			} else {
				g.playerWins()
			}
		case PAPER:
			if computerValue == SCISSORS {
				g.computerWins()
			} else {
				g.playerWins()
			}
		case SCISSORS:
			if computerValue == ROCK {
				g.computerWins()
			} else {
				g.playerWins()
			}
		default:
			// decrement i by 1, since wer're repeating the round
			g.DisplayChan <- "Invalid Choice!"
			<-g.DisplayChan
			return false
		}
	}

	return true
}

func (g *Game) computerWins() {
	g.Round.ComputerScore++
	g.DisplayChan <- "Computer Wins"
	<-g.DisplayChan
}

func (g *Game) playerWins() {
	g.Round.PlayerScore++
	g.DisplayChan <- "Player Wins"
	<-g.DisplayChan
}

func (g *Game) PrintSummary() {
	g.DisplayChan <- fmt.Sprintf(`
Final Score
-----------

Player Scored: %d/3
Computer Scored: %d/3

`, g.Round.PlayerScore, g.Round.ComputerScore)
	<-g.DisplayChan
	if g.Round.PlayerScore > g.Round.ComputerScore {
		g.DisplayChan <- "Congratulations, You Win!"
		<-g.DisplayChan

		g.DisplayChan <- ""
		<-g.DisplayChan
	} else {
		g.DisplayChan <- "Sorry You Lose, Better Luck Next Time!"
		<-g.DisplayChan
		g.DisplayChan <- ""
		<-g.DisplayChan
	}
}
