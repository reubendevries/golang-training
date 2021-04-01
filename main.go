package main

import (
	"log"

	"github.com/reubendevries/myniceprogram/helpers"
)

func main() {
	log.Println("Hello")

	var myVar helpers.SomeType

	myVar.TypeName = "Some Name"

	log.Println(myVar.TypeName)
}
