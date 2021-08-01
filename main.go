package main

import (
	"fmt"

	"github.com/eiannone/keyboard"
)

var KeyPressChan chan rune

func main() {
	KeyPressChan = make(chan rune)

	go listenForKeyPress()

	fmt.Println("Press any key or q to quit.")
	_ = keyboard.Open()
	defer func() {
		keyboard.Close()
	}()

	for {
		char, _, _ := keyboard.GetSingleKey()
		if char == 'q' || char == 'Q' {
			break
		}

		KeyPressChan <- char
	}
}

func listenForKeyPress() {
	for {
		key := <-KeyPressChan
		fmt.Println("You Pressed", string(key))
	}
}
