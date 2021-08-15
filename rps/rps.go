package rps

import (
	"math/rand"
	"time"
)

const (
	ROCK     = 0 // beats scissors. (scissors + 1) % 3 = 0
	PAPER    = 1 // beats rock. (rock + 1) % 3 = 1
	SCISSORS = 2 // beats paper. (paper + 1) % 3 = 2

)

var winMessage = []string{
	1: "Winner, Winner, Chicken Dinner",
	2: "It's your Lucky Day",
	3: "Good Job, You Won!",
}

var drawMessage = []string{
	1: "Try Again",
	2: "Great Minds Think Alike",
	3: "It's a tie!",
}

var loseMessage = []string{
	1: "Sorry, Computer Wins",
	2: "Better Luck Next Time",
	3: "You Lose!",
}

type Round struct {
	Message        string `json:"message"`
	ComputerChoice string `json:"computer_choice"`
	RoundResult    string `json:"round_result"`
}

func PlayRound(playerValue int) Round {
	rand.Seed(time.Now().UnixNano())
	computerValue := rand.Intn(3)
	computerChoice := ""
	roundResult := ""

	switch computerValue {
	case ROCK:
		computerChoice = "Computer chose ROCK"
	case PAPER:
		computerChoice = "Computer chose PAPER"
	case SCISSORS:
		computerChoice = "Computer chose SCISSORS"
	}

	messageInt := rand.Intn(3)
	message := ""

	if playerValue == computerValue {
		roundResult = "It's a draw"
		message = drawMessage[messageInt]
	} else if playerValue == (computerValue+1)%3 {
		roundResult = "Player Wins!"
		message = winMessage[messageInt]
	} else {
		roundResult = "Computer Wins!"
		message = loseMessage[messageInt]
	}

	var result Round
	result.Message = message
	result.ComputerChoice = computerChoice
	result.RoundResult = roundResult
	return result
}
