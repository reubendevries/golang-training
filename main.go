package main

import (
	"myapp/packageone"
)

// package level variable called myVar
var myVar = "This is a the package level variable."

func main() {
	// variables

	// block level variable called blockVar
	var blockVar = "This is the block level variable"

	// call the Public function called PrintMe
	packageone.PrintMe(myVar, blockVar, packageone.PackageVar)
}
