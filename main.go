package main

import "log"

// basic types (numbers, strings, boolean)
var myInt int
var myUint uint
var myFloat32 float32
var myFloat64 float64

func main() {
	myInt = 10
	myUint = 20

	myFloat32 = 10.1
	myFloat64 = 100.1

	log.Println(myInt, myUint, myFloat32, myFloat64)

	// in Go strings are immutable, which means they shouldn't
	// be changed, they still can be changed but this isn't best
	// practice as the change will force your program to allocate
	// more memory to “change” that data
	myString := "Reuben"
	log.Println(myString)

	// this will take more memory, should only be used when absoultely
	// necessary.
	//myString = "John"
	//log.Println(myString)

	var myBool = true
	myBool = false
	log.Println(myBool)

}
