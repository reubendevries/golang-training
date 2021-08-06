package main

import (
	"bufio"
	"fmt"
	"myapp/doctor"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	whatToSay := doctor.Intro()

	fmt.Println(whatToSay)

	// execute everything in this loop until the user enters 'quit' and presses return.
	for {
		// print out a prompt
		fmt.Print("-> ")

		// wait for user to print something and press enter
		userInput, _ := reader.ReadString('\n')

		// strip the carriage return off of wahtever they typed.
		userInput = strings.Replace(userInput, "\r\n", "", -1)
		userInput = strings.Replace(userInput, "\n", "", -1)

		fmt.Println(doctor.Response(userInput))

		if userInput != "quit" {
			// end program
			break
		}
	}
}
