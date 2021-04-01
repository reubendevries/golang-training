package main

import (
	"log"

	"github.com/reubendevries/myniceprogram/helpers"
)

const numPool = 10

func CalculateValue(intchan chan int) {
	randomNumber := helpers.RandomNumber(numPool)
	intchan <- randomNumber
}

func main() {
	intChan := make(chan int)
	defer close(intChan)

	go CalculateValue(intChan)

	num := <-intChan
	log.Println(num)
}
